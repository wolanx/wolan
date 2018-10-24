package docker

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
	"github.com/zx5435/wolan/compose"
)

type WolanConfig struct {
	Version       string `yaml:"version"`
	Name          string `yaml:"name"`
	Git           git
	DockerCompose string `yaml:"docker-compose"`
	Volumes       []string
}

func (this *WolanConfig) String() string {
	jsonStr, _ := json.MarshalIndent(this, "", "  ")
	return string(jsonStr)
}

type git struct {
	Url    string `yaml:"url"`
	Branch string `yaml:"branch"`
}

type WTask struct {
	WolanYAML     *WolanConfig
	StackName     string
	GitDir        string
	TaskDir       string
	ComposeConfig *compose.Configs
}

func GetTask(name string) *WTask {
	t := &WTask{}
	t.Load(name)

	t.ComposeConfig = compose.Parse(FileGetContents(t.GitDir + "/" + t.WolanYAML.DockerCompose))

	log.Println(t.WolanYAML.Git.Url)

	return t
}

func (this *WTask) Load(name string) {
	yamlFilename := config.TaskRootPath + "/" + name + "/wolan.yaml"
	file, _ := os.Open(yamlFilename)
	fileText, _ := ioutil.ReadAll(file)

	wolanConfig := &WolanConfig{}
	yaml.Unmarshal([]byte(fileText), wolanConfig)
	this.WolanYAML = wolanConfig

	hashName := wolanConfig.Name // TODO
	this.GitDir = config.GitRootPath + "/" + hashName
	this.TaskDir = config.TaskRootPath + "/" + name
	this.StackName = wolanConfig.Name
}

// clone code
func (this *WTask) GetCode() {
	log.Println("step::GetCode")

	var cmd *exec.Cmd

	if ok, _ := util.PathExists(this.GitDir); ok {
		cmd = exec.Command("git", "pull")
		cmd.Dir = this.GitDir
	} else {
		cmd = exec.Command("git", "clone", this.WolanYAML.Git.Url, this.GitDir)
	}

	out, err := cmd.CombinedOutput()
	log.Println("cmd out:", string(out))

	if err != nil {
		panic(err)
	}
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

// 构建
func (this *WTask) DoBuild() {
	log.Println("step::DoBuild")

	quickRun("make build-a", this.GitDir)
}

// 推送image
func (this *WTask) PushImage() {
	log.Println("step::PushImage")

	// TODO
}

func (t *WTask) Deploy() error {
	log.Println("WDocker::Deploy", t.StackName)

	d := NewWDocker()

	// step.1 networks 网格
	d.CreateNet(t)

	// step.2 volumes

	// step.3 services 部署多个 container
	err := d.Run(t)
	if err != nil {
		return err
	}

	return nil
}
