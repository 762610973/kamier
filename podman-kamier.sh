#!/usr/bin/env zsh

podman pull ubuntu:22.04
podman run -dit --name=base ubuntu:22.04
podman exec -it base apt update
podman exec -it base apt install wget
podman exec -it -w /usr/local/ base wget https://studygolang.com/dl/golang/go1.20.3.linux-amd64.tar.gz
podman exec -it -w /usr/local/ base tar -zxf go1.20.3.linux-amd64.tar.gz
podman exec -it -w /usr/local/ base mkdir gopath
podman exec -it base ln -s /usr/local/go/bin/go /bin/Go
podman exec -it base Go env -w GOPATH=/usr/local/gopath
podman exec -it base Go env -w GOPROXY=https://goproxy.cn,direct
podman exec -it Go get google.golang.org/grpc
podman exec -it Go get github.com/natefinch/lumberjack
podman exec -it Go get go.uber.org/zap
podman exec -it Go get go.uber.org/zap/zapcore
podman exec -it Go github.com/spf13/viper
podman exec -it Go github.com/fsnotify/fsnotify
podman exec -it Go google.golang.org/protobuf/reflect/protoreflect
podman exec -it Go google.golang.org/protobuf/runtime/protoimpl
podman exec -it Go google.golang.org/grpc/status
podman exec -it Go google.golang.org/grpc/codes

# 下载运行所需的go modules
podman commit base kamier