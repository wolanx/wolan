package handle

import (
	"io/ioutil"

	"github.com/labstack/echo"
	"github.com/zx5435/wolan/config"
	"github.com/zx5435/wolan/docker"
)

type Task struct {
	Name string `json:"name"`
}

func List(c echo.Context) error {
	var tasks []*Task

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

type JsonQuery struct {
	Name string `json:"name"`
}

func Info(c echo.Context) error {
	json := new(JsonQuery)
	c.Bind(json)
	task := docker.GetTask(json.Name)

	data := make(map[string]interface{})
	data["data"] = task

	return c.JSON(200, data)
}

func Run(c echo.Context) error {
	json := new(JsonQuery)
	c.Bind(json)

	task := docker.GetTask(json.Name)

	data := make(map[string]interface{})
	data["data"] = "success"

	err := task.Deploy()
	if err != nil {
		data["message"] = err.Error()
		return c.JSON(400, data)
	}

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
