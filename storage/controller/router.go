package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Ping(ctx context.Context, c *app.RequestContext) {
	c.String(consts.StatusOK, "OK")
}
