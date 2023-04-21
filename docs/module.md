- 日志库: 
  - [zap](https://pkg.go.dev/go.uber.org/zap)
  - [lumberjack](https://github.com/natefinch/lumberjack)
- 配置库
  - [viper](https://github.com/spf13/viper)
  - [fsnotify](https://github.com/fsnotify/fsnotify)
- http框架
  - [hertz](https://www.cloudwego.io/zh/docs/hertz/)
- rpc框架
  - [grpc](http://doc.oschina.net/grpc)
- 数据库
  - [leveldb](https://github.com/google/leveldb)


- proto命令
`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./目标文件`

- 安装protoc
- go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
