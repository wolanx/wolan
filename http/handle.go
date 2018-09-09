package http

import (
	"io/ioutil"

	"github.com/labstack/echo"
	"github.com/zx5435/wolan/config"
)

// Handler
func Index(c echo.Context) error {
	tasks := []string{}

	files, _ := ioutil.ReadDir(config.TaskRootPath)
	for _, f := range files {
		tasks = append(tasks, f.Name())
	}

	data := make(map[string]interface{})
	data["arr"] = tasks

	c.Redirect(302, "//localhost:2345/")
	return c.Render(200, "index.html", data)
}
