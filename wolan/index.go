package wolan

import "fmt"
import "os/exec"
import (
	"github.com/zx5435/wolan/util"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strings"
	"encoding/json"
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
	Config  *WolanConfig
	WorkDir string
}

func quickRun(command string, workDir string) {
	args := strings.Split(command, " ")
	fmt.Println(args)

	cmd := exec.Command(args[0], args[1:]...)
	if workDir != "" {
		cmd.Dir = workDir
	}

	out, err := cmd.CombinedOutput()
	fmt.Println("cmd out:", string(out))

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
	basePath := "/Users/kayl.zhao/go/src/github.com/zx5435/wolan/__test__"
	gitPath := basePath + "/git"

	yamlFilename := basePath + "/Config/app-1/wolan.yaml"
	file, _ := os.Open(yamlFilename)
	fileText, _ := ioutil.ReadAll(file)

	wolanConfig := &WolanConfig{}
	yaml.Unmarshal([]byte(fileText), wolanConfig)
	this.Config = wolanConfig

	hashName := "qwe" // TODO
	this.WorkDir = gitPath + "/" + hashName

	// step.1 预准备
	//this.GetCode()
	//this.DoBuild()
	this.PushImage()

	// step.2 调度
}

// clone code
func (this *WCenter) GetCode() {
	fmt.Println("step::GetCode")

	var cmd *exec.Cmd

	if ok, _ := util.PathExists(this.WorkDir); ok {
		cmd = exec.Command("git", "pull")
		cmd.Dir = this.WorkDir
	} else {
		cmd = exec.Command("git", "clone", this.Config.Git.Url, this.WorkDir)
	}

	out, err := cmd.CombinedOutput()
	fmt.Println("cmd out:", string(out))

	if err != nil {
		panic(err)
	}
}

// 构建
func (this *WCenter) DoBuild() {
	fmt.Println("step::DoBuild")

	quickRun("make build-a", this.WorkDir)
}

// 推送image
func (this *WCenter) PushImage() {
	fmt.Println("step::PushImage")

	// TODO
}
