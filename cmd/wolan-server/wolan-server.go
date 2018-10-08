package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/zx5435/wolan/config"
	"github.com/zx5435/wolan/handle"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	//wTask := wolan.NewwTask()
	//wTask.Load()
	//// step.1 预准备
	//wTask.GetCode()
	//wTask.DoBuild()
	//wTask.PushImage()
	//// step.2 调度
	//
	////return
	//
	//wDocker := docker.NewWDocker()
	//wDocker.Pull()
	//wDocker.Stop()
	//wDocker.Deploy()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", handle.List)
	e.GET("/api/user/info", handle.List)

	e.GET("/api/task/list", handle.List)
	e.GET("/api/task/:id", handle.Info)
	e.POST("/api/task/:id/run", handle.Run)

	// Load server
	e.Logger.Fatal(e.Start(":23456"))
}
