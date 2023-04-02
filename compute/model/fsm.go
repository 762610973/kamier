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
	Value any
}

type ConsensusReq struct {
	// Site 节点号
	Site int32
	// Serial 序列号
	Serial int32
	// Value 要添加的值
	Value ConsensusValue
}
