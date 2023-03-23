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

func AddData(ctx context.Context, c *app.RequestContext) {
	var d model.Data
	err := c.Bind(&d)
	if err != nil {
		zlog.Error("AddData bind object failed", zap.Error(err))
		model.ErrResponse(c, err)
	}
	err = db.InsertDocument(db.Data, d)
	if err != nil {
		model.ErrResponse(c, err)
	} else {
		model.SuccessResponse(c, nil)
	}
}
func GetData(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	err, data := db.FindDocument(db.Data, bson.M{"_id": id})
	if err != nil {
		model.ErrResponse(c, err)
	} else {
		zlog.Debug(fmt.Sprintf("Get Data by %s success", id))
		model.SuccessResponse(c, data)
	}
}

func GetAllData(ctx context.Context, c *app.RequestContext) {
	err, res := db.FindAllDocument(db.Data)
	if err != nil {
		model.ErrResponse(c, err)
	} else {
		zlog.Debug("Get All Data success")
		model.SuccessResponse(c, res)
	}
}
func DeleteData(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	err := db.DeleteDocument(db.Data, bson.M{"_id": id})
	if err != nil {
		model.ErrResponse(c, err)
	} else {
		zlog.Debug(fmt.Sprintf("Get Data by %s success", id))
		model.SuccessResponse(c, nil)
	}
}
func UpdateData(ctx context.Context, c *app.RequestContext) {
	var f model.Data
	err := c.Bind(&f)
	if err != nil {
		zlog.Error("bind data object failed", zap.Error(err))
		model.ErrResponse(c, err)
	} else {
		err := db.UpdateDocument(db.Data, f.Id, f)
		if err != nil {
			zlog.Error("update data failed", zap.Error(err))
			model.ErrResponse(c, err)
		} else {
			model.SuccessResponse(c, nil)
		}
	}
}
