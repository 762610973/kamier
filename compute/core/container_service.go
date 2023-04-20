package core

import (
	"compute/model"
	"context"
	"github.com/pkg/errors"

	"compute/api/proto/container"
)

// PrepareValue 容器内服务准备值,存放到prepareValue结构中
func (c *Core) PrepareValue(_ context.Context, req *container.PrepareReq) (*container.PrepareRes, error) {
	pid := model.Pid{
		NodeName: req.Pid.NodeName,
		Serial:   req.Pid.Serial,
	}
	p, ok := c.processTable.get(pid)
	if ok {
		p.prepared.prepareValue(req.Step, req.Value)
	} else {
		return nil, errors.New("not found process")
	}
}

// FetchValue 容器内的服务获取其他节点内容器里的值
func (c *Core) FetchValue(_ context.Context, req *container.FetchReq) (*container.FetchRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Core) AppendValue(_ context.Context, req *container.AppendReq) (*container.AppendRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Core) WatchValue(_ context.Context, req *container.WatchReq) (*container.WatchRes, error) {
	//TODO implement me
	panic("implement me")
}