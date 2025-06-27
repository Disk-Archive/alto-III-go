package main

import (
	"fmt"
	"log"
	"net"

	"alto-III-go/alto"
)

///////////////////////////////////////////
//	This is an example for test purpose  //
///////////////////////////////////////////

func main() {
	// Replace with the actual IP address and port of your AltoIII device
	ip := net.ParseIP("192.168.39.95")
	port := 6480

	a := &alto.AltoIII{
		TcpConnection: alto.TcpConnection{
			IPAddress: ip,
			Port:      port,
		},
	}

	fmt.Println("Connecting to AltoIII...")

	name := a.GetSystemName()
	fmt.Printf("System Name: %s\n", name)

	serial := a.GetSerialNumber()
	fmt.Printf("Serial Number: %s\n", serial)

	groups, err := a.GetGroups()
	if err != nil {
		log.Printf("Failed to get groups: %v\n", err)
	} else {
		fmt.Printf("Groups: %v\n", groups)
	}

	network, err := a.GetNetworkDetails()
	if err != nil {
		log.Printf("Failed to get network details: %v\n", err)
	} else {
		fmt.Println("Network Details:")
		fmt.Printf("  IP Address    : %s\n", network.Ip)
		fmt.Printf("  Netmask       : %s\n", network.Netmask)
		fmt.Printf("  Gateway       : %s\n", network.Gateway)
		fmt.Printf("  Interface Name: %s\n", network.InterfaceName)
	}

	prmConfig, err := a.GetPrometheusConfig()
	if err != nil {
		log.Printf("promethes config: %v\n", err)
	} else {
		fmt.Printf("prometheus config: %v\n", *prmConfig)
	}

	prmDb, err := a.GetPrometheusDb()
	if err != nil {
		log.Printf("promethes DB: %v\n", err)
	} else {
		fmt.Printf("prometheus DB: %v\n", *prmDb)
	}

	smartData, err := a.GetSmartDataByDiskUuid("66eca1d7-d785-45e6-8945-fb2a6e318a78")
	if err != nil {
		log.Printf("smartData: %v\n", err)
	} else {
		fmt.Printf("smartData: %v\n", *smartData)
	}
}
