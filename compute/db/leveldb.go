package db

import (
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

func StoreSerial(nodeName string, serial int) error {
	val := strconv.Itoa(serial)
	if err := db.Put([]byte(nodeName), []byte(val), nil); err != nil {
		zlog.Error("store serial failed", zap.Error(err))
		return err
	}
	return nil
}

func LoadSerial(nodeName string) (int, error) {
	serial, err := db.Get([]byte(nodeName), nil)
	if err != nil {
		//zlog.Error("load serial failed", zap.Error(err))
		return 0, err
	}
	return strconv.Atoi(string(serial))
}

func StoreOutput(pid model.Pid) error {

	return nil
}

func LoadOutput() *model.Output {
	return nil
}
