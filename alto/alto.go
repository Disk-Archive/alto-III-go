package alto

import "time"

type AltoIII struct {
	TcpConnection
}

func (a *AltoIII) GetSystemName() (altoName string) {
	altoName, _ = a.sendTcp("get|system_name", time.Second*5)
	return altoName
}

func (a *AltoIII) GetSerialNumber() (altoSerial string) {
	altoSerial, _ = a.sendTcp("get|system_serial", time.Second*5)
	return altoSerial
}
