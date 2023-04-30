package controller

import (
	"net"

	"compute/api/proto/container"
	cfg "compute/config"
	"compute/core"
	zlog "compute/log"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RunGrpcContainerServer() error {
	listener, err := net.Listen("tcp", "localhost:"+cfg.Cfg.ContainerAddr)
	if err != nil {
		zlog.Error("listen tcp failed", zap.Error(err))
		return err
	}
	cServer := grpc.NewServer()
	container.RegisterContainerServiceServer(cServer, core.C)
	if err = cServer.Serve(listener); err != nil {
		zlog.Error("container server serve failed", zap.Error(err))
		return err
	}
	return nil
}
