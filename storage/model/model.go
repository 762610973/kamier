package model

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Function struct {
	Name        string `bson:"name" json:"name"`
	Id          string `bson:"_id" json:"id"`
	Description string `bson:"description" json:"description"`
	Content     string `bson:"content" json:"content"`
}

type Data struct {
	Name    string `bson:"name" json:"name"`
	Id      string `bson:"_id" json:"id"`
	Content string `bson:"content" json:"content"`
}

type Node struct {
	Name string `bson:"name" json:"name"`
	Id   string `bson:"_id" json:"id"`
	Addr string `bson:"addr" json:"addr"`
}

type Response struct {
	Data any    `json:"data,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

const (
	Err     = "Failed"
	Success = "Success"
)

func ErrResponse(c *app.RequestContext, err error) {
	c.JSON(consts.StatusOK, Response{
		Data: err.Error(),
		Msg:  Err,
	})
}
func SuccessResponse(c *app.RequestContext, data any) {
	if data == nil {
		c.JSON(consts.StatusOK, Response{Msg: Success})
	} else {
		c.JSON(consts.StatusOK, Response{
			Data: data,
			Msg:  Success,
		})
	}
}
