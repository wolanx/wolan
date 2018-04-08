package compose

import (
	"encoding/json"
)

func (this *Service) String() string {
	jsonStr, _ := json.MarshalIndent(this, "", "  ")
	return string(jsonStr)
}

type Configs struct {
	Version string
	//Services map[string]Service
	Services map[string]*Service
}

type Service struct {
	Image    string   `yaml:"image"`
	Networks []string `yaml:"networks"`
	Ports    []string `yaml:"ports"`
}

// no use
//"github.com/docker/docker/api/types/container"
//"github.com/docker/docker/api/types/network"
//type Service struct {
//	Container *container.Config
//	Host      *container.HostConfig
//	Network   *network.NetworkingConfig
//}
