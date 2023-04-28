package function

import (
	zlog "containers/log"
	"go.uber.org/zap"
	"sync"
)

var Fnm = FnMap{
	fn: make(map[string]func()),
	mu: sync.Mutex{},
}

type FnMap struct {
	fn map[string]func()
	mu sync.Mutex
}

func (f *FnMap) Put(fnName string, fn func()) {
	f.mu.Lock()
	zlog.Debug("put fn to fnMap")
	f.fn[fnName] = fn
	f.mu.Unlock()
}

func (f *FnMap) Get(fnName string) (func(), bool) {
	f.mu.Lock()
	defer f.mu.Unlock()
	function, ok := f.fn[fnName]
	if ok {
		zlog.Info("get function by name success")
		return function, ok
	} else {
		zlog.Error("can't get function", zap.String("fnName", fnName))
		return nil, false
	}
}
