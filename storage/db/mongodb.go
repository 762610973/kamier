package db

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
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
	ctx    = context.Background()
)

const (
	Function = "function"
	Data     = "data"
	Node     = "node"
	ID       = "_id"
)

func InitMongoDB() {
	ctx1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	uri := fmt.Sprintf("mongodb://%s:%s", cfg.Cfg.Storage.MongoDBAddr, cfg.Cfg.Storage.MongoDBPort)
	client, err = mongo.Connect(ctx1, options.
		Client().ApplyURI(uri).
		SetAuth(options.Credential{
			Username: cfg.Cfg.Username,
			Password: cfg.Cfg.Password,
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
	if cfg.Cfg.Delete {
		if checkDBExist() {
			zlog.Info("database exist, then drop it")
			err = client.Database(cfg.Cfg.DBName).Drop(ctx)
			if err != nil {
				zlog.Warn("delete database failed", zap.Error(err))
			}
			zlog.Info("delete database: " + cfg.Cfg.DBName)
		}
	}
	// init database and collection
	db = client.Database(cfg.Cfg.DBName)
}

// InsertDocument 插入文档
func InsertDocument(types string, value any, id string) error {
	var err error
	err = db.Collection(types).FindOne(ctx, bson.M{ID: id}).Err()
	if err == nil {
		zlog.Error("document exist, can't insert", zap.Error(err))
		return errors.New("document exist, can't insert")
	}
	// err != nil, document not exist, insert document
	_, err = db.Collection(types).InsertOne(ctx, value)
	if err != nil {
		zlog.Error(fmt.Sprintf("insert %s failed", types), zap.Error(err))
		return err
	}
	zlog.Debug(fmt.Sprintf("insert %s success", types), zap.Any("value", value))
	return nil
}

// FindDocument 根据_id查询文档
func FindDocument(types string, filter any) (error, bson.M) {
	var err error
	var res bson.M
	err = db.Collection(types).FindOne(ctx, filter).Decode(&res)
	if err != nil {
		zlog.Error(fmt.Sprintf("find %s failed", types), zap.Error(err), zap.Any("filter", filter))
		return err, nil
	}
	return nil, res
}

// FindAllDocument 查询所有文档
func FindAllDocument(types string) (error, []bson.M) {
	var err error
	var res []bson.M
	var cur *mongo.Cursor
	cur, err = db.Collection(types).Find(ctx, bson.M{})
	if err != nil {
		zlog.Error("find collection failed", zap.Error(err))
		return err, nil
	}
	err = cur.All(ctx, &res)
	if err != nil {
		zlog.Error(fmt.Sprintf("find all %s failed", types), zap.Error(err))
		return err, nil
	}
	return nil, res
}

// DeleteDocument 根据_id删除文档
func DeleteDocument(types string, filter any) error {
	var err error
	// delete a document without check exist
	_, err = db.Collection(types).DeleteOne(ctx, filter)
	if err != nil {
		zlog.Error(fmt.Sprintf("%s delete failed", types), zap.Error(err))
		return err
	}
	zlog.Debug(fmt.Sprintf("%s delete success", types))
	return nil
}

func UpdateDocument(types string, id, update any) error {
	_, err := db.Collection(types).UpdateByID(ctx, id, update)
	if err != nil {
		zlog.Error(types+" update failed", zap.Error(err))
		return err
	}
	zlog.Debug(types + " update failed")
	return nil
}

// 检查数据库是否存在
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
