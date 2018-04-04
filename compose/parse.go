package compose

import (
	"gopkg.in/yaml.v2"
	"fmt"
)

// 解析 docker-composer.yml 的文件
// 协议 用的是 docker 内部 type
func Parse(str string) *Configs {
	fmt.Println(str)

	c := &Configs{}
	yaml.Unmarshal([]byte(str), c)

	return c
}
