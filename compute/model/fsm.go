package model

// fsm.go

const (
	Value    = "Value"
	Complete = "Complete"
	Finish   = "Finish"
	Error    = "Error"
)

// ConsensusValue 存放于共识队列上的值
type ConsensusValue struct {
	Type  string
	Value []byte
}

type ConsensusReq struct {
	// NodeName 节点名
	NodeName string
	// Serial 序列号
	Serial int64
	// Value 要添加的值
	Value ConsensusValue
}
