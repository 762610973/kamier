package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Logger 日志级别,编码格式,日志输出位置的配置
type Logger struct {
	// 日志级别
	Level string `yaml:"level"`
	// 编码格式:json、console
	Encoding     string `yaml:"encoding"`
	WriteConsole bool   `yaml:"writeConsole"`

	FileName string `yaml:"fileName"`
	// 切割之前,日志文件的最大大小(MB)
	MaxSize int `yaml:"maxSize"`
	// 保留旧文件的最大个数
	MaxBackUps int `yaml:"maxBackUps"`
	// 保留旧文件的最大天数
	MaxAge int `yaml:"maxAge"`
	// 是否压缩旧文件
	Compress bool `yaml:"compress"`
}

type NetWork struct {
	// 本机节点名
	OrgName  string `yaml:"orgName"`
	HttpPort string `yaml:"httpPort"`

	GrpcAddr string `yaml:"grpcAddr"`
	GrpcPort string `yaml:"grpcPort"`

	SocketPath string `yaml:"socketPath"`
}

type Storage struct {
	StorageUrl  string `yaml:"storageUrl"`
	LevelDBPath string `yaml:"levelDBPath"`
	MysqlAddr   string `yaml:"mysqlAddr"`
	MysqlPort   string `yaml:"mysqlPort"`
}

type config struct {
	Storage `yaml:"storage"`
	NetWork `yaml:"netWork"`
	Logger  `yaml:"logger"`
}

var Cfg config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./config/")
	viper.AddConfigPath("./compute/config/")
	viper.AddConfigPath("../config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置文件失败")
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("配置文件发生改变", e.Name)
	})
	//var cfg CoreConfig
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		panic("反序列化配置文件失败")
	}
}
