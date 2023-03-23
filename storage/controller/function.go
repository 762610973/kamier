package controller

import (
	"context"
	"fmt"

	"storage/db"
	zlog "storage/log"
	"storage/model"

	"github.com/cloudwego/hertz/pkg/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func GetFunc(_ context.Context, c *app.RequestContext) {
	id := c.Query("id")
	err, data := db.FindDocument(db.Function, bson.M{"_id": id})
	if err != nil {
		model.ErrResponse(c, err)
	} else {
		zlog.Debug(fmt.Sprintf("Get Func by %s success", id))
		model.SuccessResponse(c, data)
	}
}

func GetAllFunc(_ context.Context, c *app.RequestContext) {
	err, res := db.FindAllDocument(db.Function)
	if err != nil {
		model.ErrResponse(c, err)
	} else {
		zlog.Debug("Get All Func success")
		model.SuccessResponse(c, res)
	}
}

func AddFunc(_ context.Context, c *app.RequestContext) {
	var f model.Function
	err := c.Bind(&f)
	if err != nil {
		zlog.Error("AddFunc bind object failed", zap.Error(err))
		model.ErrResponse(c, err)
	}
	err = db.InsertDocument(db.Function, f)
	if err != nil {
		model.ErrResponse(c, err)
	}
	model.SuccessResponse(c, nil)
}
func DeleteFunc(_ context.Context, c *app.RequestContext) {
	id := c.Query("id")
	err := db.DeleteDocument(db.Function, bson.M{"_id": id})
	if err != nil {
		model.ErrResponse(c, err)
	} else {
		zlog.Debug(fmt.Sprintf("Get Function by %s success", id))
		model.SuccessResponse(c, nil)
	}
}

func UpdateFunc(_ context.Context, c *app.RequestContext) {
	var f model.Function
	err := c.Bind(&f)
	if err != nil {
		zlog.Error("bind func object failed", zap.Error(err))
		model.ErrResponse(c, err)
	} else {
		err := db.UpdateDocument(db.Function, f.Id, f)
		if err != nil {
			zlog.Error("update function failed", zap.Error(err))
			model.ErrResponse(c, err)
		} else {
			model.SuccessResponse(c, nil)
		}
	}
}
