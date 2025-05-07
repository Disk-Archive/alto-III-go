package alto

import (
	"fmt"
	"github.com/Disk-Archive/alto-III-go/http"
	"github.com/Disk-Archive/alto-III-go/meta_data"
	"net"
)

type AltoIII struct {
	IpAddress net.IP
	Port      int

	IgnoreSslErrors bool

	MetaData *meta_data.MetaData
}

func (a *AltoIII) ArchiveObject(diskId, objectName, md5 string, data []byte) (err error) {
	_, err = http.Post[interface{}](
		a.IpAddress.String(), fmt.Sprintf("/api/v1/copy/archive/object/%s?location", diskId, objectName), data,
	)
	return
}

func (a *AltoIII) DeleteObject(diskId, objectName string) (err error) {
	_, err = http.Delete[interface{}](
		a.IpAddress.String(), fmt.Sprintf("/api/v1/copy/archive/object/%s?location", diskId, objectName),
	)
	return
}
