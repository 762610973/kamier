package model

import (
	"sync"
)

// Pid 每次计算的唯一标识: 名称+计算序列号
type Pid struct {
	NodeName string // 节点名称
	Serial   int64  // 节点本次计算的唯一序列号
}

type Output struct {
}
type Pcb struct {
}

// ProcessTable 进程表,key: pid;value: pcb
type ProcessTable struct {
	m sync.Map
}

func NewPT() *ProcessTable {
	return &ProcessTable{sync.Map{}}
}

func (p *ProcessTable) Put(pid Pid, pcb *Pcb) {
	p.m.Store(pid, pcb)
}

func (p *ProcessTable) Get(pid Pid) (bool, *Pcb) {
	val, ok := p.m.Load(pid)
	if ok {
		return true, val.(*Pcb)
	}
	return false, nil
}
func (p *ProcessTable) Delete(pid Pid) {
	p.m.Delete(pid)
}
