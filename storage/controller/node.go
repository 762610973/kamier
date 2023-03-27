package controller

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"storage/db"
	zlog "storage/log"
	"storage/model"
)

func Ping(_ context.Context, c *app.RequestContext) {
	c.String(consts.StatusOK, "OK")
}

func RegisterNode(_ context.Context, c *app.RequestContext) {
	var n model.Node
	err := c.Bind(&n)
	if err != nil {
		zlog.Error("RegisterNode bind object failed", zap.Error(err))
		model.ErrResponse(c, err)
	}
	err = db.InsertDocument(db.Node, n, n.Id)
	if err != nil {
		model.ErrResponse(c, err)
	}
	model.SuccessResponse(c, nil)
}
func GetNode(_ context.Context, c *app.RequestContext) {
	id := c.Query("id")
	err, data := db.FindDocument(db.Node, bson.M{db.ID: id})
	if err != nil {
		model.ErrResponse(c, err)
	} else {
		zlog.Debug(fmt.Sprintf("Get Node by %s success", id))
		model.SuccessResponse(c, data)
	}
}
func DeleteNode(_ context.Context, c *app.RequestContext) {
	id := c.Query("id")
	err := db.DeleteDocument(db.Node, bson.M{db.ID: id})
	if err != nil {
		model.ErrResponse(c, err)
	} else {
		zlog.Debug(fmt.Sprintf("Get Node by %s success", id))
		model.SuccessResponse(c, nil)
	}
}
func UpdateNode(_ context.Context, c *app.RequestContext) {
	var f model.Node
	err := c.Bind(&f)
	if err != nil {
		zlog.Error("bind node object failed", zap.Error(err))
		model.ErrResponse(c, err)
	} else {
		err := db.UpdateDocument(db.Node, f.Id, f)
		if err != nil {
			zlog.Error("update node failed", zap.Error(err))
			model.ErrResponse(c, err)
		} else {
			model.SuccessResponse(c, nil)
		}
	}
}
func GetAllNode(_ context.Context, c *app.RequestContext) {
	err, res := db.FindAllDocument(db.Node)
	if err != nil {
		model.ErrResponse(c, err)
	} else {
		zlog.Debug("Get All Node success")
		model.SuccessResponse(c, res)
	}
}
