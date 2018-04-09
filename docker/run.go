package docker

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/zx5435/wolan/compose"
)

var (
	stackName     string
	composeConfig *compose.Configs
)

func FileGetContents(path string) string {
	file, _ := os.Open(path)
	fileText, _ := ioutil.ReadAll(file)

	return string(fileText)
}

func doLoad() {
	composeConfig = compose.Parse(FileGetContents(wCenter.WorkDir + "/" + wCenter.Config.DockerCompose))
}

func (this *WDocker) Deploy() {
	fmt.Println("WDocker::Deploy")
	doLoad()

	stackName = "cdemo"

	// step.1 networks 网格
	this.CreateNet()
	return

	// step.2 volumes

	// step.3 services 部署多个 container
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
			EndpointsConfig: map[string]*network.EndpointSettings{stackName + "_mynet": {
				Aliases: []string{name},
			}},
		}

		fmt.Printf("ExposedPorts: %+v\n", serviceContainer.ExposedPorts)
		fmt.Printf("PortBindings: %+v\n", serviceHost.PortBindings)
		//continue

		resp, err := this.cli.ContainerCreate(this.ctx, serviceContainer, serviceHost, serviceNetwork, stackName+"_"+name+".1.xxxxx")
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
	fmt.Println(stackName)
	fmt.Println(composeConfig.Networks)

	_, err := this.cli.NetworkCreate(this.ctx, stackName+"_mynet", types.NetworkCreate{})
	if err != nil {
		panic(err)
	}
}
