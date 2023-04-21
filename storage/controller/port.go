package controller

import (
	"context"
	"strconv"
	"sync"

	zlog "storage/log"
	"storage/model"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

// pid:num(单独维护,与nodeMap没有关系)
var portMap = sync.Map{}

// portNum 共识地址的端口后增加的数字
var portNum = atomic.Int64{}

// GetConsensusPortNum 获取共识端口要递增的数字
func GetConsensusPortNum(_ context.Context, c *app.RequestContext) {
	var p model.Pid
	if err := c.Bind(&p); err != nil {
		zlog.Error("bind pid object failed", zap.Error(err))
		model.ErrResponse(c, err)
		return
	}
	_, ok := portMap.Load(p)
	if !ok {
		//	不存在则store
		portMap.Store(p, portNum.Add(1))
	}
	v, _ := portMap.Load(p)
	n := strconv.FormatInt(v.(int64), 10)
	zlog.Info("return port add num:" + n)
	c.String(consts.StatusOK, n)
}

// DeleteConsensusPortNum 计算完成后,发起运算的节点来请求此方法进行删除
func DeleteConsensusPortNum(_ context.Context, c *app.RequestContext) {
	var p model.Pid
	if err := c.Bind(&p); err != nil {
		zlog.Error("bind pid object failed", zap.Error(err))
		model.ErrResponse(c, err)
		return
	}
	portNum.Dec()
	portMap.Delete(p)
	c.String(consts.StatusOK, model.Success)
}

// GetAllPortNum 返回所有信息
func GetAllPortNum(_ context.Context, c *app.RequestContext) {
	var ret [][]any
	portMap.Range(func(key, value any) bool {
		ret = append(ret, []any{key, value})
		return true
	})
	model.SuccessResponse(c, ret)
}
