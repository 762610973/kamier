package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	cfg "storage/config"
	zlog "storage/log"
	"time"
)

var (
	client = &mongo.Client{}
	db     = &mongo.Database{}
	fn     = &mongo.Collection{}
	data   = &mongo.Collection{}
	node   = &mongo.Collection{}
	ctx    = context.Background()
)

func InitMongoDB() {
	ctx1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	uri := fmt.Sprintf("mongodb://%s:%s", cfg.Cfg.Storage.MongoDBAddr, cfg.Cfg.Storage.MongoDBPort)
	client, err = mongo.Connect(ctx1, options.
		Client().ApplyURI(uri).
		SetAuth(options.Credential{
			Username: cfg.Cfg.Storage.Username,
			Password: cfg.Cfg.Storage.Password,
		}).
		SetMaxPoolSize(20))
	if err != nil {
		zlog.Panic("connect mongodb failed", zap.Error(err))
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		zlog.Panic("connect mongodb failed", zap.Error(err))
	}
	zlog.Info("connect mongodb success")
	// start with clean db
	if cfg.Cfg.Storage.Delete {
		if checkDBExist() {
			zlog.Info("database exist, then drop it")
			err = client.Database(cfg.Cfg.Storage.DBName).Drop(ctx)
			if err != nil {
				zlog.Warn("delete database failed", zap.Error(err))
			}
			zlog.Info("delete database: " + cfg.Cfg.Storage.DBName)
		}
	}

	db = client.Database(cfg.Cfg.Storage.DBName)
	fn = db.Collection("function")
	data = db.Collection("data")
	node = db.Collection("node")
}

func InsertData(types string, value any) {

}
func checkDBExist() bool {
	ctx := context.Background()
	// 此处filter填写空
	names, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		zlog.Panic("list database name failed", zap.Error(err))
	}
	exist := false
	for _, name := range names {
		if name == cfg.Cfg.Storage.DBName {
			exist = true
		}
	}
	return exist
}
