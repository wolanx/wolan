package main

import (
	"net"
	"log"
)

func main() {
	a, _ := net.LookupIP("db.zx5435.com")
	log.Println(a)
}
