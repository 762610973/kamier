package model

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Request struct {
	FunctionId string   `json:"functionId"`
	Members    []string `json:"members" vd:"len($)>0"`
}

const (
	Success = "Success"
	Err     = "Failed"
)

type Response struct {
	Data any    `json:"data,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

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
