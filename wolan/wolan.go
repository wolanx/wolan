package wolan

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/zx5435/wolan/config"
	"github.com/zx5435/wolan/util"
	"gopkg.in/yaml.v2"
)

type WolanConfig struct {
	Version       string `yaml:"version"`
	Name          string `yaml:"name"`
	Git           git
	DockerCompose string `yaml:"docker-compose"`
}

func (this *WolanConfig) String() string {
	jsonStr, _ := json.MarshalIndent(this, "", "  ")
	return string(jsonStr)
}

type git struct {
	Url    string `yaml:"url"`
	Branch string `yaml:"branch"`
}

type WCenter struct {
	Config    *WolanConfig
	StackName string
	WorkDir   string
}

func quickRun(command string, workDir string) {
	args := strings.Split(command, " ")
	log.Println(args)

	cmd := exec.Command(args[0], args[1:]...)
	if workDir != "" {
		cmd.Dir = workDir
	}

	out, err := cmd.CombinedOutput()
	log.Println("cmd out:", string(out))

	if err != nil {
		panic(err)
	}
}

var wCenter *WCenter

func NewWCenter() *WCenter {
	if wCenter != nil {
		return wCenter
	}

	wCenter = &WCenter{}
	return wCenter
}

func (this *WCenter) Run() {
	yamlFilename := config.TaskRootPath + "/app-2/wolan.yaml"
	file, _ := os.Open(yamlFilename)
	fileText, _ := ioutil.ReadAll(file)

	wolanConfig := &WolanConfig{}
	yaml.Unmarshal([]byte(fileText), wolanConfig)
	this.Config = wolanConfig

	hashName := wolanConfig.Name // TODO
	this.WorkDir = config.GitRootPath + "/" + hashName
	this.StackName = wolanConfig.Name
}

// clone code
func (this *WCenter) GetCode() {
	log.Println("step::GetCode")

	var cmd *exec.Cmd

	if ok, _ := util.PathExists(this.WorkDir); ok {
		cmd = exec.Command("git", "pull")
		cmd.Dir = this.WorkDir
	} else {
		cmd = exec.Command("git", "clone", this.Config.Git.Url, this.WorkDir)
	}

	out, err := cmd.CombinedOutput()
	log.Println("cmd out:", string(out))

	if err != nil {
		panic(err)
	}
}

// 构建
func (this *WCenter) DoBuild() {
	log.Println("step::DoBuild")

	quickRun("make build-a", this.WorkDir)
}

// 推送image
func (this *WCenter) PushImage() {
	log.Println("step::PushImage")

	// TODO
}
