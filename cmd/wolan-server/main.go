package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
)

type WolanConfig struct {
	GitUrl string `yaml:"git_url"`
}

func main() {
	fmt.Println("this wolang-client")

	basePath := "/Users/kayl.zhao/go/src/github.com/zx5435/wolan/__test__"
	gitBasePath := basePath + "/git"

	fname := basePath + "/config/app-1/wolan.yaml"
	file, _ := os.Open(fname)
	fileText, _ := ioutil.ReadAll(file)

	wolanConfig := &WolanConfig{}
	yaml.Unmarshal([]byte(fileText), wolanConfig)

	fmt.Println(wolanConfig.GitUrl)

	fmt.Println(gitBasePath)

	gname := "qwe"
	full_gname := gitBasePath + "/" + gname

	fmt.Println()

	if ok, _ := PathExists(full_gname); ok {
		fmt.Println("has git dir", full_gname)

		cmd := exec.Command("git", "pull")
		cmd.Dir = full_gname
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("cloning %v repo: %v\n%s", wolanConfig.GitUrl, err, out)
		}
		fmt.Println(string(out))
	} else {
		fmt.Println("no git dir")

		cmd := exec.Command("git", "clone", wolanConfig.GitUrl, full_gname)
		if out, err := cmd.CombinedOutput(); err != nil {
			fmt.Printf("cloning %v repo: %v\n%s", wolanConfig.GitUrl, err, out)
		}
	}

}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
