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
	//return

	wDocker := docker.NewWDocker()
	//wDocker.PullImg()
	wDocker.Deploy()

}
