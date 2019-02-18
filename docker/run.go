package docker

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)

func FileGetContents(path string) string {
	file, _ := os.Open(path)
	fileText, _ := ioutil.ReadAll(file)

	return string(fileText)
}

// 创建网络
func (d *WDocker) CreateNet(t *WTask) {
	log.Println("创建网络")

	for networkName := range t.ComposeConfig.Networks {
		labels := make(map[string]string)
		labels["com.docker.compose.project"] = t.StackName

		res, err := d.cli.NetworkCreate(d.ctx, t.StackName+"_"+networkName, types.NetworkCreate{
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
func (d *WDocker) Run(t *WTask) error {
	for serviceName, service := range t.ComposeConfig.Services {
		//log.Println(serviceName, service)

		labels := make(map[string]string)
		labels["com.docker.compose.project"] = t.StackName

		serviceContainer := &container.Config{}
		serviceContainer.Labels = labels
		serviceContainer.Image = service.Image

		serviceHost := &container.HostConfig{}

		serviceContainer.ExposedPorts, serviceHost.PortBindings = QuickPortMap(service.Ports)

		endpointsConfig := make(map[string]*network.EndpointSettings)
		for _, networkName := range service.Networks {
			endpointsConfig[t.StackName+"_"+networkName] = &network.EndpointSettings{
				Aliases: []string{serviceName},
			}
		}

		serviceNetwork := &network.NetworkingConfig{
			EndpointsConfig: endpointsConfig,
		}

		log.Printf("ExposedPorts: %+v", serviceContainer.ExposedPorts)
		log.Printf("PortBindings: %+v", serviceHost.PortBindings)
		//continue

		resp, err := d.cli.ContainerCreate(d.ctx, serviceContainer, serviceHost, serviceNetwork, t.StackName+"_"+serviceName+".1.xxxxx")
		if err != nil {
			log.Println(err)
			return err
		}

		if err := d.cli.ContainerStart(d.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
			log.Println(err)
			return err
		}

		log.Println(resp.ID)
	}
	return nil
}

func QuickPortMap(portStrs []string) (nat.PortSet, nat.PortMap) {
	exPortMap := make(nat.PortSet)
	portMap := make(nat.PortMap)
	for _, portStr := range portStrs {
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

		//log.Infof("hostIP=%s, hostPort=%s, targetPort=%s\n", hostIP, hostPort, targetPort)

		ports := []nat.PortBinding{
			{
				HostIP:   hostIP,
				HostPort: hostPort,
			},
		}
		exPortMap[targetPort] = struct{}{}
		portMap[targetPort] = ports
	}
	return exPortMap, portMap
}

func (d *WDocker) RunOne(name string, c *container.Config, h *container.HostConfig, n *network.NetworkingConfig) error {
	resp, err := d.cli.ContainerCreate(d.ctx, c, h, n, name)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := d.cli.ContainerStart(d.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Println(err)
		return err
	}

	log.Println(resp.ID)
	return nil
}
