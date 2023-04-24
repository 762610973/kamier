package main

import (
	"container/client"
	"container/config"
	"container/env"
	"container/function"
	"container/log"
)

func init() {
	config.InitConfig()
	log.InitLogger()
	_ = client.InitClient()
}

func main() {
	f, b := function.Fnm.Get(env.GetFnName())
	if b {
		f()
	}
}
