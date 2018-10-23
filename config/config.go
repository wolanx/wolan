package config

import (
	"log"
	"os"
	"path/filepath"
)

var (
	WorkPath        string
	TaskRootPath    string
	GitRootPath     string
	IngressRootPath string
)

func init() {
	pwd, _ := os.Getwd()

	WorkPath, _ = filepath.Abs(pwd + "/__work__")
	log.Println(WorkPath)

	TaskRootPath = WorkPath + "/task"
	GitRootPath = WorkPath + "/git"
	IngressRootPath = WorkPath + "/ingress"
}
