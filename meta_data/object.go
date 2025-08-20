package meta_data

import (
	"fmt"
	"github.com/Disk-Archive/alto-III-go/http"
	"github.com/google/uuid"
	"time"
)

type (
	Object struct {
		Hostname string
		UseSsl   bool
	}

	ObjectMetaData struct {
		ID uuid.UUID `json:"id"`

		CreatedAt time.Time `json:"created_at"`

		DiskId string `json:"disk_id"`

		ObjectName string `json:"object_name"`
		ObjectSize int64  `json:"object_size"`

		Md5Checksum                   string    `json:"md5_checksum"`
		Md5ChecksumCalculationDate    time.Time `json:"md5_checksum_calculation_date"`
		Sha256Checksum                string    `json:"sha256_checksum"`
		Sha256ChecksumCalculationDate time.Time `json:"sha256_checksum_calculation_date"`
		PreCopyMd5Hash                string    `json:"pre_copy_md5_hash"`

		useSsl bool
	}
)

func (o *Object) GetAll() (objects []*ObjectMetaData, err error) {
	return http.Get[[]*ObjectMetaData](o.Hostname, "/api/v1/object/object_metadata", o.UseSsl)
}

func (o *Object) GetObjectByName(objectName string) (object *ObjectMetaData, err error) {
	return http.Get[*ObjectMetaData](o.Hostname, fmt.Sprintf("/api/v1/object/object_metadata?object_name=%s", objectName), o.UseSsl)
}
