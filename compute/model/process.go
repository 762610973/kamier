package model

import (
	"sync"
)

// Pid 每次计算的唯一标识: 名称+计算序列号
type Pid struct {
	OrgName string // 节点名称
	Serial  int    // 节点本次计算的唯一序列号
}

type Output struct {
}
type Pcb struct {
}

type ProcessTable struct {
	sync.Map
}

func NewPT() *ProcessTable {
	return &ProcessTable{sync.Map{}}
}

func (p *ProcessTable) Put(pid Pid, pcb *Pcb) {
	p.Map.Store(pid, pcb)
}

func (p *ProcessTable) Get(pid Pid) (bool, *Pcb) {
	val, ok := p.Map.Load(pid)
	if ok {
		return true, val.(*Pcb)
	}
	return false, nil
}
func (p *ProcessTable) Delete(pid Pid) {
	p.Map.Delete(pid)
}
