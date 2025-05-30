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

func (a *AltoIII) GetGroups() (groupsList []string, err error) {
	res, _ := a.sendTcp("get|groups", time.Second*5)
	groupsList, err = ParseAltoResponse(res)
	return
}

func (a *AltoIII) GetNetworkDetails() (networkDetails *Network, err error) {
	res, _ := a.sendTcp("get|network_address", time.Second*5)
	resArr, err := ParseAltoResponse(res)
	if err != nil {
		return nil, err
	}
	networkDetails = &Network{
		Ip:            resArr[1],
		Netmask:       resArr[2],
		Gateway:       resArr[3],
		InterfaceName: resArr[4],
	}
	return
}

func (a *AltoIII) GetPrometheusConfig() (config *string, err error) {
	res, _ := a.sendTcp("get|prm_file|prometheus|config", time.Second*5)
	resArr, err := ParseAltoResponse(res)
	if err != nil {
		return nil, err
	}
	return &resArr[1], err
}
