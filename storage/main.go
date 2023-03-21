package main

import (
	"log"
	"os"
	"os/signal"
	ctl "storage/controller"

	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
)

func main() {
	h := server.New(config.Option{F: func(c *config.Options) {
		c.Addr = ":1112"
		c.Network = "tcp"
	}})
	h.Use(recovery.Recovery())
	h.GET("/ping", ctl.Ping)

	err := h.Run()
	if err != nil {
		return
	}
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		log.Println("Shutdown Server ...")
	}()

}
