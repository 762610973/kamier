package db

import (
	cfg "compute/config"
	zlog "compute/log"
	"compute/model"
	"github.com/syndtr/goleveldb/leveldb"
	"go.uber.org/zap"
)

var db *leveldb.DB

const (
	Serial = "serial_"
)

func InitLeveldb() {
	var err error
	db, err = leveldb.OpenFile(cfg.Cfg.LevelDBPath, nil)
	if err != nil {
		zlog.Panic("open leveldb file failed, init leveldb failed", zap.Error(err))
	}
}

func StoreSerial() {
	//db.Put()
}

func LoadSerial() {

}

func StoreOutput() error {
	return nil
}

func LoadOutput() *model.Output {
	return nil
}
