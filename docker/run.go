package docker

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"os"
	"io/ioutil"
	"github.com/zx5435/wolan/compose"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

func FileGetConnents(path string) string {
	file, _ := os.Open(path)
	fileText, _ := ioutil.ReadAll(file)

	return string(fileText)
}

func (this *WDocker) Deploy() {
	//fmt.Println("WDocker::Deploy")
	//fmt.Println(wCenter.Config)

	//this.CreateNet()
	//return

	composeConfig := compose.Parse(FileGetConnents(wCenter.WorkDir + "/" + wCenter.Config.DockerCompose))
	//fmt.Println(composeConfig)

	//netStr := "cdemo_mynet"

	// 部署多个 container
	for name, service := range composeConfig.Services {
		serviceContainer := &container.Config{}

		serviceHost := &container.HostConfig{}
		for _, port := range service.Ports {
			fmt.Println(port)

			ports := []nat.PortBinding{
				{
					HostIP:   "qwe:123",
					HostPort: "2323",
				},
			}
			portMap := make(nat.PortMap)
			portMap["80"] = ports

			//serviceHost.PortBindings["80"] = []nat.PortBinding{
			//	{
			//		HostIP:   "qwe:123",
			//		HostPort: "2323",
			//	},
			//}
			serviceHost.PortBindings = portMap
		}

		serviceNetwork := &network.NetworkingConfig{
			EndpointsConfig: map[string]*network.EndpointSettings{"cdemo_mynet": &network.EndpointSettings{
				Aliases: []string{name},
			}},
		}

		fmt.Println(name)
		fmt.Println(service)
		fmt.Printf("%+v\n\n", serviceHost.PortBindings)
		continue

		resp, err := this.cli.ContainerCreate(this.ctx, serviceContainer, serviceHost, serviceNetwork, "cdemo_"+name+"_1")
		if err != nil {
			panic(err)
		}

		if err := this.cli.ContainerStart(this.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
			panic(err)
		}

		fmt.Println(resp.ID)
	}
}

func (this *WDocker) CreateNet() {
	this.cli.NetworkCreate(this.ctx, "cdemo_mynet", types.NetworkCreate{})
}
