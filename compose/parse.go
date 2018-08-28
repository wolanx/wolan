package compose

import (
	"gopkg.in/yaml.v2"
)

// 解析 docker-composer.yml 的文件
// 协议 用的是 docker 内部 type
func Parse(str string) *Configs {
	c := &Configs{}
	yaml.Unmarshal([]byte(str), c)

	return c
}
