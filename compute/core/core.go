package core

import (
	"sync"

	"compute/client"
	"compute/consensus"
	"compute/model"
)

var C = NewCore()

type Core struct {
	*processTable
	lock sync.Mutex
}

func NewCore() *Core {
	return &Core{
		processTable: newPT(),
	}
}

// StartProcess 启动进程,准备pid,建立共识节点,建立共识,启动容器发起计算
func (c *Core) StartProcess(pid model.Pid, funcId string, members []string, callback chan<- *model.Output, errCh chan error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	p := new(pcb)
	p.callback = callback
	if len(members) == 1 {
		// 只有一个成员,单节点计算,不需要建立共识,同时容器内直接执行即可
		goto label
	} else {
		// 多节点参与计算
		r, err := consensus.NewRaft(members)
		if err != nil {
			errCh <- err
		}
		p.consensus = r
		if err = r.BuildConsensus(); err != nil {
			errCh <- err
		}
	}
label:
	if err := client.GenerateTempFile(funcId); err != nil {
		errCh <- err
	}
	// 将当前进程插入进程表
	c.put(pid, p)
	go c.startContainer(errCh, callback)
}

// startContainer 启动容器执行计算方法
func (c *Core) startContainer(errCh chan error, callback chan<- *model.Output) {

}
