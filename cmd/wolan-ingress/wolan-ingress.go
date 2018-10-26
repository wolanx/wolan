package main

import (
	"flag"
	"github.com/zx5435/wolan/ingress"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})

	var (
		env     string
		action  string
		domains Domains
	)

	flag.StringVar(&env, "env", "dev", "env")
	flag.StringVar(&action, "s", "", "new|renew")
	flag.Var(&domains, "d", "www.example.com")

	flag.Parse()

	if env != "prod" {
		ingress.AcmeURL = "https://acme-staging.api.letsencrypt.org/directory"
	} else {
		ingress.AcmeURL = "https://acme-v01.api.letsencrypt.org/directory"
	}

	switch action {
	case "new":
		ingress.LogoNum(0).Info(action, domains)
		err := ingress.RunNew(domains)
		if err != nil {
			ingress.LogoNum(0).Info(err.Error())
		}
	case "renew":
		ingress.LogoNum(0).Info(action, domains)
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
