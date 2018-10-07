package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type WolanConfig struct {
	GitUrl string `yaml:"git_url"`
}

func main() {
	fmt.Println("this wolang-client")

	fname := "/Users/kayl.zhao/go/src/github.com/zx5435/wolan/__work__/config/cicdcm/wolan.yaml"
	file, _ := os.Open(fname)
	fileText, _ := ioutil.ReadAll(file)

	wolanConfig := &WolanConfig{}
	yaml.Unmarshal([]byte(fileText), wolanConfig)

	fmt.Println(wolanConfig)

}
