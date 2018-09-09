package config

import (
	"log"
	"os"
	"path/filepath"
)

var (
	WorkPath     string
	TaskRootPath string
	GitRootPath  string
	ViewPath     string
)

func init() {
	pwd, _ := os.Getwd()

	WorkPath, _ = filepath.Abs(pwd + "/__work__")
	ViewPath, _ = filepath.Abs(pwd + "/views")
	log.Println(WorkPath)

	TaskRootPath = WorkPath + "/task"
	GitRootPath = WorkPath + "/git"
}
