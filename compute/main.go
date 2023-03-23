package main

import (
	"compute/config"
	"compute/core"
	zlog "compute/log"
	"compute/server/controller"
	"context"
	"github.com/cloudwego/hertz/pkg/app/server"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"time"
)

func init() {
	config.InitConfig()
	zlog.InitLogger()
}

func main() {
	var err error
	c := core.NewCore()
	var h *server.Hertz
	go func() {
		err = controller.RunHttpServer(c, h)
		if err != nil {
			zlog.Error("start http server failed", zap.Error(err))
			os.Exit(1)
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
}
