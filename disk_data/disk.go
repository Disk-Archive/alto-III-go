package disk_data

import (
	"fmt"
	"github.com/Disk-Archive/alto-III-go/alto"
	"github.com/Disk-Archive/alto-III-go/http"
	"github.com/google/uuid"
	"time"
)

type (
	DiskAPI struct {
		Hostname    string
		Port        int
		Credentials *alto.AltoBasicAuthCredentials
	}
	Disk struct {
		Id uuid.UUID `bun:"type:uuid,pk"`

		CreatedAt           time.Time        `json:"created_at"`
		DataPartition       string           `json:"data_partition"`
		DeviceLocation      string           `json:"device_location"`
		FreeSpace           uint64           `json:"free_space"`
		Group               uuid.UUID        `json:"group"`
		Health              DiskHealthStatus `json:"health"` // ok, warn, error
		IsMounted           bool             `json:"is_mounted"`
		IsSpinning          bool             `json:"is_spinning"`
		IsFormated          bool             `json:"is_formated"`
		MetadataPartition   string           `json:"metadata_partition"`
		ModelNumber         string           `json:"model_number"`
		Manufacture         string           `json:"manufacture"`
		SerialNumber        string           `json:"serial_number"`
		TotalSize           uint64           `json:"total_size"`
		RemoveStatus        *RemoveStatus    `json:"remove_status"`
		DataPartitionId     uuid.UUID        `json:"data_partition_id"`
		MetadataPartitionId uuid.UUID        `json:"metadata_partition_id"`
	}
)

func New(hostname string, port int) *DiskAPI {
	return &DiskAPI{
		Hostname: hostname,
		Port:     port,
	}
}

func (d *DiskAPI) GetAllDisks() ([]Disk, error) {
	type res struct {
		Disks []Disk `json:"disks"`
	}
	data, err := http.Get[res](d.Hostname, "/api/v1/disk", d.Credentials.Username, d.Credentials.Password, true, true)
	return data.Disks, err
}

func (d *DiskAPI) GetDiskById(id uuid.UUID) (disk Disk, err error) {
	return http.Get[Disk](d.Hostname, fmt.Sprintf("/api/v1/disk/%s", id), d.Credentials.Username, d.Credentials.Password, true, true)
}

func (d *DiskAPI) BringDiskOnline(diskId uuid.UUID) (mountPoint string, err error) {
	type Response struct {
		MountPoint string `json:"mount_path"`
	}
	result, err := http.Get[Response](d.Hostname, fmt.Sprintf("/api/v1/disk/online/%s", diskId), d.Credentials.Username, d.Credentials.Password, true, true)
	return result.MountPoint, err
}

func (d *DiskAPI) TakeDiskOffline(diskId uuid.UUID) (err error) {
	type Response struct {
		Message string `json:"message"`
	}
	_, err = http.Get[Response](d.Hostname, fmt.Sprintf("/api/v1/disk/offline/%s", diskId), d.Credentials.Username, d.Credentials.Password, true, true)
	return err
}
