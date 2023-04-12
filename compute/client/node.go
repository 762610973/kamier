package client

import (
	"context"
	"fmt"

	"compute/api/proto/node"
	cfg "compute/config"
	zlog "compute/log"
	"compute/model"

	"github.com/imroc/req/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	registerPath = "/node/register"
	getNodePath  = "/node/get"
)

var nodeMap = make(map[string]string)

// RegisterNode 启动时注册本节点地址
func RegisterNode() error {
	_, err := req.R().SetBodyJsonMarshal(map[string]string{
		"name": cfg.Cfg.NodeName,
		"addr": fmt.Sprintf("%s:%s", cfg.Cfg.GrpcAddr, cfg.Cfg.GrpcPort),
	}).Post(fmt.Sprintf("%s%s", cfg.Cfg.Storage, registerPath))
	if err != nil {
		return err
	}
	zlog.Info("register node success")
	return nil
}

// GetHost 通过节点名从storage获取对应节点的地址，准备执行时获取节点信息，并将信息缓存下来
func GetHost(nodeName string) (string, error) {
	resp, err := req.
		R().
		AddQueryParams("org,").
		Get(cfg.Cfg.StorageUrl + getNodePath)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

// Ipc 请求参与运算的节点
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

// Start 发起启动计算的请求
func Start(nodeName string, funcId string, pid *model.Pid) error {
	clientConn, err := grpc.Dial(nodeMap[nodeName])
	if err != nil {
		zlog.Error(fmt.Sprintf("grpc dial %s failed", nodeName))
		return err
	}
	client := node.NewNodeServiceClient(clientConn)
	_, err = client.Start(context.Background(), &node.StartReq{
		FuncId: funcId,
		Pid: &node.Pid{
			NodeName: pid.NodeName,
			Serial:   pid.Serial,
		},
	})
	if err != nil {
		zlog.Error("grpc start failed", zap.Error(err))
		return err
	}
	zlog.Debug("grpc start success")
	return nil
}

// Fetch 从其他节点获取值
func Fetch(nodeName string, pid *model.Pid, targetName, sourceName string, step int64) ([]byte, error) {
	clientConn, err := grpc.Dial(nodeMap[nodeName])
	if err != nil {
		zlog.Error(fmt.Sprintf("grpc dial %s failed", nodeName))
		return nil, err
	}
	client := node.NewNodeServiceClient(clientConn)
	res, err := client.Fetch(context.Background(), &node.FetchReq{
		Pid: &node.Pid{
			NodeName: pid.NodeName,
			Serial:   pid.Serial,
		},
		TargetName: targetName,
		SourceName: sourceName,
		Step:       step,
	})
	if err != nil {
		zlog.Error("grpc fetch failed", zap.Error(err))
		return nil, err
	}
	zlog.Debug("grpc prepare success")
	return res.Res, nil
}

// Prepare 发起准备请求
func Prepare(nodeName string, members []string) error {
	clientConn, err := grpc.Dial(nodeMap[nodeName])
	if err != nil {
		zlog.Error(fmt.Sprintf("grpc dial %s failed", nodeName))
		return err
	}
	client := node.NewNodeServiceClient(clientConn)
	_, err = client.Prepare(context.Background(), &node.PrepareReq{
		Members: members,
	})
	if err != nil {
		zlog.Error("grpc prepare failed", zap.Error(err))
		return err
	}
	zlog.Debug("grpc prepare success")
	return nil
}
