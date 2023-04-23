package consensus

import (
	"encoding/json"
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

	leaderNotifyCh := make(chan bool, 1)
	rcfg.NotifyCh = leaderNotifyCh
	rcfg.NoSnapshotRestoreOnStart = false
	rcfg.LogLevel = "ERROR"
	rcfg.LocalID = raft.ServerID(raftId(members))
	addr, err := getPort(pid, fmt.Sprintf("http://%s:%s", cfg.Cfg.LocalAddr, cfg.Cfg.GrpcPort))
	if err != nil {
		zlog.Error("get port failed", zap.Error(err))
		return nil, err
	}
	tcpAddr, _ := net.ResolveTCPAddr("tcp", addr)
	transport, err := raft.NewTCPTransport(tcpAddr.String(), tcpAddr, 4, 4*time.Second, os.Stderr)
	if err != nil {
		zlog.Error("new tcp transport failed", zap.Error(err))
		return nil, err
	}
	snapshotStore := raft.NewInmemSnapshotStore()
	zlog.Debug("new memory snapshot store")
	temp := randString(10)
	lPath := "raftFile/logPath/" + temp
	_ = os.MkdirAll(lPath, 0777)
	logStore, err := raftDb.NewBoltStore(lPath)
	if err != nil {
		zlog.Error("new bolt store failed")
		return nil, err
	}
	sPath := "raftFile/stableStorePath/" + temp
	_ = os.MkdirAll(sPath, 0777)
	stableStore, err := raftDb.NewBoltStore(sPath)
	if err != nil {
		zlog.Error("new stable store failed", zap.Error(err))
		return nil, err
	}
	rf, err := raft.NewRaft(rcfg, newFsm(), logStore, stableStore, snapshotStore, transport)
	if err != nil {
		zlog.Error("new raft failed", zap.Error(err))
		return nil, err
	}
	zlog.Info("new raft node success")
	return &Raft{
		raft:     rf,
		tempFile: []string{lPath, sPath},
		pid:      pid,
	}, nil
}

func getPort(pid model.Pid, addr string) (string, error) {
	addr = addr[7:]
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
	tempInt, err := strconv.ParseInt(n, 64, 10)
	if err != nil {
		zlog.Error("parse int64 failed", zap.Error(err), zap.String("int64->", n))
		return "", err
	}
	// 将端口转换为数字
	portInt, err := strconv.ParseInt(port, 64, 10)
	if err != nil {
		zlog.Error("parse port to int64 failed", zap.Error(err), zap.String("int64->", n))
		return "", err
	}
	port = strconv.FormatInt(portInt+tempInt, 10)
	return host + ":" + port, nil

}

// BuildConsensus 构建集群共识
func (r *Raft) BuildConsensus() error {
	zlog.Info("build consensus")
	serverCfg := r.raft.GetConfiguration().Configuration()
	if len(serverCfg.Servers) > 0 {
		return errors.New("start failed, config exists")
	}
	for idx, node := range r.members {
		id := strconv.Itoa(idx)
		addr, err := getPort(r.pid, client.Nodemap.Get(node))
		if err != nil {
			return err
		}
		server := raft.Server{
			ID:      raft.ServerID(id),
			Address: raft.ServerAddress(addr),
		}
		serverCfg.Servers = append(serverCfg.Servers, server)
	}
	if err := r.raft.BootstrapCluster(serverCfg).Error(); err != nil {
		zlog.Error("boot strap cluster failed", zap.Error(err))
		return err
	}
	zlog.Info("build consensus success")
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
	err = errors.Join(r.raft.Shutdown().Error())
	return err
}

func raftId(members []string) string {
	var res int
	for idx, member := range members {
		if member == cfg.Cfg.NodeName {
			res = idx + 1
			break
		}
	}
	return strconv.Itoa(res)
}

func (r *Raft) Push(pid model.Pid, arg []byte) {
	// get leader and ipc to push value

	for {
		leaderAddr, leaderId := r.raft.LeaderWithID()
		if string(leaderId) == "" {
			//还没有选举出leader
			time.Sleep(time.Duration(cfg.Cfg.LeaderElection) * time.Millisecond)
		} else {
			// leader已经选举出来了
			// 判断是否是本节点
			if string(leaderAddr) == r.raftAddr {
				var v model.ConsensusValue
				err := json.Unmarshal(arg, &v)
				if err != nil {
					zlog.Panic("unmarshal consensus arg failed", zap.Error(err))
				}
				cmd, err := json.Marshal(model.ConsensusReq{
					NodeName: pid.NodeName,
					Serial:   pid.Serial,
					Value: model.ConsensusValue{
						Type:  v.Type,
						Value: v.Value,
					},
				})
				if err != nil {
					zlog.Error("marshal consensus value failed", zap.Error(err))
				}
				apply := r.raft.Apply(cmd, 4*time.Second)
				if apply.Error() != nil {
					zlog.Panic("raft apply failed")
				}
				break
			} else {
				err := client.Ipc(string(leaderAddr), pid, arg)
				if err != nil {
					zlog.Error("ipc failed", zap.Error(err))
				}
				break
			}

		}
	}
}

func (r *Raft) Watch(targetNode string) model.ConsensusValue {
	waitCh := make(chan model.ConsensusValue, 1)
	r.watchValue(targetNode, waitCh)
	defer zlog.Debug("watch value from consensus success")
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
