package main

import (
	"fmt"
	"github.com/zx5435/wolan/pkg"
)

func main() {
	fmt.Println("Wolang Server")

	wCenter := new(pkg.WCenter)
	wCenter.Run()

}
