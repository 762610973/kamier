package client

import (
	"compute/api/proto/node"
	zlog "compute/log"
	"compute/model"
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// storage.go 请求storage获取计算函数，生成本地的临时文件，挂载至容器内

func getFunc() error {
	return nil
}

func generateTempFile() error {
	return nil
}

func Ipc(nodeName string, pid *model.Pid, arg []byte) ([]byte, error) {
	clientConn, err := grpc.Dial(nodeMap[nodeName])
	if err != nil {
		zlog.Error(fmt.Sprintf("grpc dial %s failed", nodeName))
		return nil, err
	}
	client := node.NewNodeServiceClient(clientConn)
	res, err := client.Ipc(context.Background(), &node.IpcReq{
		Pid: &node.Pid{
			NodeName: pid.NodeName,
			Serial:   pid.Serial,
		},
		Arg: arg,
	})
	if err != nil {
		zlog.Error("grpc ipc failed", zap.Error(err))
		return nil, err
	}
	zlog.Debug("grpc ipc success")
	return res.Res, nil
}
