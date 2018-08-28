package main

import (
	"log"

	"github.com/zx5435/wolan/docker"
	"github.com/zx5435/wolan/wolan"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	wCenter := wolan.NewWCenter()
	wCenter.Run()
	// step.1 预准备
	//wCenter.GetCode()
	//wCenter.DoBuild()
	//wCenter.PushImage()
	// step.2 调度

	return

	wDocker := docker.NewWDocker()
	wDocker.PullImg()
	wDocker.Deploy()
	wDocker.Down()
}
