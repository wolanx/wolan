package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/zx5435/wolan/config"
	"github.com/zx5435/wolan/handle"
	"regexp"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	echo.NotFoundHandler = func(c echo.Context) error {
		// render your 404 page
		return c.File("frontend/build/index.html")
		//return c.String(http.StatusNotFound, "404")
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/api/user/info", handle.Test)

	e.GET("/api/task/list", handle.List)
	e.GET("/api/task/:id", handle.Info)
	e.POST("/api/task/:id/run", handle.Run)

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: func(c echo.Context) bool {
			reg := regexp.MustCompile(`api`)
			if reg.Match([]byte(c.Request().RequestURI)) {
				return true
			} else {
				return false
			}
		},
		Root:   "frontend/build",
		Browse: true,
	}))

	// Load server
	e.Logger.Fatal(e.Start(":23456"))
}
