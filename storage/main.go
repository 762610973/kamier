package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	cfg "storage/config"
	ctl "storage/controller"
	"storage/log"

	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
)

func init() {
	cfg.InitConfig()
	log.InitLogger()
}

func main() {
	h := server.New(config.Option{F: func(c *config.Options) {
		c.Addr = fmt.Sprintf(":%s", cfg.Cfg.NetWork.HttpPort)
		c.Network = "tcp"
	}})
	h.Use(recovery.Recovery())
	h.GET("/ping", ctl.Ping)
	// 公共函数

	h.POST("/function/add", ctl.AddFunc)
	h.GET("/function/get/", ctl.GetFunc)
	h.GET("/function/getAllFunc", ctl.GetAllFunc)
	h.DELETE("/function/delete/", ctl.DeleteFunc)
	h.PUT("/function/update/", ctl.UpdateFunc)

	// 公共数据
	h.POST("/data/add", ctl.AddData)
	h.GET("/data/get/", ctl.GetData)
	h.GET("/data/getAllFunc", ctl.GetAllData)
	h.DELETE("/data/delete/", ctl.DeleteData)
	h.PUT("/data/update/", ctl.UpdateData)

	// 节点
	h.POST("/node/add", ctl.RegisterNode)
	h.GET("/node/get/", ctl.GetNode)
	h.DELETE("/node/delete/", ctl.DeleteNode)
	h.POST("/node/update/", ctl.UpdateNode)
	h.GET("/node/getAllNode", ctl.GetAllNode)

	err := h.Run()
	if err != nil {
		return
	}
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		_ = h.Shutdown(context.Background())
	}()
}
