package main

import (
	"flag"

	"github.com/zx5435/wolan/common/log"
	"github.com/zx5435/wolan/lib/ingress"
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

	if env != "prod" {
		// test
		action = "new"
		domains = []string{"www.test.com"}
	}
	ingress.TEST = env != "prod"
	ingress.LoadConfig()

	log.Infof("action: %s, domains: %s", action, domains)

	switch action {
	case "new":
		ingress.RunNew(domains)
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
