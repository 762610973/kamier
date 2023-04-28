package process

import (
	"encoding/json"

	"containers/client"
	"containers/env"
	zlog "containers/log"

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
		err := client.AppendValue(env.Pid, completeType, []byte("1"))
		if err != nil {
			zlog.Error("append value failed", zap.Error(err))
		}
	} else {
		types, value, err := client.WatchValue(env.Pid, node.Name)
		if err != nil {
			zlog.Error("watch value failed", zap.Error(err))
		}
		if types == completeType && string(value) == "1" {
			zlog.Info("watch value success")
		}
	}
label:
	return Function{
		node: node,
		fn:   fn,
	}
}

var step int64 = 1

func (f *Function) ComputeCallAt(node Node) (int, error) {
	if env.GetMembersLength() == "1" {
		return f.fn(), nil
	}
	step = (step + 1) / 2
	defer func() {
		step++
	}()
	if node.IsSelf() {
		data, err := json.Marshal(f.fn())
		if err != nil {
			zlog.Error("marshal fn() failed", zap.Error(err))
			return -1, err
		}
		err = client.PrepareValue(env.Pid, step, data)
		if err != nil {
			zlog.Error("prepare value failed", zap.Error(err))
		}
		return f.fn(), nil
	} else {
		value, err := client.FetchValue(env.Pid, node.Name, env.GetSelfName(), 1)
		if err != nil {
			zlog.Error("fetch value failed", zap.Error(err))
			return -1, err
		}
		var res int
		err = json.Unmarshal(value, &res)
		if err != nil {
			zlog.Error("unmarshal value failed", zap.Error(err))
			return -1, err
		}
		return res, nil
	}
}
