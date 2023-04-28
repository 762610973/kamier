package exec

import (
	"containers/env"
	"containers/function"
	zlog "containers/log"
	"containers/process"
	"fmt"
	"go.uber.org/zap"
)

func init() {
	zlog.Info("init compute file, put func to function map")
	function.Fnm.Put(env.GetFnName(), do)
}

func fnN1() int {
	return 7
}
func fnN2() int {
	return 3
}
func do() {
	n1 := process.Node{Name: "node1"}
	n2 := process.Node{Name: "node2"}
	f1 := process.CallAt(n1, fnN1)
	f2 := process.CallAt(n2, fnN2)
	res1, err := f1.ComputeCallAt(n1)
	if err != nil {
		zlog.Error("compute call at n1 failed", zap.Error(err))
		return
	}
	res2, err := f2.ComputeCallAt(n1)
	if err != nil {
		zlog.Error("compute call at n2 failed", zap.Error(err))
		return
	}
	fmt.Print(res1 + res2)
}
