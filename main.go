// //////////////////////
// Example for testing
// //////////////////////
package main

import (
	"github.com/Disk-Archive/alto-III-go/disk"
	"github.com/Disk-Archive/alto-III-go/file_data"
	"github.com/google/uuid"
)

func main() {

	api := file_data.FileAPI{
		Hostname: "192.168.39.95",
		Port:     443,
	}

	data, err := api.GetFilesByGroup(uuid.MustParse("6920a1f0-c3fb-4db7-be6b-3a7712f08de7"))
	if err != nil {
		panic(err)
	}
	for _, file := range data {
		println(file.Filename)
	}

	data, err = api.GetAllFiles()
	if err != nil {
		panic(err)
	}
	for _, file := range data {
		println(file.Filename)
	}

	DiskApi := disk.DiskAPI{
		Hostname: "192.168.39.95",
		Port:     443,
	}

	DiskData, err := DiskApi.GetAllDisks()
	if err != nil {
		panic(err)
	}
	for _, disk := range DiskData {
		println(disk.SerialNumber)
	}

}
