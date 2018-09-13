package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/zx5435/wolan/config"
	"github.com/zx5435/wolan/http"
)

// git clone https://github.com/golang/crypto.git

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", http.Index)

	// Start server
	e.Logger.Fatal(e.Start(":23456"))

	//wCenter := wolan.NewWCenter()
	//wCenter.Run()
	//// step.1 预准备
	//wCenter.GetCode()
	////wCenter.DoBuild()
	////wCenter.PushImage()
	//// step.2 调度
	//
	////return
	//
	//wDocker := docker.NewWDocker()
	////wDocker.Pull()
	//wDocker.Stop()
	////wDocker.Deploy()
}
