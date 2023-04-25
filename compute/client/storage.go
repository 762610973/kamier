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

// funcRes 接收storage的响应参数
type funcRes struct {
	Data struct {
		Name        string `json:"name,omitempty"`
		Id          string `json:"id,omitempty"`
		Description string `json:"description,omitempty"`
		Content     string `json:"content,omitempty"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func getFunc(id string) ([2]string, error) {
	res, err := req.
		R().
		AddQueryParam("id", id).
		Get(cfg.Cfg.StorageUrl + "/function/get")
	if err != nil {
		zlog.Error("get func by id failed", zap.Error(err))
		return [2]string{}, err
	}
	var f funcRes
	err = res.UnmarshalJson(&f)
	if err != nil {
		zlog.Error("unmarshal response failed", zap.Error(err))
		return [2]string{}, err
	}
	return [2]string{f.Data.Content, f.Data.Name}, nil
}

// GenerateTempFile 根据id从storage获取计算方法,base64解码后保存在文件中,并挂载至容器内
func GenerateTempFile(funcId string) ([2]string, error) {
	m, err := getFunc(funcId)
	if err != nil {
		return [2]string{}, err
	}
	res, err := base64.StdEncoding.DecodeString(m[0])
	if err != nil {
		zlog.Error("base64 decode failed", zap.Error(err))
	}
	var tempPath string
	err = os.Mkdir("../container/"+m[1], 0777)
	if err != nil {
		zlog.Error("mkdir package failed", zap.Error(err))
	}
	tempPath = "../container/" + m[1] + "/" + m[1] + ".go"
	//tempPath = "../container/exec/" + m[1] + ".go"
	file, err := os.Create(tempPath)
	if err != nil {
		zlog.Error("create file failed", zap.Error(err))
		return [2]string{}, nil
	}
	if _, err = file.Write(res); err != nil {
		zlog.Error("write data to temp file failed", zap.Error(err))
		return [2]string{}, nil
	}
	m[0] = tempPath
	zlog.Info("generate temp file success", zap.Any("m->", m))
	return m, nil
}
