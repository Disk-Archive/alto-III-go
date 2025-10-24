package alto

import (
	"fmt"
	"github.com/Disk-Archive/alto-III-go/groups"
	"github.com/Disk-Archive/alto-III-go/http"
	"github.com/Disk-Archive/alto-III-go/meta_data"
	"github.com/google/uuid"
	"net/url"
)

type (
	AltoIII struct {
		Hostname string
		Port     int

		UseSsl         bool
		InsecureSslReq bool

		Groups   *groups.Groups
		MetaData *meta_data.MetaData

		Credentials *AltoBasicAuthCredentials
	}

	AltoBasicAuthCredentials struct {
		Username string
		Password string
	}
)

func New(hostname string, port int, username, password string, useSsl, insecureSslReq bool) (altoIII *AltoIII) {
	if port < 0 || port > 65535 {
		return nil
	}

	return &AltoIII{
		Hostname: hostname,
		Port:     port,
		Credentials: &AltoBasicAuthCredentials{
			Username: username, Password: password,
		},
		UseSsl:         useSsl,
		InsecureSslReq: insecureSslReq,
		MetaData: &meta_data.MetaData{
			Object: &meta_data.Object{
				Hostname:       hostname,
				Port:           port,
				UseSsl:         useSsl,
				InsecureSslReq: insecureSslReq,
				Credentials: &meta_data.AltoBasicAuthCredentials{
					Username: username, Password: password,
				},
			},
		},
		Groups: groups.New(hostname, username, password, port, useSsl),
	}
}

func (a *AltoIII) ArchiveObject(groupId, diskId, objectName, md5 string, data []byte, bucketId uuid.UUID) (err error) {
	_, err = http.Post[interface{}](
		a.Hostname, fmt.Sprintf("/api/v1/copy/archive/object?location=%s&disk_id=%s&group_id=%s&bucket_id=%s", url.QueryEscape(objectName), diskId, groupId, bucketId), md5, a.Credentials.Username, a.Credentials.Password, a.Port, data, a.UseSsl, a.InsecureSslReq,
	)
	return
}

func (a *AltoIII) RestoreObject(groupId, diskId, objectName, md5 string) (fileBytes []byte, err error) {
	fileBytes, err = http.Get[[]byte](a.Hostname, fmt.Sprintf("/api/v1/copy/restore/object?location=%s&disk_id=%s&group_id=%s", url.QueryEscape(objectName), diskId, groupId), a.Credentials.Username, a.Credentials.Password, a.Port, a.UseSsl, a.InsecureSslReq)
	return
}

func (a *AltoIII) DeleteObject(objectId string) (err error) {
	_, err = http.Delete[interface{}](
		a.Hostname, fmt.Sprintf("/api/v1/copy/delete/object/%s", objectId), a.Credentials.Username, a.Credentials.Password, a.Port, a.UseSsl, a.InsecureSslReq,
	)
	return
}

func (a *AltoIII) CopyObject(sPath, dPath string, sBucket, dBucket uuid.UUID) (err error) {
	_, err = http.Patch[interface{}](
		a.Hostname, fmt.Sprintf("/api/v1/copy/copy/object?s_path=%s&d_path=%s&s_bucket_id=%s&d_bucket_id=%s", url.QueryEscape(sPath), url.QueryEscape(dPath), sBucket, dBucket), a.Credentials.Username, a.Credentials.Password, a.Port, a.UseSsl, a.InsecureSslReq,
	)
	return
}
