package controller

import (
	"compute/core"
	"context"

	cfg "compute/config"
	zlog "compute/log"
	"compute/model"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"go.uber.org/zap"
)

func RunHttpServer() *server.Hertz {
	h := server.New(config.Option{F: func(o *config.Options) {
		o.Addr = ":" + cfg.Cfg.NetWork.HttpPort
		o.DisablePrintRoute = true
	}})
	h.Use(recovery.Recovery())
	h.POST("/syncCompute", syncCompute)
	h.POST("/asyncCompute", asyncCompute)
	h.POST("/getOutput", getOutput)
	return h
}

func syncCompute(_ context.Context, c *app.RequestContext) {
	var r model.Request
	var err error
	err = c.BindAndValidate(&r)
	if err != nil {
		zlog.Error("bindAndValidate failed", zap.Error(err))
		model.ErrResponse(c, err)
		return
	}
	output, err := core.SyncCompute(r)
	if err != nil {
		zlog.Error("sync compute failed", zap.Error(err))
		model.ErrResponse(c, err)
		return
	}
	zlog.Info("sync compute success")
	model.SuccessResponse(c, output)
}

func asyncCompute(_ context.Context, c *app.RequestContext) {
	var r model.Request
	var err error
	err = c.BindAndValidate(&r)
	if err != nil {
		zlog.Error("bindAndValidate failed", zap.Error(err))
		model.ErrResponse(c, err)
		return
	}
	pid, err := service.ASyncCompute(r)
	if err != nil {
		zlog.Error("async compute failed", zap.Error(err))
		model.ErrResponse(c, err)
		return
	}
	zlog.Info("async compute start success")
	model.SuccessResponse(c, pid)
}

func getOutput(_ context.Context, c *app.RequestContext) {
	var p model.Pid
	var err error
	err = c.Bind(&p)
	if err != nil {
		zlog.Error("bind failed", zap.Error(err))
		return
	}
	output, err := service.GetOutput(p)
	if err != nil {
		zlog.Error("get output by pid failed", zap.Error(err), zap.Any("pid", p))
		model.ErrResponse(c, err)
	}
	zlog.Info("get output success")
	model.SuccessResponse(c, output)
}
