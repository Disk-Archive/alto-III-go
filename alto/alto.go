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

	UseSsl          bool
	IgnoreSslErrors bool

	Groups   *groups.Groups
	MetaData *meta_data.MetaData
}

func New(hostname string, port int, useSsl, ignoreSsl bool) (altoIII *AltoIII) {
	if port < 0 || port > 65535 {
		return nil
	}

	return &AltoIII{
		Hostname:        hostname,
		Port:            port,
		UseSsl:          useSsl,
		IgnoreSslErrors: ignoreSsl,
		MetaData: &meta_data.MetaData{
			Object: &meta_data.Object{
				Hostname: hostname,
				UseSsl:   useSsl,
			},
		},
		Groups: groups.New(hostname, useSsl),
	}
}

func (a *AltoIII) ArchiveObject(groupId, diskId, objectName, md5 string, data []byte) (err error) {
	_, err = http.Post[interface{}](
		a.Hostname, fmt.Sprintf("/api/v1/copy/archive/object?location=%s&disk_id=%s&group_id=%s", objectName, diskId, groupId), md5, data, a.UseSsl,
	)
	return
}

func (a *AltoIII) RestoreObject(groupId, diskId, objectName, md5 string) (fileBytes []byte, err error) {
	fileBytes, err = http.Get[[]byte](a.Hostname, fmt.Sprintf("/api/v1/copy/restore/object?location=%s&disk_id=%s&group_id=%s", objectName, diskId, groupId), a.UseSsl)
	return
}

func (a *AltoIII) DeleteObject(objectId string) (err error) {
	_, err = http.Delete[interface{}](
		a.Hostname, fmt.Sprintf("/api/v1/copy/delete/object/%s", objectId), a.UseSsl,
	)
	return
}
