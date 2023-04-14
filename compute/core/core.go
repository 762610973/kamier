package core

import (
	"sync"

	"compute/model"
)

var C = NewCore()

type Core struct {
	*model.ProcessTable
	lock         sync.Mutex
	prepareTable sync.Map
}

func NewCore() *Core {
	return &Core{
		ProcessTable: model.NewPT(),
		lock:         sync.Mutex{},
		prepareTable: sync.Map{},
	}
}

// StartProcess 启动进程,准备pid,建立共识节点,建立共识,启动容器发起计算
func (c *Core) StartProcess(pid *model.Pid, funcId string, members []string, callback chan<- *model.Output) {

}

func (c *Core) startContainer() {

}
