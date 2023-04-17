package client

import (
	"encoding/base64"
	"os"

	cfg "compute/config"
	zlog "compute/log"

	"github.com/imroc/req/v3"
	"go.uber.org/zap"
)

// storage.go 请求storage获取计算函数，生成本地的临时文件，挂载至容器内

func getFunc(id string) (string, error) {
	res, err := req.
		R().
		AddQueryParam("id", id).
		Get(cfg.Cfg.StorageUrl + "/function/get")
	if err != nil {
		zlog.Error("get func by id failed", zap.Error(err))
		return "", err
	}
	var m map[string]any
	err = res.Unmarshal(&m)
	if err != nil {
		zlog.Error("unmarshal response failed", zap.Error(err))
		return "", err
	}
	return m["content"].(string), nil
}

// GenerateTempFile 根据id从storage获取计算方法,base64解码后保存在文件中,并挂载至容器内
func GenerateTempFile(funcId string) error {
	data, err := getFunc(funcId)
	if err != nil {
		return err
	}
	res, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		zlog.Error("base64 decode failed", zap.Error(err))
	}
	// TODO 尚未确定路径
	var tempPath string
	file, err := os.Create(tempPath)
	if err != nil {
		zlog.Error("create file failed", zap.Error(err))
		return nil
	}
	if _, err = file.Write(res); err != nil {
		zlog.Error("write data to temp file failed", zap.Error(err))
		return nil
	}
	zlog.Info("generate temp file success")
	return nil
}
