package consensus

import (
	"compute/model"
	"math/rand"
	"os"
	"strconv"
	"time"

	cfg "compute/config"

	zlog "compute/log"
	"github.com/hashicorp/raft"
	raftDb "github.com/hashicorp/raft-boltdb"
	"go.uber.org/zap"
)

type Raft struct {
	*fsm
	raft     *raft.Raft
	members  []string
	tempFile []string
	raftAddr string
}

func NewRaft(members []string) (*Raft, error) {
	rcfg := raft.DefaultConfig()

	leaderNotifyCh := make(chan bool, 1)
	rcfg.NotifyCh = leaderNotifyCh
	rcfg.NoSnapshotRestoreOnStart = false
	rcfg.LogLevel = "ERROR"
	rcfg.LocalID = raft.ServerID(raftId(members))

	transport, err := raft.NewTCPTransport()
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
	}, nil
}

func (r *Raft) BuildConsensus() error {

	return nil
}

// ShutDown 关闭共识
func (r *Raft) ShutDown() error {
	return r.raft.Shutdown().Error()
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
				r.pushValue(model.ConsensusReq{
					NodeName: "",
					Serial:   0,
					Value:    model.ConsensusValue{},
				})
				break
			} else {
				// request leader to pushValue
				// TODO
				break
			}

		}
	}
}

func (r *Raft) Watch(targetNode string) model.ConsensusValue {
	waitCh := make(chan model.ConsensusValue, 1)
	r.WatchValue(targetNode, waitCh)
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
