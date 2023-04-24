package function

import "sync"

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
	f.fn[fnName] = fn
	f.mu.Unlock()
}

func (f *FnMap) Get(fnName string) (func(), bool) {
	f.mu.Lock()
	function, ok := f.fn[fnName]
	if ok {
		return function, ok
	} else {
		return nil, false
	}
}
