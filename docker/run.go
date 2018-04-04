package docker

import (
	"fmt"
	"github.com/docker/docker/api/types/container"
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

func (this *WDocker) Run() {
	fmt.Println("WDocker::Run")

	fmt.Println(wCenter.Config)

	yml := compose.Parse(FileGetConnents(wCenter.WorkDir + "/" + wCenter.Config.DockerCompose))
	fmt.Printf("%#v\n", yml)
	fmt.Println(yml)

	return

	imgName := "zx5435/cdemo-php:a"

	resp, err := this.cli.ContainerCreate(this.ctx, &container.Config{
		Image: imgName,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := this.cli.ContainerStart(this.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}
