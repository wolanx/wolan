package main

import (
	"os"

	"github.com/elastic/beats/v7/filebeat/cmd"
	inputs "github.com/elastic/beats/v7/filebeat/input/default-inputs"
	_ "github.com/zx5435/wolan/wolan-logging/loki"
)

// go run main.go -e -c 2.yml
func main() {
	if err := cmd.Filebeat(inputs.Init, cmd.FilebeatSettings()).Execute(); err != nil {
		os.Exit(1)
	}
}
