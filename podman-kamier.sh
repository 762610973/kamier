#!/usr/bin/env zsh

#podman pull ubuntu:22.04
podman run -dit --name=base ubuntu:22.04
podman exec -it base apt update
podman exec -it base apt install -y wget
#podman exec -it -w /usr base sudo
podman exec -it -w /usr/local/ base wget https://studygolang.com/dl/golang/go1.20.3.linux-amd64.tar.gz
podman exec -it -w /usr/local/ base tar -zxf go1.20.3.linux-amd64.tar.gz
podman exec -it -w /usr/local/ base mkdir gopath
podman exec -it base ln -s /usr/local/go/bin/go /bin/Go
podman exec -it base Go env -w GOPATH=/usr/local/gopath
podman exec -it base Go env -w GOPROXY=https://goproxy.cn,direct
podman exec -it -w /root base mkdir initModule
podman exec -it -w /root/initModule base Go mod init initModule
podman exec -it -w /root/initModule base Go get google.golang.org/grpc
podman exec -it -w /root/initModule base Go get github.com/natefinch/lumberjack
podman exec -it -w /root/initModule base Go get go.uber.org/zap
podman exec -it -w /root/initModule base Go get go.uber.org/zap/zapcore
podman exec -it -w /root/initModule base Go get github.com/spf13/viper
podman exec -it -w /root/initModule base Go get github.com/fsnotify/fsnotify
podman exec -it -w /root/initModule base Go get google.golang.org/protobuf/reflect/protoreflect
podman exec -it -w /root/initModule base Go get google.golang.org/protobuf/runtime/protoimpl
podman exec -it -w /root/initModule base Go get google.golang.org/grpc/status
podman exec -it -w /root/initModule base Go get google.golang.org/grpc/codes
podman exec -it -w /root/initModule base Go get golang.org/x/text
podman exec -it -w /root/initModule base Go get google.golang.org/genproto
podman exec -it -w /root/initModule base Go get google.golang.org/protobuf
podman exec -it -w /root/initModule base Go get gopkg.in/ini.v1
podman exec -it -w /root/initModule base Go get gopkg.in/yaml.v3
podman exec -it -w /root/initModule base Go get github.com/pelletier/go-toml/v2
podman exec -it -w /root/initModule base Go get github.com/hashicorp/hcl
podman exec -it -w /root/initModule base Go get github.com/subosito/gotenv
podman exec -it -w /root/initModule base Go get go.uber.org/multierr
podman exec -it -w /root/initModule base Go get go.uber.org/atomic
podman exec -it -w /root/initModule base Go get golang.org/x/net
podman exec -it -w /root/initModule base Go get gopkg.in/yaml.v3
podman exec -it -w /root/initModule base Go get github.com/spf13/jwalterweatherman
podman exec -it -w /root/initModule base Go get github.com/spf13/cast
podman exec -it -w /root/initModule base Go get github.com/spf13/afero

# 下载运行所需的go modules
podman commit base kamier:1.0
podman stop base