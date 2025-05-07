package alto

import (
	"github.com/Disk-Archive/alto-III-go/meta_data"
	"net"
)

type AltoIII struct {
	IpAddress net.IP
	Port      int

	IgnoreSslErrors bool

	MetaData *meta_data.MetaData
}
