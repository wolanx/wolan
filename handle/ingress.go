package handle

import (
	"github.com/labstack/echo"
	"github.com/zx5435/wolan/docker"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

func IngressStart(c echo.Context) error {
	d := docker.NewWDocker()
	d.RunOne("wolan-ingress", &container.Config{
		Image: "zx5435/wolan:ingress",
	}, &container.HostConfig{
		PortBindings: docker.QuickPortMap([]string{
			"1313:80",
		}),
	}, &network.NetworkingConfig{
	})

	data := make(map[string]interface{})
	data["data"] = "start"

	return c.JSON(200, data)
}
