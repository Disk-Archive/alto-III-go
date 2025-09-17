// //////////////////////
// Example for testing
// //////////////////////
package main

import (
	"github.com/Disk-Archive/alto-III-go/alto"
	"github.com/Disk-Archive/alto-III-go/disk_data"
	"github.com/Disk-Archive/alto-III-go/file_data"
	"github.com/Disk-Archive/alto-III-go/meta_data"
	"github.com/google/uuid"
)

func main() {

	api := file_data.FileAPI{
		Hostname: "192.168.39.95",
		Port:     443,
		Credentials: &alto.AltoBasicAuthCredentials{
			Username: "admin",
			Password: "admin",
		},
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

	DiskApi := disk_data.DiskAPI{
		Hostname: "192.168.39.95",
		Port:     443,
		Credentials: &alto.AltoBasicAuthCredentials{
			Username: "admin",
			Password: "admin",
		},
	}

	DiskData, err := DiskApi.GetAllDisks()
	if err != nil {
		panic(err)
	}
	for _, disk := range DiskData {
		println(disk.SerialNumber)
	}

	objectApi := meta_data.Object{
		Hostname: "192.168.39.95",
		UseSsl:   true,
		Credentials: &meta_data.AltoBasicAuthCredentials{
			Username: "admin",
			Password: "admin",
		},
	}

	println("bucket objects: ")
	objectData, err := objectApi.GetObjectsByBucketId(uuid.MustParse("908be984-a7fa-4e99-9574-1eb38aaac50f"))
	if err != nil {
		panic(err)
	}
	for _, object := range objectData {
		println(object.ObjectName)
	}
}
