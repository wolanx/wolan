package docker

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"os"
	"io/ioutil"
	"github.com/zx5435/wolan/compose"
	"github.com/docker/docker/api/types/network"
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
	fmt.Println(composeConfig)

	//netStr := "cdemo_mynet"

	// 部署多个 container
	for name, service := range composeConfig.Services {
		service.Host.NetworkMode = "cdemo_mynet"
		fmt.Println(name)
		//fmt.Printf("%#v\n", service.Container)
		//fmt.Printf("%#v\n", service.Host)

		service.Network = &network.NetworkingConfig{
			EndpointsConfig: map[string]*network.EndpointSettings{"cdemo_mynet": &network.EndpointSettings{
				Aliases: []string{name},
			}},
		}

		fmt.Printf("%#v\n", service.Network)
		fmt.Println(composeConfig)
		//break

		resp, err := this.cli.ContainerCreate(this.ctx, service.Container, service.Host, service.Network, "cdemo_"+name+"_1")
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
