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

func Ping(ctx context.Context, c *app.RequestContext) {
	c.String(consts.StatusOK, "OK")
}

func GetFunc(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	err, data := db.FindDocument(db.Function, bson.M{"_id": id})
	if err != nil {
		model.ErrResponse(c, err)
	} else {
		zlog.Debug(fmt.Sprintf("Get Func by %s success", id))
		model.SuccessResponse(c, data)
	}
}

func GetAllFunc(ctx context.Context, c *app.RequestContext) {
	err, res := db.FindAllDocument(db.Function)
	if err != nil {
		model.ErrResponse(c, err)
	} else {
		zlog.Debug("Get All Func success")
		model.SuccessResponse(c, res)
	}
}

func AddFunc(ctx context.Context, c *app.RequestContext) {
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
func DeleteFunc(ctx context.Context, c *app.RequestContext) {

}

func UpdateFunc(ctx context.Context, c *app.RequestContext) {

}

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

}
func UpdateData(ctx context.Context, c *app.RequestContext) {

}

func RegisterNode(ctx context.Context, c *app.RequestContext) {
	var n model.Node
	err := c.Bind(&n)
	if err != nil {
		zlog.Error("RegisterNode bind object failed", zap.Error(err))
		model.ErrResponse(c, err)
	}
	err = db.InsertDocument(db.Node, n)
	if err != nil {
		model.ErrResponse(c, err)
	}
	model.SuccessResponse(c, nil)
}

func GetNode(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	err, data := db.FindDocument(db.Node, bson.M{"_id": id})
	if err != nil {
		model.ErrResponse(c, err)
	} else {
		zlog.Debug(fmt.Sprintf("Get Node by %s success", id))
		model.SuccessResponse(c, data)
	}
}
func DeleteNode(ctx context.Context, c *app.RequestContext) {

}
func UpdateNode(ctx context.Context, c *app.RequestContext) {

}
func GetAllNode(ctx context.Context, c *app.RequestContext) {
	err, res := db.FindAllDocument(db.Node)
	if err != nil {
		model.ErrResponse(c, err)
	} else {
		zlog.Debug("Get All Node success")
		model.SuccessResponse(c, res)
	}
}
