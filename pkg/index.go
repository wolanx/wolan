package pkg

import "fmt"
import "os/exec"
import (
	"github.com/zx5435/wolan/util"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strings"
)

type WolanConfig struct {
	GitUrl string `yaml:"git_url"`
}

type WCenter struct {
	config  *WolanConfig
	workDir string
}

func (this *WCenter) Run() {
	basePath := "/Users/kayl.zhao/go/src/github.com/zx5435/wolan/__test__"
	gitPath := basePath + "/git"

	yamlFilename := basePath + "/config/app-1/wolan.yaml"
	file, _ := os.Open(yamlFilename)
	fileText, _ := ioutil.ReadAll(file)

	wolanConfig := &WolanConfig{}
	yaml.Unmarshal([]byte(fileText), wolanConfig)

	fmt.Printf("%#v\n\n", wolanConfig)

	hashName := "qwe" // TODO
	this.workDir = gitPath + "/" + hashName

	//this.GetCode()
	//this.DoBuild()
	this.PushImage()
}

// clone code
func (this *WCenter) GetCode() {
	fmt.Println("step::GetCode")

	var cmd *exec.Cmd

	if ok, _ := util.PathExists(this.workDir); ok {
		cmd = exec.Command("git", "pull")
		cmd.Dir = this.workDir
	} else {
		cmd = exec.Command("git", "clone", this.config.GitUrl, this.workDir)
	}

	out, err := cmd.CombinedOutput()
	fmt.Println("cmd out:", string(out))

	if err != nil {
		panic(err)
	}
}

// 构建
func (this *WCenter) DoBuild() {
	fmt.Printf("step::DoBuild")

	cmd := exec.Command("make", "build-a")
	cmd.Dir = this.workDir

	out, err := cmd.CombinedOutput()
	fmt.Println("cmd out:", string(out))

	if err != nil {
		panic(err)
	}
}

// 推送image
func (this *WCenter) PushImage() {
	fmt.Printf("step::PushImage")

	quickRun("make build-b", "")
}

func quickRun(command string, workDir string) {
	args := strings.Split(command," ")
	fmt.Println(args)
	return

	cmd := exec.Command("make", "build-a")
	if workDir != "" {
		cmd.Dir = workDir
	}

	out, err := cmd.CombinedOutput()
	fmt.Println("cmd out:", string(out))

	if err != nil {
		panic(err)
	}
}
