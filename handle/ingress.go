package handle

import (
	"github.com/labstack/echo"
	"github.com/zx5435/wolan/docker"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"log"
)

func IngressStart(c echo.Context) error {
	task := docker.GetTask("ingress")

	volumes := make(map[string]struct{})
	volumes[task.TaskDir] = struct{}{}

	binds := []string{task.TaskDir + "/conf.d:/etc/nginx/conf.d:rw"}

	for _, dir := range task.WolanYAML.Volumes {
		binds = append(binds, dir)
	}

	log.Println(binds)

	d := docker.NewWDocker()
	dConfig := &container.Config{
		Image:   "zx5435/wolan:ingress",
		Volumes: volumes,
	}
	dHost := &container.HostConfig{
		Binds: binds,
	}

	dConfig.ExposedPorts, dHost.PortBindings = docker.QuickPortMap([]string{
		"80:80",
		"443:443",
	})

	err := d.RunOne("wolan-ingress", dConfig, dHost, &network.NetworkingConfig{})

	data := make(map[string]interface{})

	if err != nil {
		data["data"] = err.Error()
		return c.JSON(400, data)
	}

	data["data"] = "start"

	return c.JSON(200, data)
}
