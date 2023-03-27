package main

import (
	"compute/config"
	"compute/core"
	"compute/db"
	zlog "compute/log"
	"compute/server/controller"
	"context"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"strings"
	"time"
)

func init() {
	config.InitConfig()
	zlog.InitLogger()
	db.InitLeveldb()
}

func main() {
	c := core.NewCore()
	h := controller.RunHttpServer(c)
	go func() {
		go func() {
			zlog.Info("start http server")
			err := h.Run()
			if err != nil {
				if strings.Contains(err.Error(), "use of closed network connection") {
					zlog.Info("begin graceful shutdown...")
				} else {
					zlog.Error("run http server failed", zap.Error(err))
				}
			}
		}()
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, os.Kill)
	select {
	case <-signalCh:
		//优雅退出
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := h.Shutdown(ctx)
		if err != nil {
			zlog.Error("graceful shutdown failed", zap.Error(err))
		}
		zlog.Info("graceful shutdown...")
	}
}
