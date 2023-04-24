package env

import (
	zlog "container/log"
	"os"
	"strconv"
)

const (
	selfName      = "SelfName"
	socketPath    = "SocketPath"
	serial        = "Serial"
	nodeName      = "NodeName"
	fnName        = "FnName"
	membersLength = "MembersLength"
	hostIp        = "Host_IP"
)

func GetSelfName() string {
	return os.Getenv(selfName)
}

func GetSocketPath() string {
	return os.Getenv(socketPath)
}

func GetSerial() int64 {
	serial := os.Getenv(serial)
	i, err := strconv.ParseInt(serial, 64, 10)
	if err != nil {
		return -1
	}
	return i
}

func GetNodeName() string {
	return os.Getenv(nodeName)
}
func GetFnName() string {
	return os.Getenv(fnName)
}

func GetMembersLength() string {
	return os.Getenv(membersLength)
}

func GetHostIp() string {
	host := os.Getenv(hostIp)
	zlog.Info("host ip = " + host)
	return host
}
