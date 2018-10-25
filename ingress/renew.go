package ingress

import (
	"context"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/acme"
)

func RunRenew(args []string) {
	var client *acme.Client
	certExpiry := 365 * 24 * time.Hour

	for _, domain := range args {
		domainConfPath := filepath.Join(siteConfDir, domain+".conf")

		if _, err := os.Stat(domainConfPath); os.IsNotExist(err) {
			logf("%s conf: %v", domain, err)
			continue
		}

		conf, err := parseSiteConf(domainConfPath)

		if err != nil {
			Fatalf("%s conf: %v", domain, err)
		}

		for _, cert := range conf.Certificates {

			c, err := parseCertificate(cert.fullchain)

			if err != nil {
				Fatalf("%s cert: %v", domain, err)
			}

			if !strings.Contains(c.Issuer.CommonName, "Let's Encrypt") {
				logf("%s Issuer '%s' not support acme, skip.", filepath.Base(cert.fullchain), c.Issuer.CommonName)
				continue
			}

			days := int(c.NotAfter.Sub(time.Now()).Hours() / 24)

			if days > allowRenewDays {
				logf("%s %d days valid, skip.", filepath.Base(cert.fullchain), days)
				continue
			}

			if client == nil {

				accountKey, err := anyKey(filepath.Join(configDir, accountKeyFile))

				if err != nil {
					Fatalf("account key: %v", err)
				}

				client = &acme.Client{
					Key:          accountKey,
					DirectoryURL: directoryURL,
				}

				if _, err := readConfig(); os.IsNotExist(err) {
					if err := register(client); err != nil {
						Fatalf("register: %v", err)
					}
				}
			}

			req := &x509.CertificateRequest{
				Subject:  pkix.Name{CommonName: c.Subject.CommonName},
				DNSNames: c.DNSNames,
			}

			privkey, err := memKey(cert.privkey)

			if err != nil {
				Fatalf("privkey: %v", err)
			}

			csr, err := x509.CreateCertificateRequest(rand.Reader, req, privkey)

			if err != nil {
				Fatalf("csr: %v", err)
			}

			for _, dnsName := range c.DNSNames {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)

				if err := authz(ctx, client, conf.DomainPublicDir, dnsName); err != nil {
					Fatalf("authz %s: %v", dnsName, err)
				}
				cancel()
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
			defer cancel()
			certs, _, err := client.CreateCert(ctx, csr, certExpiry, true)

			if err != nil {
				Fatalf("cert: %v", err)
			}

			var pemcert []byte
			for _, b := range certs {
				b = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: b})
				pemcert = append(pemcert, b...)
			}

			if err := writeKey(cert.privkey, privkey); err != nil {
				Fatalf("privkey: %v", err)
			}

			if err := ioutil.WriteFile(cert.fullchain, pemcert, 0644); err != nil {
				Fatalf("cert: %v", err)
			}
		}
	}

	if client != nil {
		if err := NginxReload(); err != nil {
			Fatalf("nginx: %v", err)
		}
	}
}
