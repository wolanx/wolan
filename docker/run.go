package docker

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/zx5435/wolan/compose"
	"log"
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
	stackName = "cdemo"
	log.Println("WDocker::Deploy", stackName)

	doLoad()

	// step.1 networks 网格
	this.CreateNet()

	// step.2 volumes

	// step.3 services 部署多个 container
	this.CreateContainer()
}

// 创建网络
func (this *WDocker) CreateNet() {
	log.Println(composeConfig.Networks)

	for networkName := range composeConfig.Networks {
		labels := make(map[string]string)
		labels["com.docker.compose.project"] = stackName

		res, err := this.cli.NetworkCreate(this.ctx, stackName+"_"+networkName, types.NetworkCreate{
			CheckDuplicate: true,
			//Driver: "overlay",
			//Scope:  "swarm",
			Labels: labels,
		})
		if err != nil {
			log.Printf(err.Error())
		} else {
			log.Printf("%#v\n", res)
		}
	}
}

// 创建容器
func (this *WDocker) CreateContainer() {
	for serviceName, service := range composeConfig.Services {
		log.Println(serviceName, service)

		labels := make(map[string]string)
		labels["com.docker.compose.project"] = stackName

		serviceContainer := &container.Config{}
		serviceContainer.Labels = labels
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

			//log.Printf("hostIP=%s, hostPort=%s, targetPort=%s\n", hostIP, hostPort, targetPort)

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

		endpointsConfig := make(map[string]*network.EndpointSettings)
		for _, networkName := range service.Networks {
			endpointsConfig[stackName+"_"+networkName] = &network.EndpointSettings{
				Aliases: []string{serviceName},
			}
		}

		serviceNetwork := &network.NetworkingConfig{
			EndpointsConfig: endpointsConfig,
		}

		log.Printf("ExposedPorts: %+v\nPortBindings: %+v\n", serviceContainer.ExposedPorts, serviceHost.PortBindings)
		//continue

		resp, err := this.cli.ContainerCreate(this.ctx, serviceContainer, serviceHost, serviceNetwork, stackName+"_"+serviceName+".1.xxxxx")
		if err != nil {
			panic(err)
		}

		if err := this.cli.ContainerStart(this.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
			panic(err)
		}

		log.Println(resp.ID)
	}
}
