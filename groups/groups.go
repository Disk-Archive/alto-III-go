package groups

import (
	"encoding/json"
	"github.com/Disk-Archive/alto-III-go/http"
	"time"
)

type (
	Groups struct {
		Hostname string
		Groups   []Group
	}

	Group struct {
		Id        string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	}
)

func New(hostname string) (groups *Groups) {
	groups = &Groups{
		Hostname: hostname,
	}
	groups.Groups, _ = groups.GetGroups()
	return
}

func (g *Groups) CreateGroup(group *Group) (err error) {
	data, err := json.Marshal(group)
	_, err = http.Post[Group](g.Hostname, "/api/v1/groups/create", "", data)
	return err
}

func (g *Groups) GetGroups() (groups []Group, err error) {
	result, err := http.Get[[]Group](g.Hostname, "/api/v1/groups")
	return result, err
}
