package alto

import (
	"errors"
	"fmt"
	"github.com/Disk-Archive/alto-III-go/http"
	"github.com/Disk-Archive/alto-III-go/meta_data"
)

type AltoIII struct {
	Hostname string
	Port     int

	IgnoreSslErrors bool

	MetaData *meta_data.MetaData
}

func New(hostname string, port int, ignoreSsl bool) (altoIII *AltoIII, err error) {
	if port < 0 || port > 65535 {
		return nil, errors.New("port out of range 0-65535")
	}

	return &AltoIII{
		Hostname:        hostname,
		Port:            port,
		IgnoreSslErrors: ignoreSsl,
		MetaData: &meta_data.MetaData{
			Object: &meta_data.Object{
				Hostname: hostname,
			},
		},
	}, nil
}

func (a *AltoIII) ArchiveObject(diskId, objectName, md5 string, data []byte) (err error) {
	_, err = http.Post[interface{}](
		a.Hostname, fmt.Sprintf("/api/v1/copy/archive/object/%s?location=%s", diskId, objectName), data,
	)
	return
}

func (a *AltoIII) DeleteObject(diskId, objectName string) (err error) {
	_, err = http.Delete[interface{}](
		a.Hostname, fmt.Sprintf("/api/v1/copy/delete/%s?location=%s", diskId, objectName),
	)
	return
}
