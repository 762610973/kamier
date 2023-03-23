package controller

import (
	cfg "compute/config"
	"compute/core"
	zlog "compute/log"
	"compute/model"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type controller struct {
	*core.Core
}

func RunHttpServer(core *core.Core, h *server.Hertz) error {
	h = server.New(config.Option{F: func(o *config.Options) {
		o.Addr = ":" + cfg.Cfg.NetWork.HttpPort
	}})
	h.Use(recovery.Recovery())
	ctl := controller{core}
	h.POST("/syncCompute", ctl.SyncCompute)
	h.POST("/asyncCompute")
	h.POST("/getOutput")
	err := h.Run()
	if err != nil {
		zlog.Error("start http server failed", zap.Error(err))
		return err
	}
	return nil
}

func RunGrpcProcessServer(core *core.Core) (*grpc.Server, error) {
	return nil, nil
}

func (ctl *controller) SyncCompute(_ context.Context, c *app.RequestContext) {
	var r model.Request
	var err error
	//err = c.Bind(&r)
	err = c.BindAndValidate(&r)
	if err != nil {
		zlog.Error("bindAndValidate failed", zap.Error(err))
		model.ErrResponse(c, err)
		return
	}
	c.JSON(consts.StatusOK, "success")
}
