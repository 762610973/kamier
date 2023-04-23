package core

import (
	"compute/client"
	zlog "compute/log"
	"compute/model"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"go.uber.org/zap"

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
		return &container.PrepareRes{}, nil
	} else {
		return nil, errors.New("not found process")
	}
}

// FetchValue 容器内的服务获取其他节点内容器里的值
func (c *Core) FetchValue(_ context.Context, req *container.FetchReq) (*container.FetchRes, error) {
	pid := model.Pid{
		NodeName: req.Pid.NodeName,
		Serial:   req.Pid.Serial,
	}
	res, err := client.Fetch(pid, req.TargetName, req.SourceName, req.Step)
	if err != nil {
		zlog.Error("fetch value failed", zap.Error(err))
		return nil, err
	}
	return &container.FetchRes{Res: res}, nil
}

// AppendValue 容器内服务向共识队列中添加值
func (c *Core) AppendValue(_ context.Context, req *container.AppendReq) (*container.AppendRes, error) {
	pid := model.Pid{
		NodeName: req.Pid.NodeName,
		Serial:   req.Pid.Serial,
	}
	p, ok := c.processTable.get(pid)
	if !ok {
		return nil, errors.New("pid not found")
	} else {
		v := model.ConsensusValue{
			Type:  req.Type,
			Value: req.Value,
		}
		arg, err := json.Marshal(v)
		if err != nil {
			zlog.Error("marshal consensus value failed", zap.Error(err))
			return nil, err
		}
		p.consensus.Push(pid, arg)
		zlog.Debug("append value success")
		return &container.AppendRes{}, nil
	}

}

func (c *Core) WatchValue(_ context.Context, req *container.WatchReq) (*container.WatchRes, error) {
	pid := model.Pid{
		NodeName: req.Pid.NodeName,
		Serial:   req.Pid.Serial,
	}
	p, ok := c.processTable.get(pid)
	if !ok {
		return nil, errors.New("pid not found")
	}
	res := p.consensus.Watch(req.TargetName)
	zlog.Debug("watch value success")
	return &container.WatchRes{
		Type:  res.Type,
		Value: res.Value,
	}, nil

}
