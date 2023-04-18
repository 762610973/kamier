package core

import (
	"context"

	"compute/api/proto/node"
)

func (c *Core) Prepare(_ context.Context, req *node.PrepareReq) (*node.PrepareRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Core) Start(_ context.Context, req *node.StartReq) (*node.StartRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Core) Ipc(_ context.Context, req *node.IpcReq) (*node.IpcRes, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Core) Fetch(_ context.Context, req *node.FetchReq) (*node.FetchRes, error) {
	//TODO implement me
	panic("implement me")
}
