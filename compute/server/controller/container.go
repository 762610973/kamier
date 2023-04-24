package controller

import (
	"net"
	"os"

	"compute/api/proto/container"
	cfg "compute/config"
	"compute/core"
	zlog "compute/log"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RunGrpcContainerServer(cServer *grpc.Server) (err error) {
	if _, err = os.Stat(cfg.Cfg.SocketPath); err == nil {
		if err = os.Remove(cfg.Cfg.SocketPath); err != nil {
			zlog.Error("remove grpc listen socket file failed", zap.Error(err))
			return err
		}
	}
	zlog.Debug("delete socket file")
	listener, err := net.ListenUnix("unix", &net.UnixAddr{
		Name: cfg.Cfg.SocketPath,
		Net:  "unix",
	})
	if err != nil {
		zlog.Error("net.Listen failed", zap.Error(err))
		return err
	}
	if err = os.Chmod(cfg.Cfg.SocketPath, 0777); err != nil {
		zlog.Error("socket file elevate privileges failed", zap.Error(err))
		return err
	}
	container.RegisterContainerServiceServer(cServer, core.C)
	if err = cServer.Serve(listener); err != nil {
		zlog.Error("container server serve failed", zap.Error(err))
		return err
	}
	return
}
