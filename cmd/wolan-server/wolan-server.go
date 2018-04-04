package main

import (
	"github.com/zx5435/wolan/wolan"
	"github.com/zx5435/wolan/docker"
)

func main() {
	wCenter := wolan.NewWCenter()
	wCenter.Run()

	wDocker := docker.NewWDocker()
	//wDocker.PullImg()
	wDocker.Deploy()

}
