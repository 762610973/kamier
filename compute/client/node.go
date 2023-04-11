package client

import (
	"fmt"

	cfg "compute/config"
	zlog "compute/log"

	"github.com/imroc/req/v3"
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
