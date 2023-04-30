package main

import (
	"fmt"

	"container/client"
	"container/config"
	"container/env"
	"container/log"
	"container/process"

	"go.uber.org/zap"
)

func init() {
	config.InitConfig()
	log.InitLogger()
	env.InitPid()
	err := client.InitClient()
	if err != nil {
		log.Error("init client failed", zap.Error(err))
	}
}

func fnN1() int {
	return 7
}
func fnN2() int {
	return 3
}
func main() {
	n1 := process.Node{Name: "org1"}
	n2 := process.Node{Name: "org2"}
	f1 := process.CallAt(n1, fnN1)
	f2 := process.CallAt(n2, fnN2)
	res1, err := f1.ComputeCallAt(n1)
	if err != nil {
		log.Error("compute call at n1 failed", zap.Error(err))
		return
	}
	res2, err := f2.ComputeCallAt(n1)
	if err != nil {
		log.Error("compute call at n2 failed", zap.Error(err))
		return
	}
	fmt.Print(res1 + res2)
}
