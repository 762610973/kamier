package db

import (
	"testing"

	"compute/config"
	zlog "compute/log"

	"github.com/stretchr/testify/assert"
)

func TestStoreSerial(t *testing.T) {
	config.InitConfig()
	zlog.InitLogger()
	InitLeveldb()
	err := StoreSerial("org1", 1)
	assert.Equal(t, nil, err)
	err = StoreSerial("org12", 12)
	assert.Equal(t, nil, err)
}

func TestLoadSerial(t *testing.T) {
	config.InitConfig()
	zlog.InitLogger()
	InitLeveldb()
	serial, err := LoadSerial("org1")
	assert.Equal(t, err, nil)
	assert.Equal(t, 1, serial)
	_, err = LoadSerial("org11")
	t.Log(assert.ErrorContains(t, err, "leveldb: not found"))
	t.Log(assert.Error(t, err))
}
