package main

import (
	"log"
	"net"
)

func main() {
	a, _ := net.LookupIP("db.zx5435.com")
	log.Println(a)
}
