package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"go.uber.org/zap"
	"os"
	"os/signal"
	cfg "storage/config"
	ctl "storage/controller"
	"storage/db"
	zlog "storage/log"
	"strings"
	"time"
)

func init() {
	cfg.InitConfig()
	zlog.InitLogger()
	db.InitMongoDB()
}
func main() {
	h := server.Default(config.Option{F: func(c *config.Options) {
		c.Addr = fmt.Sprintf(":%s", cfg.Cfg.HttpPort)
		c.Network = "tcp"
		c.DisablePrintRoute = true
	}})
	h.GET("/ping", ctl.Ping)
	// 公共函数

	h.POST("/function/add", ctl.AddFunc)
	h.GET("/function/get", ctl.GetFunc)
	h.GET("/function/getAllFunc", ctl.GetAllFunc)
	h.DELETE("/function/delete/", ctl.DeleteFunc)
	h.PUT("/function/update/", ctl.UpdateFunc)

	// 公共数据
	h.POST("/data/add", ctl.AddData)
	h.GET("/data/get", ctl.GetData)
	h.GET("/data/getAllFunc", ctl.GetAllData)
	h.DELETE("/data/delete/", ctl.DeleteData)
	h.PUT("/data/update/", ctl.UpdateData)

	// 节点
	h.POST("/node/add", ctl.RegisterNode)
	h.GET("/node/get/", ctl.GetNode)
	h.DELETE("/node/delete/", ctl.DeleteNode)
	h.POST("/node/update/", ctl.UpdateNode)
	h.GET("/node/getAllNode", ctl.GetAllNode)
	go func() {
		err := h.Run()
		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				zlog.Info("begin graceful shutdown...")
			} else {
				zlog.Error("run http server failed", zap.Error(err))
			}
		}
	}()
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, os.Kill)
	select {
	case <-signalCh:
		//优雅退出
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := h.Shutdown(ctx)
		zlog.Info("graceful shutdown...")
		if err != nil {
			return
		}
	}
	zlog.Info("graceful shutdown...http server shutdown")
}
