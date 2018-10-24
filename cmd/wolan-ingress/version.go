package main

import "fmt"

var (
	version = "v1.0.0"
	osarch  string // set by ldflags

	cmdVersion = &cmd{
		run:       runVersion,
		UsageLine: "version",
		Short:     "display ngx version",
		Long:      "display ngx version and build info.\n",
	}
)

func runVersion(args []string) {
	fmt.Println(version, osarch)
}
