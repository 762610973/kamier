package service

import (
	gclient "compute/client"
	"compute/config"
	"compute/db"
	zlog "compute/log"
	"compute/model"
	"context"
	"fmt"
	"go.uber.org/zap"
	"sync"
)

// SyncCompute 同步计算处理逻辑,发起准备,启动计算
func SyncCompute(req model.Request) (*model.Output, error) {
	// 本节点prepare
	pid, err := preparePid()
	if err != nil {
		return nil, err
	}
	err = allNodePrepare(req.Members)
	if err != nil {
		zlog.Error("all node prepare failed", zap.Error(err))
		return nil, err
	}
	allNodeStart(req.Members, req.FunctionId, pid)

	return nil, nil
}
func ASyncCompute(req model.Request) (*model.Pid, error) {
	return nil, nil
}
func GetOutput(pid model.Pid) (*model.Output, error) {
	return db.LoadOutput(pid)
}

var (
	mu = sync.Mutex{}
	// baseSerial 从1开始递增
	baseSerial = int64(1)
)

type done chan struct{}

// preparePid 准备执行计算,校验pid
func preparePid() (pid *model.Pid, err error) {
	var serial int64
	mu.Lock()
	defer mu.Unlock()
	serial, err = db.LoadSerial(config.Cfg.NodeName)
	if err != nil {
		// 序列号不存在
		if serial == db.SerialNotExists {
			err = db.StoreSerial(config.Cfg.NodeName, baseSerial)
			if err != nil {
				zlog.Error("store serial failed", zap.Error(err))
				return nil, err
			}
			serial = baseSerial
		}
		// 解析序列号失败
		if serial == db.SerialParseFailed {
			return nil, err
		}
	}
	// 序列号存在,则加1并存储
	serial = serial + 1
	err = db.StoreSerial(config.Cfg.NodeName, serial)
	if err != nil {
		zlog.Error("store serial failed", zap.Error(err))
		return nil, err
	}
	zlog.Info("preparePid success")
	return &model.Pid{
		NodeName: config.Cfg.NodeName,
		Serial:   serial,
	}, nil
}

// allNodePrepare 并发请求除本节点外的所有节点,其中一个节点准备失败则本次计算失败
func allNodePrepare(members []string) error {
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// doneCh 防止死锁
	doneCh := make([]done, len(members))
	for index, member := range members {
		doneCh[index] = make(done)
		go func(member string, d done) {
			if member == config.Cfg.NodeName {
				d <- struct{}{}
			} else {
				err = gclient.Prepare(member, members)
				if err != nil {
					zlog.Error(fmt.Sprint(member, "prepare failed"))
					cancel()
					return
				}
				d <- struct{}{}
			}
		}(member, doneCh[index])
	}
	for _, done := range doneCh {
		select {
		case <-done:
			continue
		case <-ctx.Done():
			return err
		}
	}
	zlog.Info("all node start success")
	return nil
}

func allNodeStart(members []string, funcId string, pid *model.Pid) {
	var wg sync.WaitGroup
	var err error
	for _, member := range members {
		wg.Add(1)
		go func(member string) {
			defer wg.Done()
			if member != config.Cfg.NodeName {
				err = gclient.Start(member, funcId, pid)
				if err != nil {
					zlog.Error("start failed", zap.Error(err))
				}
			}
		}(member)
	}
	wg.Wait()
	zlog.Debug("all node begin start")
}
