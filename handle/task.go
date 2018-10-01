package handle

import (
	"io/ioutil"

	"github.com/labstack/echo"
	"github.com/zx5435/wolan/config"
)

type Task struct {
	Name string `json:"name"`
}

func List(c echo.Context) error {
	tasks := []*Task{}

	files, _ := ioutil.ReadDir(config.TaskRootPath)
	for _, f := range files {
		tasks = append(tasks, &Task{
			Name: f.Name(),
		})
	}

	data := make(map[string]interface{})
	data["data"] = tasks

	return c.JSON(200, data)
}

func Info(c echo.Context) error {
	data := make(map[string]interface{})
	data["data"] = "info"

	return c.JSON(200, data)
}


func Run(c echo.Context) error {
	data := make(map[string]interface{})
	data["data"] = "start"

	return c.JSON(200, data)
}


func Stop(c echo.Context) error {
	data := make(map[string]interface{})
	data["data"] = "start"

	return c.JSON(200, data)
}


func Rm(c echo.Context) error {
	data := make(map[string]interface{})
	data["data"] = "start"

	return c.JSON(200, data)
}
