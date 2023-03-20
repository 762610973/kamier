package main

import (
	"github.com/762610973/kamier/config"
	"github.com/762610973/kamier/log"
)

func init() {
	config.InitConfig()
	log.InitLogger()

}

func main() {
	log.Zlog.Info("test color")
	log.Zlog.Error("test color")
	log.Zlog.Warn("test color")
	log.Zlog.Debug("test debug")
}
