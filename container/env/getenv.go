package env

import (
	zlog "container/log"
	"go.uber.org/zap"
	"os"
	"strconv"
)

const (
	selfName      = "SelfName"
	socketPath    = "SocketPath"
	serial        = "Serial"
	nodeName      = "NodeName"
	membersLength = "MembersLength"
	hostIp        = "Host_IP"
)

func GetSelfName() string {
	name := os.Getenv(selfName)
	zlog.Info("get self name: " + name)
	return name
}

func GetContainerAddr() string {
	addr := os.Getenv(socketPath)
	zlog.Info("get socket path: " + addr)
	return addr
}

func GetSerial() int64 {
	s := os.Getenv(serial)
	i, err := strconv.Atoi(s)
	if err != nil {
		zlog.Error("get serial failed", zap.Error(err))
		return -1
	}
	return int64(i)
}

func GetNodeName() string {
	name := os.Getenv(nodeName)
	zlog.Info("get node name: " + name)
	return name
}

func GetMembersLength() string {
	l := os.Getenv(membersLength)
	zlog.Info("membersLength = " + l)
	return l
}

func GetHostIp() string {
	host := os.Getenv(hostIp)
	zlog.Info("host ip = " + host)
	return host
}
