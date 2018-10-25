package main

import (
	"flag"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	var action string
	var domains Domains

	flag.StringVar(&action, "s", "", "new|renew")
	flag.Var(&domains, "d", "www.example.com")

	flag.Parse()

	switch action {
	case "new":
		log.Println(action, domains)
		runNew(domains)
	case "renew":
		log.Println(action, domains)
		runRenew(domains)
	default:
		usageAndExit("-s cannot be new|renew.")
	}
}

type Domains []string

func (i *Domains) String() string {
	return "my string representation"
}

func (i *Domains) Set(value string) error {
	*i = append(*i, value)
	return nil
}
