package core

import (
	gclient "compute/client"
	cfg "compute/config"
	zlog "compute/log"
	"compute/model"
	"context"
	"fmt"

	"compute/api/proto/node"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Prepare 参与运算的结点准备启动时校验本节点名即可
func (c *Core) Prepare(_ context.Context, req *node.PrepareReq) (*node.PrepareRes, error) {
	go func() {
		for _, member := range req.Members {
			go func(m string) {
				host, err := gclient.GetHost(m)
				if err != nil {
					zlog.Error("get host failed", zap.Error(err))
					return
				}
				gclient.Nodemap.Put(m, host)
			}(member)
		}
	}()
	var exist bool
	for _, member := range req.Members {
		if member == cfg.Cfg.NodeName {
			exist = true
			break
		}
	}
	if exist {
		zlog.Info("prepare exec complete")
		return &node.PrepareRes{}, nil
	} else {
		zlog.Error(fmt.Sprint(cfg.Cfg.NodeName, "not int compute members"))
		return nil, errors.New(fmt.Sprint(cfg.Cfg.NodeName, "not int compute members"))
	}
}

// Start 参与运算的节点开始启动,不需要准备工作,直接启动即可
func (c *Core) Start(_ context.Context, req *node.StartReq) (*node.StartRes, error) {
	errCh := make(chan error, 1)
	c.StartProcess(model.Pid{
		NodeName: req.Pid.NodeName,
		Serial:   req.Pid.Serial,
	}, req.FuncId, req.Members, nil, errCh)
	err := <-errCh
	if err != nil {
		zlog.Error("[slave] start process failed", zap.Error(err))
		return nil, err
	} else {
		zlog.Debug("[slave] start process success")
		return &node.StartRes{}, nil
	}

}

// Ipc 容器内程序向共识队列中添加至,通过此方法
func (c *Core) Ipc(_ context.Context, req *node.IpcReq) (*node.IpcRes, error) {
	pid := model.Pid{
		NodeName: req.NodeName,
		Serial:   req.Serial,
	}
	p, ok := c.processTable.get(pid)
	if !ok {
		zlog.Error(PidNotExistsErr)
		return nil, errors.New(PidNotExistsErr)
	}
	p.consensus.Push(pid, req.Arg)
	return &node.IpcRes{}, nil
}

// Fetch 其他节点从本节点获取值,调用fetchValue
func (c *Core) Fetch(_ context.Context, req *node.FetchReq) (*node.FetchRes, error) {
	pid := model.Pid{
		NodeName: req.Pid.NodeName,
		Serial:   req.Pid.Serial,
	}
	p, ok := c.processTable.get(pid)
	if ok {
		v := make(chan value, 1)
		p.prepared.fetchValue(req.Step, v)
		zlog.Debug("waiting channel to send value")
		res := <-v
		zlog.Debug("wait channel to send value")
		return &node.FetchRes{Res: res}, nil
	} else {
		zlog.Error("can't fetch value, pid not found")
		return nil, errors.New("can't fetch value")
	}
}
