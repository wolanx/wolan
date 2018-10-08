package docker

import (
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

func (d *WDocker) Stop(t *WTask) {
	log.Println("down", t.StackName)

	f := filters.NewArgs(filters.Arg("name", "hehe"))

	arr, _ := d.cli.ContainerList(d.ctx, types.ContainerListOptions{
		Filters: f,
	})

	for _, one := range arr {
		log.Printf("%#v", one)
		d.cli.ContainerStop(d.ctx, one.ID, nil)
		d.cli.ContainerRemove(d.ctx, one.ID, types.ContainerRemoveOptions{})
	}
}
