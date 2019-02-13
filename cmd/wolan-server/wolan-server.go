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
	//echo.NotFoundHandler = func(c echo.Context) error {
	//	reg := regexp.MustCompile("api")
	//	log.Println(c.Request().RequestURI, reg.Match([]byte(c.Request().RequestURI)))
	//	if reg.Match([]byte(c.Request().RequestURI)) {
	//		return c.String(http.StatusNotFound, `{"name":"Not Found","message":"页面未找到。"}`)
	//	} else {
	//		return c.File("frontend/build/index.html")
	//	}
	//}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.File("frontend/build/index.html")
	})

	// user
	e.GET("/api/user/info", handle.Test)
	// ingress
	e.GET("/api/ingress/start", handle.IngressStart)
	e.GET("/api/ingress/reload", handle.Test)
	// task
	e.GET("/api/task/list", handle.List)
	e.GET("/api/task/:id", handle.Info)
	e.POST("/api/task/:id/run", handle.Run)

	e.Static("/static", "frontend/build/static")

	// Load server
	e.Logger.Fatal(e.Start(":23456"))
}
