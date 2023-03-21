package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Ping(ctx context.Context, c *app.RequestContext) {
	c.String(consts.StatusOK, "OK")
}

func GetFunc(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	c.String(consts.StatusOK, id)
}

func GetAllFunc(ctx context.Context, c *app.RequestContext) {

}
func AddFunc(ctx context.Context, c *app.RequestContext) {

}
func DeleteFunc(ctx context.Context, c *app.RequestContext) {

}
func UpdateFunc(ctx context.Context, c *app.RequestContext) {

}

func AddData(ctx context.Context, c *app.RequestContext) {

}
func GetData(ctx context.Context, c *app.RequestContext) {

}

func GetAllData(ctx context.Context, c *app.RequestContext) {

}
func DeleteData(ctx context.Context, c *app.RequestContext) {

}
func UpdateData(ctx context.Context, c *app.RequestContext) {

}

func RegisterNode(ctx context.Context, c *app.RequestContext) {

}

func GetNode(ctx context.Context, c *app.RequestContext) {

}
func DeleteNode(ctx context.Context, c *app.RequestContext) {

}
func UpdateNode(ctx context.Context, c *app.RequestContext) {

}
func GetAllNode(ctx context.Context, c *app.RequestContext) {

}
