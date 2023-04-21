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
#podman exec -it Go get
# 下载运行所需的go modules
podman commit base kamier