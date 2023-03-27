package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	cfg "storage/config"
	zlog "storage/log"
	"storage/model"
	"testing"
)

func TestFindDocument(t *testing.T) {
	cfg.InitConfig()
	zlog.InitLogger()
	InitMongoDB()
	const _id = "_id"
	var s bson.M
	fmt.Println(db.Collection(Function).FindOne(ctx, bson.M{_id: "33333"}).Err().Error())
	fmt.Println(db.Collection(Function).FindOne(ctx, bson.M{_id: "33333"}).Decode(&s).Error())
	function := model.Function{
		Name:        "1",
		Id:          "2",
		Description: "3",
		Content:     "4",
	}

	_, err := db.Collection(Function).InsertOne(ctx, function)
	if err != nil {
		zlog.Error("insert document failed", zap.Error(err))
		return
	}
	fmt.Println(FindDocument(Function, bson.M{_id: "2"}))
}

func TestFindAllDocument(t *testing.T) {
	cfg.InitConfig()
	zlog.InitLogger()
	InitMongoDB()
	f1 := model.Function{
		Name:        "1",
		Id:          "1",
		Description: "1",
		Content:     "1",
	}
	f2 := model.Function{
		Name:        "1",
		Id:          "11",
		Description: "1",
		Content:     "1",
	}
	_ = InsertDocument(Function, f1, "1")
	_ = InsertDocument(Function, f2, "11")
	fmt.Println(FindAllDocument(Function))
}
func TestInsertDocument(t *testing.T) {
	cfg.InitConfig()
	zlog.InitLogger()
	InitMongoDB()
	f1 := model.Function{
		Name:        "1",
		Id:          "1",
		Description: "1",
		Content:     "1",
	}
	err := InsertDocument(Function, f1, "1")
	assert.Equal(t, err, nil)
	err = InsertDocument(Function, f1, "1")
	assert.NotEqual(t, err, nil)
}
