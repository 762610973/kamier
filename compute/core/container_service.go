package core

import (
	"context"

	"compute/api/proto/container"
)

func (c *Core) PrepareValue(_ context.Context, req *container.PrepareReq) (*container.PrepareRes, error) {
	//TODO implement me
	panic("implement me")
}

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
