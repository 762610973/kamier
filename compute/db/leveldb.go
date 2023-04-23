package db

import (
	"encoding/json"
	"strconv"

	cfg "compute/config"
	zlog "compute/log"
	"compute/model"

	"github.com/syndtr/goleveldb/leveldb"
	"go.uber.org/zap"
)

var db *leveldb.DB

func InitLeveldb() {
	var err error
	db, err = leveldb.OpenFile(cfg.Cfg.LevelDBPath, nil)
	if err != nil {
		zlog.Panic("open leveldb file failed, init leveldb failed", zap.Error(err))
	}
}

func CloseDB() {
	_ = db.Close()
	zlog.Info("close db")
}

func StoreSerial(nodeName string, serial int64) error {
	val := strconv.FormatInt(serial, 10)
	if err := db.Put([]byte(nodeName), []byte(val), nil); err != nil {
		return err
	}
	return nil
}

const (
	SerialNotExists   = -1
	SerialParseFailed = -2
)

// LoadSerial 从leveldb中加载序列号
func LoadSerial(nodeName string) (int64, error) {
	serial, err := db.Get([]byte(nodeName), nil)
	if err != nil {
		zlog.Error("load serial failed", zap.Error(err))
		return SerialNotExists, err
	}
	res, err := strconv.ParseInt(string(serial), 10, 64)
	if err != nil {
		zlog.Error("parse serial failed", zap.Error(err))
		return SerialParseFailed, err
	}
	return res, nil
}

// StoreOutput 保存结果
func StoreOutput(pid model.Pid, output model.Output) error {
	key, err := json.Marshal(pid)
	if err != nil {
		zlog.Error("marshal pid as key failed", zap.Error(err))
		return err
	}
	val, err := json.Marshal(output)
	if err != nil {
		zlog.Error("marshal output as value failed", zap.Error(err))
		return err
	}
	if err = db.Put(key, val, nil); err != nil {
		zlog.Error("db.put failed", zap.Error(err))
		return err
	}
	return nil
}

func LoadOutput(pid model.Pid) (*model.Output, error) {
	key, err := json.Marshal(pid)
	if err != nil {
		zlog.Error("marshal pid failed", zap.Error(err), zap.Any("pid", pid))
		return nil, err
	}
	output, err := db.Get(key, nil)
	if err != nil {
		zlog.Error("get output failed", zap.Error(err), zap.Any("key", pid))
		return nil, err
	}
	var res model.Output
	err = json.Unmarshal(output, &res)
	if err != nil {
		zlog.Error("unmarshal output failed", zap.Error(err))
		return nil, err
	}
	zlog.Info("load output success")
	return &res, nil
}
