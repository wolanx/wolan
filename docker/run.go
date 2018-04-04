package docker

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"os"
	"io/ioutil"
	"github.com/zx5435/wolan/compose"
)

func FileGetConnents(path string) string {
	file, _ := os.Open(path)
	fileText, _ := ioutil.ReadAll(file)

	return string(fileText)
}

func (this *WDocker) Deploy() {
	fmt.Println("WDocker::Deploy")
	fmt.Println(wCenter.Config)

	composeConfig := compose.Parse(FileGetConnents(wCenter.WorkDir + "/" + wCenter.Config.DockerCompose))
	fmt.Println(composeConfig)

	// 部署多个 container
	for name, containerConfig := range composeConfig.Services {
		fmt.Println(name, containerConfig)

		resp, err := this.cli.ContainerCreate(this.ctx, containerConfig, nil, nil, "cdemo_"+name+"_1")
		if err != nil {
			panic(err)
		}

		if err := this.cli.ContainerStart(this.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
			panic(err)
		}

		fmt.Println(resp.ID)
	}
}
