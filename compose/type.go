package compose

import (
	"encoding/json"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

func (this *Configs) String() string {
	jsonStr, _ := json.MarshalIndent(this, "", "  ")
	return string(jsonStr)
}

type Configs struct {
	Version string
	//Services map[string]Service
	Services map[string]*Service
}

// no use
type Service struct {
	Container *container.Config
	Host      *container.HostConfig
	Network   *network.NetworkingConfig
}
