package main

import (
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zx5435/wolan/config"
	_ "github.com/zx5435/wolan/config"
	"github.com/zx5435/wolan/http"
)

// git clone https://github.com/golang/crypto.git

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Echo instance
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob(config.ViewPath + "/*.html")),
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = t

	// Routes
	e.GET("/", http.Index)

	// Start server
	e.Logger.Fatal(e.Start(":2346"))

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
