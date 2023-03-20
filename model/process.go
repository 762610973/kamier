package model

// Pid 每次计算的唯一标识: 名称+计算序列号
type Pid struct {
	OrgName string // 节点名称
	Serial  int    // 节点本次计算的唯一序列号
}

type Output struct {
}
