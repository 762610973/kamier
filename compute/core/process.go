package core

import (
	"errors"
	"os"
	"sync"

	"compute/consensus"
	zlog "compute/log"
	"compute/model"

	"go.uber.org/zap"
)

// processTable 进程表,key: pid;value: pcb
type processTable struct {
	m         sync.Map
	callback  chan<- *model.Output
	consensus consensus.Raft
}

type pcb struct {
	prepared  *prepareValue
	callback  chan<- *model.Output
	consensus *consensus.Raft
	// 存储计算方法的临时文件路径
	tempFilePath string
}

func newPT() *processTable {
	return &processTable{m: sync.Map{}}
}

func (pT *processTable) put(pid model.Pid, pcb *pcb) {
	pT.m.Store(pid, pcb)
}

func (pT *processTable) get(pid model.Pid) (p *pcb, ok bool) {
	val, ok := pT.m.Load(pid)
	if ok {
		return val.(*pcb), true
	}
	return nil, false
}
func (pT *processTable) delete(pid model.Pid) {
	pT.m.Delete(pid)
}

// shutdown 关闭进程
func (p *pcb) shutdown() (err error) {
	err = p.consensus.ShutDown()
	if err != nil {
		zlog.Error("shutdown consensus failed", zap.Error(err))
		err = errors.Join(err)
	}
	if err = os.Remove(p.tempFilePath); err != nil {
		zlog.Error("remove temp file failed", zap.Error(err))
		err = errors.Join(err)
	}
	return err
}
