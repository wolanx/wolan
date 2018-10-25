package ingress

import (
	"io/ioutil"
	"regexp"
)

var (
	confSslRegex = regexp.MustCompile(`(root|ssl_certificate|ssl_certificate_key|ssl_session_ticket_key|ssl_dhparam|ssl_trusted_certificate)\s+([a-z0-9_\-\.\/]+?);`)
)

type ngxCertificate struct {
	privkey   string
	fullchain string
}

type ngxSiteConf struct {
	Certificates          []ngxCertificate
	SslSessionTicketKey   string
	SslDHParam            string
	SslTrustedCertificate string
	DomainPublicDir       string
}

func parseSiteConf(confFilename string) (*ngxSiteConf, error) {
	text, err := ioutil.ReadFile(confFilename)

	if err != nil {
		return nil, err
	}

	matches := confSslRegex.FindAllStringSubmatch(string(text), -1)

	var conf = &ngxSiteConf{}
	var cert *ngxCertificate

	for _, match := range matches {

		if len(match) == 3 {
			key, value := match[1], match[2]

			switch key {
			case "ssl_certificate":
				if cert == nil {
					cert = &ngxCertificate{
						fullchain: value,
					}
				} else {
					cert.fullchain = value
					conf.Certificates = append(conf.Certificates, *cert)
					cert = nil
				}
			case "ssl_certificate_key":
				if cert == nil {
					cert = &ngxCertificate{
						privkey: value,
					}
				} else {
					cert.privkey = value
					conf.Certificates = append(conf.Certificates, *cert)
					cert = nil
				}
			case "ssl_session_ticket_key":
				conf.SslSessionTicketKey = value
			case "ssl_dhparam":
				conf.SslDHParam = value
			case "ssl_trusted_certificate":
				conf.SslTrustedCertificate = value
			case "root":
				if conf.DomainPublicDir == "" {
					conf.DomainPublicDir = value
				}
			}
		}

	}

	return conf, nil
}
