package controller

import (
	"context"
	"sync"

	zlog "storage/log"
	"storage/model"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Ping(_ context.Context, c *app.RequestContext) {
	c.String(consts.StatusOK, "OK")
}

// name:addr
var nodeMap = sync.Map{}

// RegisterNode 启动时注册节点,执行完删除节点
func RegisterNode(_ context.Context, c *app.RequestContext) {
	var n model.Node
	err := c.Bind(&n)
	if err != nil {
		zlog.Error("RegisterNode bind object failed", zap.Error(err))
		model.ErrResponse(c, err)
		return
	}
	nodeMap.Store(n.Name, n.Addr)
	zlog.Info("register node success", zap.Any("node", n))
	model.SuccessResponse(c, nil)
}

// GetNode 获取节点名对应的地址
func GetNode(_ context.Context, c *app.RequestContext) {
	name := c.Query("name")
	value, ok := nodeMap.Load(name)
	if !ok {
		zlog.Error("not found node", zap.String("name", name))
		model.ErrResponse(c, errors.New("not found"))
	} else {
		zlog.Info("get node success")
		c.String(consts.StatusOK, value.(string))
	}
}

// DeleteNode 删除节点
func DeleteNode(_ context.Context, c *app.RequestContext) {
	name := c.Query("name")
	nodeMap.Delete(name)
	c.String(consts.StatusOK, model.Success)
}

// GetAllNode 获取所有节点的信息
func GetAllNode(_ context.Context, c *app.RequestContext) {
	var m map[string]string
	nodeMap.Range(func(key, value any) bool {
		m[key.(string)] = value.(string)
		return true
	})
	model.SuccessResponse(c, m)
}
