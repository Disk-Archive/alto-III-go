package alto

import (
	"fmt"
	"github.com/Disk-Archive/alto-III-go/groups"
	"github.com/Disk-Archive/alto-III-go/http"
	"github.com/Disk-Archive/alto-III-go/meta_data"
)

type AltoIII struct {
	Hostname string
	Port     int

	IgnoreSslErrors bool

	Groups   *groups.Groups
	MetaData *meta_data.MetaData
}

func New(hostname string, port int, ignoreSsl bool) (altoIII *AltoIII) {
	if port < 0 || port > 65535 {
		return nil
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
		Groups: groups.New(hostname),
	}
}

func (a *AltoIII) ArchiveObject(diskId, objectName, md5 string, data []byte) (err error) {
	_, err = http.Post[interface{}](
		a.Hostname, fmt.Sprintf("/api/v1/copy/archive/object/%s?location=%s", diskId, objectName), "", data,
	)
	return
}

func (a *AltoIII) DeleteObject(objectId string) (err error) {
	_, err = http.Delete[interface{}](
		a.Hostname, fmt.Sprintf("/api/v1/copy/delete/object/%s", objectId),
	)
	return
}
