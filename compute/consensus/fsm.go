package consensus

import "compute/model"

// fsm 无需锁保护
type fsm struct {
	Queue   []value
	Pointer int
	Watch   *watch
}

func newFsm() *fsm {
	return &fsm{
		Queue:   make([]value, 0, 10),
		Pointer: 0,
		Watch:   nil,
	}
}

type watch struct {
	site int32
	ch   chan model.ConsensusValue
}

// 队列中的值,site表示放到队列当中的节点的节点号
// 这两个字段是要序列化的,必须可导出
type value struct {
	// 节点号
	Site int32
	// 队列中的值
	Value model.ConsensusValue
}
