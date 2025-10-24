package alto

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

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

func (a *AltoIII) GetPrometheusDb() (config *string, err error) {
	res, _ := a.sendTcp("get|db", time.Second*5)

	if strings.HasPrefix(res, "<?xml") {
		return &res, nil
	}

	return nil, fmt.Errorf("error getting prometheus DB %v: ", res)
}

func (a *AltoIII) GetSmartDataByDiskUuid(uuid string) (config *string, err error) {
	cmd := fmt.Sprintf("get|disk_smart_data|%s", uuid)
	res, _ := a.sendTcp(cmd, time.Second*5)
	resArr, err := ParseAltoResponse(res)
	if err != nil {
		return nil, err
	}
	return &resArr[1], err
}

func (a *AltoIII) GetChassis() ([]Chassis, error) {
	// Get total number of chassis
	res, _ := a.sendTcp("get|chassis|count", time.Second*5)
	resArr, err := ParseAltoResponse(res)
	if err != nil {
		return nil, fmt.Errorf("failed to parse chassis count: %v", err)
	}

	chassisCount, err := strconv.Atoi(resArr[1])
	if err != nil {
		return nil, fmt.Errorf("invalid chassis count value: %v", resArr[1])
	}

	chassisList := make([]Chassis, 0, chassisCount)

	// Loop through chassis indexes and build struct
	for i := 0; i < chassisCount; i++ {
		parseField := func(cmd string) (int, error) {
			fieldRes, _ := a.sendTcp(cmd, time.Second*5)
			fieldArr, err := ParseAltoResponse(fieldRes)
			if err != nil {
				return 0, fmt.Errorf("failed to parse command '%s': %v", cmd, err)
			}

			val, convErr := strconv.Atoi(fieldArr[1])
			if convErr != nil {
				return 0, fmt.Errorf("invalid integer from '%s': %v", cmd, fieldArr[1])
			}
			return val, nil
		}

		rows, err := parseField(fmt.Sprintf("get|chassis|rows|%d", i))
		if err != nil {
			return nil, err
		}

		firstSlot, err := parseField(fmt.Sprintf("get|chassis|first_slot|%d", i))
		if err != nil {
			return nil, err
		}

		lastSlot, err := parseField(fmt.Sprintf("get|chassis|last_slot|%d", i))
		if err != nil {
			return nil, err
		}

		totalSlots, err := parseField(fmt.Sprintf("get|chassis|total_slots|%d", i))
		if err != nil {
			return nil, err
		}

		layout, activeSlots := ChassisSlotLayoutGet(totalSlots)

		chassisList = append(chassisList, Chassis{
			TotalSlots:  totalSlots,
			ActiveSlots: activeSlots,
			Rows:        rows,
			FirstSlot:   firstSlot,
			LastSlot:    lastSlot,
			Layout:      layout,
		})
	}

	return chassisList, nil
}
