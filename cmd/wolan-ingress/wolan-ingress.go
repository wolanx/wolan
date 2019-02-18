package main

import (
	"flag"
	"github.com/zx5435/wolan/src/ingress"
	"github.com/zx5435/wolan/src/log"
)

func main() {
	var (
		env     string
		action  string
		domains Domains
	)

	flag.StringVar(&env, "env", "dev", "[dev|prod]")
	flag.StringVar(&action, "s", "", "[new|renew]")
	flag.Var(&domains, "d", "www.example.com")

	flag.Parse()

	if env != "prod" {
		ingress.AcmeURL = "https://acme-staging.api.letsencrypt.org/directory"
	} else {
		ingress.AcmeURL = "https://acme-v01.api.letsencrypt.org/directory"
	}

	log.Infof("action: %s, domains: %s", action, domains)

	switch action {
	case "new":
		err := ingress.RunNew(domains)
		if err != nil {
			ingress.LogoNum(0).Info(err.Error())
		}
	case "renew":
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
