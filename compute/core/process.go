package core

import (
	"sync"

	"compute/consensus"
	"compute/model"
)

type pcb struct {
	prepared  *prepareValue
	callback  chan<- *model.Output
	consensus *consensus.Raft
}

// processTable 进程表,key: pid;value: pcb
type processTable struct {
	m         sync.Map
	callback  chan<- *model.Output
	consensus consensus.Raft
}

func newPT() *processTable {
	return &processTable{m: sync.Map{}}
}

func (p *processTable) put(pid model.Pid, pcb *pcb) {
	p.m.Store(pid, pcb)
}

func (p *processTable) get(pid model.Pid) (bool, *pcb) {
	val, ok := p.m.Load(pid)
	if ok {
		return true, val.(*pcb)
	}
	return false, nil
}
func (p *processTable) delete(pid model.Pid) {
	p.m.Delete(pid)
}

func (p *processTable) shutdown() error {
	return nil
}
