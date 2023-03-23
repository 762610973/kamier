package db

import (
	cfg "storage/config"
	"storage/log"
	"storage/model"
	"testing"
)

func TestInsertDocument(t *testing.T) {
	cfg.InitConfig()
	log.InitLogger()
	InitMongoDB()
	fn.InsertOne(ctx, model.Function{
		Name:        "1",
		Id:          10,
		Description: "1",
		Content:     "1",
	})
}
