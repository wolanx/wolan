package compose

import (
	"gopkg.in/yaml.v2"
	"fmt"
)

func Parse(str string) *Config {
	fmt.Println(str)

	c := &Config{}
	yaml.Unmarshal([]byte(str), c)

	return c
}
