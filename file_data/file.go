package file_data

import (
	"fmt"
	"github.com/Disk-Archive/alto-III-go/http"
	"github.com/google/uuid"
	"time"
)

type (
	FileAPI struct {
		Hostname string
		Port     int
	}

	File struct {
		ID uuid.UUID `json:"id"`

		CreatedAt time.Time `json:"created_at"`
		Filename  string    `json:"filename"`
		FileSize  int64     `json:"file_size"`
		FilePath  string    `json:"file_path"`

		DiskId string `json:"disk_id"`

		Md5Checksum                string    `json:"md5_checksum"`
		Md5ChecksumCalculationDate time.Time `json:"md5_checksum_calculation_date"`

		Sha256Checksum                string    `json:"sha256_checksum"`
		Sha256ChecksumCalculationDate time.Time `json:"sha256_checksum_calculation_date"`

		PreCopyMd5Hash string `json:"pre_copy_md5_hash"`
	}
)

func New(host string, port int) *FileAPI {
	return &FileAPI{
		Hostname: host,
		Port:     port,
	}
}

func (f *FileAPI) GetFilesByGroup(groupId uuid.UUID) (file []File, err error) {

	res, err := http.Get[[]File](f.Hostname, fmt.Sprintf("/api/v1/file/by-group-id/%s", groupId))

	return res, nil
}

func (f *FileAPI) GetAllFiles() (file []File, err error) {

	res, err := http.Get[[]File](f.Hostname, "/api/v1/file/file_metadata")

	return res, nil
}
