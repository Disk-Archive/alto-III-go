package alto

import (
	"fmt"
	"github.com/Disk-Archive/alto-III-go/groups"
	"github.com/Disk-Archive/alto-III-go/http"
	"github.com/Disk-Archive/alto-III-go/meta_data"
)

type (
	AltoIII struct {
		Hostname string
		Port     int

		UseSsl          bool
		IgnoreSslErrors bool

		Groups   *groups.Groups
		MetaData *meta_data.MetaData

		Credentials *AltoBasicAuthCredentials
	}

	AltoBasicAuthCredentials struct {
		Username string
		Password string
	}
)

func New(hostname string, port int, username, password string, useSsl, ignoreSsl bool) (altoIII *AltoIII) {
	if port < 0 || port > 65535 {
		return nil
	}

	return &AltoIII{
		Hostname: hostname,
		Port:     port,
		Credentials: &AltoBasicAuthCredentials{
			Username: username, Password: password,
		},
		UseSsl:          useSsl,
		IgnoreSslErrors: ignoreSsl,
		MetaData: &meta_data.MetaData{
			Object: &meta_data.Object{
				Hostname: hostname,
				UseSsl:   useSsl,
				Credentials: &meta_data.AltoBasicAuthCredentials{
					Username: username, Password: password,
				},
			},
		},
		Groups: groups.New(hostname, username, password, useSsl),
	}
}

func (a *AltoIII) ArchiveObject(groupId, diskId, objectName, md5 string, data []byte) (err error) {
	_, err = http.Post[interface{}](
		a.Hostname, fmt.Sprintf("/api/v1/copy/archive/object?location=%s&disk_id=%s&group_id=%s", objectName, diskId, groupId), md5, a.Credentials.Username, a.Credentials.Password, data, a.UseSsl,
	)
	return
}

func (a *AltoIII) RestoreObject(groupId, diskId, objectName, md5 string) (fileBytes []byte, err error) {
	fileBytes, err = http.Get[[]byte](a.Hostname, fmt.Sprintf("/api/v1/copy/restore/object?location=%s&disk_id=%s&group_id=%s", objectName, diskId, groupId), a.Credentials.Username, a.Credentials.Password, a.UseSsl)
	return
}

func (a *AltoIII) DeleteObject(objectId string) (err error) {
	_, err = http.Delete[interface{}](
		a.Hostname, fmt.Sprintf("/api/v1/copy/delete/object/%s", objectId), a.Credentials.Username, a.Credentials.Password, a.UseSsl,
	)
	return
}
