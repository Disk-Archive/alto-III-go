package groups

import (
	"encoding/json"
	"github.com/Disk-Archive/alto-III-go/http"
	"time"
)

type (
	Groups struct {
		Hostname string
		Port     int
		Groups   []Group
		UseSsl   bool

		Credentials *AltoBasicAuthCredentials
	}

	Group struct {
		Id        string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`

		useSsl bool

		Credentials *AltoBasicAuthCredentials
	}

	AltoBasicAuthCredentials struct {
		Username string
		Password string
	}
)

func New(hostname, username, password string, port int, useSsl bool) (groups *Groups) {
	groups = &Groups{
		Hostname:    hostname,
		Port:        port,
		UseSsl:      useSsl,
		Credentials: &AltoBasicAuthCredentials{Username: username, Password: password},
	}
	groups.Groups, _ = groups.GetGroups()
	return
}

func (g *Groups) CreateGroup(group *Group) (err error) {
	data, err := json.Marshal(group)
	_, err = http.Post[Group](g.Hostname, "/api/v1/groups/create", "", g.Credentials.Username, g.Credentials.Password, "application/json", g.Port, data, g.UseSsl, true)
	return err
}

func (g *Groups) GetGroups() (groups []Group, err error) {
	result, err := http.Get[[]Group](g.Hostname, "/api/v1/groups", g.Credentials.Username, g.Credentials.Password, g.Port, g.UseSsl, true)
	return result, err
}
