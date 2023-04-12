package service

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
