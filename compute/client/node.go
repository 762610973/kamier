package client

import (
	cfg "compute/config"
	zlog "compute/log"
	"fmt"
	"github.com/imroc/req/v3"
)

// RegisterNode 启动时注册本节点地址
func RegisterNode() error {
	_, err := req.R().SetBodyJsonMarshal(map[string]string{
		"name": cfg.Cfg.OrgName,
		"addr": fmt.Sprintf("%s:%s", cfg.Cfg.GrpcAddr, cfg.Cfg.GrpcPort),
	}).Post(fmt.Sprintf("%s%s", cfg.Cfg.Storage, "/node/register"))
	if err != nil {
		return err
	}
	zlog.Info("register node success")
	return nil
}

// GetHost 通过节点名从storage获取对应节点的地址
func GetHost(orgName string) string {

	return ""
}
