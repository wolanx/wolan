package docker

import (
	"io"
	"os"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"fmt"
	"github.com/zx5435/wolan/wolan"
)

var wCenter *wolan.WCenter

func init() {
	fmt.Println("init")
	wCenter = wolan.NewWCenter()
}

type WDocker struct {
	cli *client.Client
	ctx context.Context
}

func NewWDocker() *WDocker {
	wDocker := new(WDocker)

	wDocker.ctx = context.Background()
	//cli, err := client.NewClientWithOpts(client.WithHost("tcp://192.168.199.7:2376"))
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	wDocker.cli = cli

	return wDocker
}

func (wDocker *WDocker) PullImg() {
	fmt.Println("WDocker::pullImg")

	imgName := "192.168.199.115:5000/cdemo-php:a"

	out, err := wDocker.cli.ImagePull(wDocker.ctx, imgName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
}
