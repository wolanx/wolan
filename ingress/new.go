package ingress

import (
	"context"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"time"
	"fmt"
	"net"

	"golang.org/x/crypto/acme"
	"errors"
)

func RunNew(args []string) error {
	LogInfo(directoryURL)

	if len(args) == 0 {
		return errors.New("args = 0")
	}

	if err := MkdirAll(configDir, 0700); err != nil {
		Fatalf("config dir: %v", err)
	}
	if err := MkdirAll(siteConfDir, 0700); err != nil {
		Fatalf("site conf dir: %v", err)
	}
	if err := MkdirAll(siteRootDir, 0755); err != nil {
		Fatalf("site root dir: %v", err)
	}

	conf, err := fetchResource(siteConfFile)
	if err != nil {
		Fatalf("read conf: %v", err)
	}

	index, err := fetchResource(siteIndexFile)
	if err != nil {
		Fatalf("read index: %v", err)
	}

	confTpl, err := template.New("siteConf").Parse(string(conf))
	if err != nil {
		Fatalf("parse conf: %v", err)
	}

	indexTpl, err := template.New("siteIndex").Parse(string(index))
	if err != nil {
		Fatalf("parse index: %v", err)
	}

	var domains []string

	for _, domain := range args {
		domainConfPath := filepath.Join(siteConfDir, domain+".conf")
		domainRootDir := filepath.Join(siteRootDir, domain)
		domainPublicDir := filepath.Join(domainRootDir, "public")
		domainIndexPath := filepath.Join(domainPublicDir, siteIndexFile)

		if err := MkdirAll(domainRootDir, 0755); err != nil {
			Fatalf("%s root: %v", domainConfPath, err)
		}

		data := struct {
			SiteRoot string
			Domain   string
			WithSSL  bool
		}{
			SiteRoot: siteRootDir,
			Domain:   domain,
			WithSSL:  false,
		}

		if err := writeTpl(confTpl, domainConfPath, data); err != nil {
			if os.IsExist(err) {
				LogWarn(fmt.Sprintf("%s conf: %v", domainConfPath, err))
				continue
			} else {
				Fatalf("%s conf: %v", domain, err)
			}
		}

		domains = append(domains, domain)

		if err := MkdirAll(domainPublicDir, 0755); err != nil {
			Fatalf("%s public: %v", domain, err)
		}

		if err := writeTpl(indexTpl, domainIndexPath, data); err != nil {
			LogInfof("%s index: %v", domain, err)
		}
	}

	var client *acme.Client

	if len(domains) > 0 {
		if err := NginxReload(); err != nil {
			Fatalf("nginx: %v", err)
		}

		time.Sleep(time.Second * 5)

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

	certExpiry := 365 * 24 * time.Hour

	for _, domain := range domains {
		LogInfo(domain)
		ipArr, err := net.LookupIP(domain)
		LogInfo(ipArr)

		if err != nil {
			Fatalf("%s lookup: %v", domain, err)
		}

		domainConfPath := filepath.Join(siteConfDir, domain+".conf")
		domainRootDir := filepath.Join(siteRootDir, domain)
		domainPublicDir := filepath.Join(domainRootDir, "public")

		data := struct {
			SiteRoot string
			Domain   string
			WithSSL  bool
		}{
			SiteRoot: siteRootDir,
			Domain:   domain,
			WithSSL:  true,
		}

		if err := editTpl(confTpl, domainConfPath, data); err != nil {
			Fatalf("%s conf: %v", domain, err)
		}

		conf, err := parseSiteConf(domainConfPath)

		if err != nil {
			Fatalf("%s conf: %v", domain, err)
		}

		if err := writeResource(conf.SslDHParam); err != nil {
			Fatalf("dhparam: %v", err)
		}

		if err := writeResource(conf.SslSessionTicketKey); err != nil {
			Fatalf("ticket: %v", err)
		}

		if err := writeResource(conf.SslTrustedCertificate); err != nil {
			Fatalf("%s ocsp: %v", domain, err)
		}

		req := &x509.CertificateRequest{
			Subject: pkix.Name{CommonName: domain},
		}

		dnsNames := []string{
			domain,
		}

		wwwDomain := "www." + domain

		if wwwIPArr, err := net.LookupIP(wwwDomain); err == nil {
			if reflect.DeepEqual(ipArr, wwwIPArr) {
				dnsNames = append(dnsNames, wwwDomain)
			}
		}

		for _, cert := range conf.Certificates {
			if err := sameDir(cert.privkey, 0700); err != nil {
				Fatalf("dir: %v", err)
			}

			privkey, err := anyKey(cert.privkey)

			if err != nil {
				Fatalf("privkey: %v", err)
			}

			req.DNSNames = dnsNames

			csr, err := x509.CreateCertificateRequest(rand.Reader, req, privkey)

			if err != nil {
				Fatalf("csr: %v", err)
			}

			for _, dnsName := range dnsNames {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)

				if err := authz(ctx, client, domainPublicDir, dnsName); err != nil {
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

			if err := ioutil.WriteFile(cert.fullchain, pemcert, 0644); err != nil {
				Fatalf("cert: %v", err)
			}

		}

		if err := NginxReload(); err != nil {
			Fatalf("nginx: %v", err)
		}
	}
	return nil
}
