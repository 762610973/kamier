package core

import (
	"sync"

	"compute/model"
)

type Core struct {
	processTable *model.ProcessTable
	lock         sync.Mutex
	prepareTable sync.Map
}

func NewCore() *Core {
	return &Core{
		processTable: model.NewPT(),
		lock:         sync.Mutex{},
		prepareTable: sync.Map{},
	}
}
