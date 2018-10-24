package compose

import (
	"encoding/json"
)

// no use
//"github.com/docker/docker/api/types/container"
//"github.com/docker/docker/api/types/network"
//type Service struct {
//	Container *container.WolanYAML
//	Host      *container.HostConfig
//	Network   *network.NetworkingConfig
//}
type Configs struct {
	Version string
	//Services map[string]Service
	Services map[string]*Service
	Networks map[string]*Network
}

//nginx:
//  image: zx5435/cdemo-nginx:a
//  networks:
//    - mynet
//  ports:
//    - "80"
type Service struct {
	Image    string   `yaml:"image"`
	Networks []string `yaml:"networks"`
	Ports    []string `yaml:"ports"`
}

//mynet:
//  driver: overlay
//  attachable: true
//  config:
//    - subnet: 172.28.0.0/16
//  external:
//    name: my-pre-existing-network
type Network struct {
	Driver string `yaml:"driver"`
}

func (this *Service) String() string {
	jsonStr, _ := json.MarshalIndent(this, "", "  ")
	return string(jsonStr)
}

func (this *Network) String() string {
	jsonStr, _ := json.MarshalIndent(this, "", "  ")
	return string(jsonStr)
}
