package main

import (
	"container/client"
	"container/config"
	"container/env"
	"container/function"
	"container/log"
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

func main() {
	f, b := function.Fnm.Get(env.GetFnName())
	if b {
		f()
	}
}
