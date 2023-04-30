package controller

import (
	"net"

	"compute/api/proto/node"
	cfg "compute/config"
	"compute/core"
	zlog "compute/log"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RunGrpcNodeServer 启动节点间grpcServer
func RunGrpcNodeServer() (err error) {
	zlog.Debug("run grpc node server")
	listener, err := net.Listen("tcp", ":"+cfg.Cfg.GrpcPort)
	if err != nil {
		zlog.Error("net.Listen failed", zap.Error(err))
		return err
	}
	gServer := grpc.NewServer()
	node.RegisterNodeServiceServer(gServer, core.C)
	if err = gServer.Serve(listener); err != nil {
		zlog.Error("run grpc node server failed", zap.Error(err))
		return err
	}
	return
}
