package model

// Pid 每次计算的唯一标识: 名称+计算序列号
type Pid struct {
	NodeName string // 节点名称
	Serial   int64  // 节点本次计算的唯一序列号
}

type Output struct {
	StdOut string `json:"output,omitempty"`
	StdErr string `json:"err,omitempty"`
}
