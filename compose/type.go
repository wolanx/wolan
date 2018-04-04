package compose

import "encoding/json"

func (this *Config) String() string {
	jsonStr, _ := json.MarshalIndent(this, "", "  ")
	return string(jsonStr)
}

type Config struct {
	Version  string
	Services map[string]Service
}

type Service struct {
	Image string
}
