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
	"strings"
	"github.com/pkg/errors"
)

func FileGetConnents(path string) string {
	file, _ := os.Open(path)
	fileText, _ := ioutil.ReadAll(file)

	return string(fileText)
}

func (this *WDocker) Deploy() {
	//this.CreateNet()
	//return

	composeConfig := compose.Parse(FileGetConnents(wCenter.WorkDir + "/" + wCenter.Config.DockerCompose))
	//fmt.Println(composeConfig)

	//netStr := "cdemo_mynet"

	// 部署多个 container
	for name, service := range composeConfig.Services {
		fmt.Println(name)

		serviceContainer := &container.Config{}
		serviceContainer.Image = service.Image

		exPortMap := make(nat.PortSet)
		portMap := make(nat.PortMap)

		serviceHost := &container.HostConfig{}
		for _, portStr := range service.Ports {
			portArr := strings.Split(portStr, ":")

			var (
				hostIP     string
				hostPort   string
				targetPort nat.Port
			)
			switch len(portArr) {
			case 3:
				hostIP = portArr[0]
				hostPort = portArr[1]
				targetPort = nat.Port(portArr[2])
			case 2:
				hostIP = ""
				hostPort = portArr[0]
				targetPort = nat.Port(portArr[1])
			case 1:
				hostIP = ""
				hostPort = ""
				targetPort = nat.Port(portArr[0])
			default:
				panic(errors.New("port arr : !<=3"))
			}

			fmt.Printf("hostIP=%s, hostPort=%s, targetPort=%s\n", hostIP, hostPort, targetPort)

			ports := []nat.PortBinding{
				{
					HostIP:   hostIP,
					HostPort: hostPort,
				},
			}
			exPortMap[targetPort] = struct{}{}
			portMap[targetPort] = ports
		}
		serviceContainer.ExposedPorts = exPortMap
		serviceHost.PortBindings = portMap

		serviceNetwork := &network.NetworkingConfig{
			EndpointsConfig: map[string]*network.EndpointSettings{"cdemo_mynet": {
				Aliases: []string{name},
			}},
		}

		fmt.Printf("%+v\n\n", serviceContainer.ExposedPorts)
		fmt.Printf("%+v\n\n", serviceHost.PortBindings)
		//continue

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
