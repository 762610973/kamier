package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"time"

	"compute/client"
	"compute/config"
	"compute/db"
	zlog "compute/log"
	"compute/server/controller"

	"go.uber.org/zap"
)

func init() {
	config.InitConfig()
	zlog.InitLogger()
	db.InitLeveldb()
	if err := client.RegisterNode(); err != nil {
		zlog.Panic("register node failed", zap.Error(err))
	}
}

func main() {
	h := controller.RunHttpServer()
	var err error
	go func() {
		zlog.Info("start http server")
		err = h.Run()
		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				zlog.Info("begin graceful shutdown...")
			} else {
				zlog.Error("run http server failed", zap.Error(err))
			}
		}
	}()
	errCh := make(chan error, 1)
	go func() {
		err = controller.RunGrpcNodeServer()
		if err != nil {
			errCh <- err
		}
	}()
	go func() {
		err = controller.RunGrpcContainerServer()
		if err != nil {
			errCh <- err
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, os.Kill)
	select {
	case <-signalCh:
		//优雅退出
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = h.Shutdown(ctx); err != nil {
			zlog.Error("graceful shutdown failed", zap.Error(err))
		}
		db.CloseDB()
		zlog.Info("graceful shutdown...")
	case e := <-errCh:
		zlog.Error("start grpc server failed", zap.Error(e))
	}
}
