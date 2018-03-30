package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type WolanConfig struct {
	GitUrl string `yaml:"git_url"`
}

func main() {
	fmt.Println("this wolang-client")

	fname := "/Users/kayl.zhao/go/src/github.com/zx5435/wolan/__test__/config/app-1/wolan.yaml"
	file,_ := os.Open(fname)
	fileText, _ := ioutil.ReadAll(file)

	wolanConfig := &WolanConfig{}
	yaml.Unmarshal([]byte(fileText), wolanConfig)

	fmt.Println(wolanConfig)


}
