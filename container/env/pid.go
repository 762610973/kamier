package env

type PID struct {
	NodeName string
	Serial   int64
}

var Pid PID

func InitPid() {
	Pid = PID{
		NodeName: GetNodeName(),
		Serial:   GetSerial(),
	}
}
