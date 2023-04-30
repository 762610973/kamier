package consensus

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"compute/client"
	cfg "compute/config"
	zlog "compute/log"
	"compute/model"

	"github.com/hashicorp/raft"
	raftDb "github.com/hashicorp/raft-boltdb"
	"github.com/imroc/req/v3"
	"go.uber.org/zap"
)

type Raft struct {
	*fsm
	raft     *raft.Raft
	pid      model.Pid
	members  []string
	tempFile []string
	raftAddr string
}

// NewRaft new a raft node
func NewRaft(pid model.Pid, members []string) (*Raft, error) {
	rcfg := raft.DefaultConfig()
	rcfg.SnapshotThreshold = 3
	rcfg.SnapshotInterval = 10 * time.Second
	leaderNotifyCh := make(chan bool, 1)
	rcfg.NotifyCh = leaderNotifyCh
	rcfg.NoSnapshotRestoreOnStart = false
	rcfg.LogLevel = "ERROR"
	rcfg.LocalID = raft.ServerID(getRaftId(members))
	zlog.Debug("get raft id success", zap.Any("raftId", rcfg.LocalID))
	addr, err := getPort(pid, fmt.Sprintf("%s:%s", cfg.Cfg.LocalAddr, cfg.Cfg.GrpcPort))
	//addr, err := getPort(pid, fmt.Sprintf("http://%s:%s", cfg.Cfg.LocalAddr, cfg.Cfg.GrpcPort))
	if err != nil {
		zlog.Error("get port failed", zap.Error(err))
		return nil, err
	}
	zlog.Info("get port success", zap.Any("localAddr: ", addr))
	tcpAddr, _ := net.ResolveTCPAddr("tcp", addr)
	transport, err := raft.NewTCPTransport(addr, tcpAddr, 4, 4*time.Second, os.Stderr)
	if err != nil {
		zlog.Error("new tcp transport failed", zap.Error(err))
		return nil, err
	}
	snapshotStore := raft.NewInmemSnapshotStore()
	zlog.Debug("new memory snapshot store")
	temp := randString(10)
	lPath := "raftFile/logPath/"
	_ = os.MkdirAll(lPath, 0777)

	logStore, err := raftDb.NewBoltStore(lPath + temp)
	if err != nil {
		zlog.Error("new bolt store failed", zap.Error(err))
		return nil, err
	}
	sPath := "raftFile/stableStorePath/"
	_ = os.MkdirAll(sPath, 0777)
	stableStore, err := raftDb.NewBoltStore(sPath + temp)
	if err != nil {
		zlog.Error("new stable store failed", zap.Error(err))
		return nil, err
	}
	f := newFsm()
	rf, err := raft.NewRaft(rcfg, f, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		zlog.Error("new raft failed", zap.Error(err))
		return nil, err
	}
	zlog.Info("new raft node success")
	return &Raft{
		raft:     rf,
		tempFile: []string{lPath + temp, sPath + temp},
		pid:      pid,
		members:  members,
		fsm:      f,
		raftAddr: addr,
	}, nil
}

func getPort(pid model.Pid, addr string) (string, error) {
	//addr = addr[7:]
	host, port := "", ""
	for i, v := range addr {
		if v == ':' {
			host = addr[:i]
			port = addr[i+1:]
		}
	}
	res, err := req.R().SetBodyJsonMarshal(pid).Post(cfg.Cfg.StorageUrl + "/consensus/")
	if err != nil {
		zlog.Error("get consensus port num failed", zap.Error(err))
		return "", err
	}
	n := res.String()
	// 获取要递增的数字
	tempInt, err := strconv.Atoi(n)
	if err != nil {
		zlog.Error("parse int64 failed", zap.Error(err), zap.String("int64->", n))
		return "", err
	}
	// 将端口转换为数字
	portInt, err := strconv.Atoi(port)
	if err != nil {
		zlog.Error("parse port to int64 failed", zap.Error(err), zap.String("int64->", n))
		return "", err
	}
	port = strconv.Itoa(portInt + tempInt)
	zlog.Info("get port success")
	return host + ":" + port, nil

}

// BuildConsensus 构建集群共识
func (r *Raft) BuildConsensus() error {
	zlog.Info("build consensus")
	servers := r.raft.GetConfiguration().Configuration().Servers
	r.raft.GetConfiguration().Configuration()
	if len(servers) > 0 {
		zlog.Error("servers.len > 0")
		return errors.New("start failed, config exists")
	}
	servers = []raft.Server{}
	for idx, node := range r.members {
		id := strconv.Itoa(idx + 1)
		addr, err := getPort(r.pid, client.Nodemap.Get(node))
		if err != nil {
			zlog.Error("build consensus failed", zap.Error(err))
			return err
		}
		server := raft.Server{
			Suffrage: raft.Voter,
			ID:       raft.ServerID(id),
			Address:  raft.ServerAddress(addr),
		}
		zlog.Debug(fmt.Sprint(server))
		servers = append(servers, server)
	}
	zlog.Debug("cluster servers", zap.Any("servers: ", servers))
	if err := r.raft.BootstrapCluster(raft.Configuration{Servers: servers}).Error(); err != nil {
		zlog.Error("boot strap cluster failed", zap.Error(err))
		return err
	}
	zlog.Info("boot strap cluster success")
	return nil
}

// ShutDown 关闭共识
func (r *Raft) ShutDown() error {
	var err error
	for _, path := range r.tempFile {
		err = os.Remove(path)
		if err != nil {
			zlog.Error("remove path failed", zap.String("path", path))
			err = errors.Join(err)
		}
	}
	err = r.raft.Shutdown().Error()
	if err != nil {
		zlog.Error("shutdown consensus failed", zap.Error(err))
		err = errors.Join(err)
	}
	if err != nil {
		return err
	}
	return nil
}

func getRaftId(members []string) string {
	for idx, member := range members {
		if member == cfg.Cfg.NodeName {
			return strconv.Itoa(idx + 1)
		}
	}
	return ""
}

func (r *Raft) Push(pid model.Pid, arg []byte) {
	for {
		leaderAddr, leaderId := r.raft.LeaderWithID()
		if string(leaderId) == "" {
			zlog.Debug("leader is coming")
			time.Sleep(time.Duration(cfg.Cfg.LeaderElection) * time.Millisecond)
		} else {
			// leader已经选举出来了
			// 判断是否是本节点
			if string(leaderAddr) == r.raftAddr {
				zlog.Debug("self node is leader")
				apply := r.raft.Apply(arg, 4*time.Second)
				if apply.Error() != nil {
					zlog.Panic("raft apply failed")
				}
				break
			} else {
				zlog.Debug("self node is not leader")
				s := string(leaderId)
				id, err := strconv.Atoi(s)
				if err != nil {
					zlog.Error("convert string to int failed", zap.Error(err))
				}
				leaderName := r.members[id-1]
				err = client.Ipc(leaderName, pid, arg)
				if err != nil {
					zlog.Error("ipc failed", zap.Error(err))
				}
				break
			}

		}
	}
}

func (r *Raft) Watch(targetNode string) model.ConsensusValue {
	zlog.Debug("watch value", zap.String("targetNode", targetNode))
	waitCh := make(chan model.ConsensusValue, 1)
	r.fsm.watchValue(targetNode, waitCh)
	zlog.Debug("watch value from consensus")
	return <-waitCh
}

const randTempPath = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rune(randTempPath[rand.Intn(len(randTempPath))])
	}
	return string(b)
}
