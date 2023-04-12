package consensus

import (
	"encoding/json"
	"io"

	zlog "compute/log"
	"compute/model"

	"github.com/hashicorp/raft"
	"go.uber.org/zap"
)

// raft.go

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

func (f *fsm) Apply(l *raft.Log) any {
	var tmp model.ConsensusReq
	err := json.Unmarshal(l.Data, &tmp)
	if err != nil {
		zlog.Error("applied data decode failed", zap.Error(err))
		return err
	}
	f.pushValue(&model.ConsensusReq{
		Site:   tmp.Site,
		Serial: tmp.Serial,
		Value:  tmp.Value,
	})
	return nil
}

func (f *fsm) pushValue(req *model.ConsensusReq) {
	zlog.Debug("push data to fsm", zap.Int32("[site]", req.Site), zap.Any("[value]", req.Value))
	f.Queue = append(f.Queue, value{
		Site:  req.Site,
		Value: req.Value,
	})
	if f.Watch != nil {
		zlog.Debug("There is a node waiting for a value")
		// 此时正在等待的节点等到了对应节点
		if f.Watch.site == req.Site {
			// 将值发送到正在等待的节点的channel中
			f.Watch.ch <- req.Value
			// 更新pointer
			/*
					siteA SiteB
					A     A
					A     A
					A     A
					B     B
				* time: ->->->->->->
				* siteA: A->A->A->B
				* SiteB:    A->A->A->B
				! 节点A率先放置了三个完成标记,等待B的完成标记
				! 节点B依次获取A的三个完成标记,Pointer++,然后放置B的完成标记,此时pointer就可以更新到当前位置了
			*/
			// TODO: 此处f.pointer可以是++
			//f.Pointer++
			f.Pointer = len(f.Queue)
			f.Watch = nil
		}
	}
}

func (f *fsm) watchValue(targetSite int32, waiter chan model.ConsensusValue) {
	zlog.Debug("Watch value from", zap.Int32("[targetSite]", targetSite))
	for i := f.Pointer; i < len(f.Queue); i++ {
		if f.Queue[i].Site == targetSite {
			// 消耗一个完成标记
			f.Pointer++
			waiter <- f.Queue[i].Value
			return
		}
	}
	f.Watch = &watch{
		site: targetSite,
		// 传递waiter
		ch: waiter,
	}
}

// Restore 读取本地数据恢复快照
func (f *fsm) Restore(snapshot io.ReadCloser) error {
	var r fsm
	err := json.NewDecoder(snapshot).Decode(&r)
	if err != nil {
		zlog.Error("restore failed", zap.Error(err))
		return err
	}
	zlog.Debug("restore success")
	return nil
}

// Release 完成快照后的回调函数
func (f *fsm) Release() {}

// Persist 持久化必要信息
func (f *fsm) Persist(sink raft.SnapshotSink) error {
	data, err := json.Marshal(f)
	if err != nil {
		zlog.Error("encode snapshot failed", zap.Error(err))
		return err
	}
	_, err = sink.Write(data)
	if err != nil {
		zlog.Error("sink.Write failed", zap.Error(err))
		return err
	}
	if err = sink.Close(); err != nil {
		zlog.Error("sink.close failed", zap.Error(err))
		return err
	}
	return nil
}

// Snapshot 快照包括整个fsm结构
func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	return f, nil
}
