package compose

import (
	"encoding/json"
	"github.com/docker/docker/api/types/container"
)

func (this *Config) String() string {
	jsonStr, _ := json.MarshalIndent(this, "", "  ")
	return string(jsonStr)
}

type Config struct {
	Version string
	//Services map[string]Service
	Services map[string]*container.Config
}

// no use
type Service struct {
	Image string
}
