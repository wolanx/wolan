package main

import (
	"flag"
	"github.com/zx5435/wolan/ingress"
	log "github.com/sirupsen/logrus"
)

// docker cp ./cmd/wolan-ingress/wolan-ingress wolan-ingress:/usr/bin/wolan-ingress
// docker cp ./tpl/ingress/rc wolan-ingress:/go/src/github.com/zx5435/wolan/tpl/ingress/rc
func main() {
	log.SetFormatter(&log.TextFormatter{})

	var action string
	var domains Domains

	flag.StringVar(&action, "s", "", "new|renew")
	flag.Var(&domains, "d", "www.example.com")

	flag.Parse()

	switch action {
	case "new":
		ingress.LogInfo(action, domains)
		err := ingress.RunNew(domains)
		if err != nil {
			ingress.LogInfoNum(2)(err.Error())
		}
	case "renew":
		ingress.LogInfo(action, domains)
		ingress.RunRenew(domains)
	default:
		ingress.UsageAndExit("-s cannot be new|renew.")
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
