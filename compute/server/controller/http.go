package controller

import (
	"compute/server/service"
	"context"

	cfg "compute/config"
	"compute/core"
	zlog "compute/log"
	"compute/model"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go.uber.org/zap"
)

type controller struct {
	*core.Core
}

func RunHttpServer(core *core.Core) *server.Hertz {
	h := server.New(config.Option{F: func(o *config.Options) {
		o.Addr = ":" + cfg.Cfg.NetWork.HttpPort
		o.DisablePrintRoute = true
	}})
	h.Use(recovery.Recovery())
	ctl := controller{core}
	h.POST("/syncCompute", ctl.SyncCompute)
	h.POST("/asyncCompute")
	h.POST("/getOutput")
	return h
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
	output, err := service.SyncCompute(r)
	if err != nil {
		zlog.Error("sync compute failed", zap.Error(err))
		model.ErrResponse(c, err)
	}
	zlog.Info("sync compute success")
	c.JSON(consts.StatusOK, "success")
}
