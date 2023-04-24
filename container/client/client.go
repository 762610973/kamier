package client

import (
	"context"
	"go.uber.org/zap"

	"container/env"
	zlog "container/log"
	"container/proto/container"

	"google.golang.org/grpc"
)

var client container.ContainerServiceClient

func InitClient() (err error) {
	var clientConn *grpc.ClientConn
	clientConn, err = grpc.Dial(env.GetHostIp())
	if err != nil {
		zlog.Error("grpc dial failed", zap.Error(err))
		return err
	}
	client = container.NewContainerServiceClient(clientConn)
	return nil
}

func PrepareValue(pid env.PID, step int64, value []byte) error {
	_, err := client.PrepareValue(context.Background(), &container.PrepareReq{
		Pid: &container.Pid{
			NodeName: pid.NodeName,
			Serial:   pid.Serial,
		},
		Step:  step,
		Value: value,
	})
	if err != nil {
		zlog.Error("prepare value failed", zap.Error(err))
		return err
	}
	return nil
}

func FetchValue(pid env.PID, targetName, sourceName string, step int64) ([]byte, error) {
	res, err := client.FetchValue(context.Background(), &container.FetchReq{
		Pid: &container.Pid{
			NodeName: pid.NodeName,
			Serial:   pid.Serial,
		},
		TargetName: targetName,
		SourceName: sourceName,
		Step:       step,
	})
	if err != nil {
		zlog.Error("fetch value failed", zap.Error(err))
		return nil, err
	}
	zlog.Error("fetch value success", zap.String("res->", res.String()))
	return res.Res, nil
}
