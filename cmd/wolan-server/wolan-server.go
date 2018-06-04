package main

import (
	"github.com/zx5435/wolan/docker"
	"github.com/zx5435/wolan/wolan"
)

func main() {
	wCenter := wolan.NewWCenter()
	wCenter.Run(false)

	wDocker := docker.NewWDocker()
	//wDocker.PullImg()
	wDocker.Deploy()
}
