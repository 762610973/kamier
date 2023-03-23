package main

import (
	"compute/config"
	"compute/core"
	zlog "compute/log"
	"compute/server/controller"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

func init() {
	config.InitConfig()
	zlog.InitLogger()
}

func main() {
	var err error
	c := core.NewCore()
	var h *server.Hertz
	err = controller.RunHttpServer(c, h)
	if err != nil {
		zlog.Error("run http server failed", zap.Error(err))
		return
	}
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		status := <-quit
		fmt.Println("status: ", status)
		zlog.Info("shutdown all server")
		_ = h.Shutdown(context.Background())
	}()
}
