package docker

import (
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

func (this *WDocker) Stop() {
	log.Println("down", wCenter.StackName)

	f := filters.NewArgs(filters.Arg("name", "hehe"))

	arr, _ := this.cli.ContainerList(this.ctx, types.ContainerListOptions{
		Filters: f,
	})

	for _, one := range arr {
		log.Printf("%#v", one)
		this.cli.ContainerStop(this.ctx, one.ID, nil)
		this.cli.ContainerRemove(this.ctx, one.ID, types.ContainerRemoveOptions{})
	}
}
