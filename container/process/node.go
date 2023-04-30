package process

import (
	"encoding/json"

	"container/client"
	"container/env"
	zlog "container/log"

	"go.uber.org/zap"
)

type Node struct {
	Name string
}

// IsSelf 判断该节点是否是本节点
func (n Node) IsSelf() bool {
	return n.Name == env.GetSelfName()
}

type Function struct {
	node Node
	fn   func() int
}

const (
	completeType = "complete"
)

func CallAt(node Node, fn func() int) Function {
	if env.GetMembersLength() == "1" {
		goto label
	}
	if node.IsSelf() {
		err := client.AppendValue(env.Pid, completeType, []byte("1"), node.Name)
		if err != nil {
			zlog.Error("[Call At] append value failed", zap.Error(err))
		}
		zlog.Debug("[Call At] append value success")
	} else {
		types, value, err := client.WatchValue(env.Pid, node.Name)
		if err != nil {
			zlog.Error("[Call At] watch value failed", zap.Error(err))
		}
		if types == completeType && string(value) == "1" {
			zlog.Info("[Call At] watch value success")
		}
	}
label:
	return Function{
		node: node,
		fn:   fn,
	}
}

var step int64 = 1

func (f *Function) ComputeCallAt(targetNode Node) (int, error) {
	if env.GetMembersLength() == "1" {
		return f.fn(), nil
	}
	step = (step + 1) / 2
	defer func() {
		step++
	}()
	// 本机数据
	if f.node.IsSelf() {
		// 本节点运算,直接返回
		if targetNode.IsSelf() {
			err := client.AppendValue(env.Pid, completeType, []byte("1"), targetNode.Name)
			if err != nil {
				zlog.Error("append value failed", zap.Error(err))
				return -1, nil
			}
			zlog.Info("[ComputeCallAt] append value success")
			return f.fn(), nil
		} else {
			// 本机数据在其他节点运算
			data, err := json.Marshal(f.fn())
			if err != nil {
				zlog.Error("marshal fn() failed", zap.Error(err))
				return -1, err
			}
			// 准备值,并且等待对方完成
			err = client.PrepareValue(env.Pid, step, data)
			if err != nil {
				zlog.Error("prepare value failed", zap.Error(err))
				return -1, nil
			}
			_, _, err = client.WatchValue(env.Pid, targetNode.Name)
			if err != nil {
				zlog.Error("watch value failed", zap.Error(err))
				return -1, err
			}
			return 0, nil
		}
	} else {
		//! 非本机数据
		if targetNode.IsSelf() {
			// 非本机数据,在本节点运算
			// 从目标节点获取值
			zlog.Debug("fetch value", zap.String("target", f.node.Name), zap.String("source", env.GetSelfName()))
			value, err := client.FetchValue(env.Pid, f.node.Name, env.GetSelfName(), step)
			if err != nil {
				zlog.Error("[ComputeCallAt] fetch value failed", zap.Error(err))
				return -1, err
			}
			var res int
			err = json.Unmarshal(value, &res)
			if err != nil {
				zlog.Error("[ComputeCallAt] unmarshal value failed", zap.Error(err))
				return -1, err
			}
			// 获取到值之后反序列化,然后添加一个完成标记
			err = client.AppendValue(env.Pid, completeType, []byte("1"), env.GetSelfName())
			if err != nil {
				zlog.Error("[ComputeCallAt] append value failed", zap.Error(err))
				return -1, nil
			}
			return res, nil
		} else {
			_, _, err := client.WatchValue(env.Pid, targetNode.Name)
			if err != nil {
				zlog.Error("非本机数据在其他节点运算,watch complete failed", zap.Error(err))
				return -1, nil
			} else {
				return 0, nil
			}
		}
	}

}
