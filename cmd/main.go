package main

import (
	"github.com/762610973/kamier/config"
	"github.com/762610973/kamier/log"
)

func main() {
	config.InitConfig()
	log.InitLogger()
	log.Zlog.Info("testInfo")
}
