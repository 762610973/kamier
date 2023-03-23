package core

import (
	"compute/model"
	"sync"
)

type Core struct {
	processTable *model.ProcessTable
	lock         sync.Mutex
	watermarks   sync.Map
	prepareTable sync.Map
}

func NewCore() *Core {
	return &Core{
		processTable: model.NewPT(),
		lock:         sync.Mutex{},
		watermarks:   sync.Map{},
		prepareTable: sync.Map{},
	}
}
