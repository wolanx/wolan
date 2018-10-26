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
	"time"
	"fmt"
	"net"

	"golang.org/x/crypto/acme"
	"errors"
)

func LoadTpl() (*template.Template, *template.Template) {
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
	return confTpl, indexTpl
}

func arg2dm(args []string, confTpl *template.Template, indexTpl *template.Template) []string {
	var domains []string
	for _, domain := range args {
		domainConfPath := filepath.Join(confDir, domain+".conf")
		domainRootDir := filepath.Join(wwwDir, domain)
		domainPublicDir := filepath.Join(domainRootDir, "public")
		domainIndexPath := filepath.Join(domainPublicDir, siteIndexFile)

		if err := MkDirAll(domainRootDir, 0755); err != nil {
			Fatalf("%s root: %v", domainConfPath, err)
		}

		data := struct {
			SiteRoot string
			Domain   string
			WithSSL  bool
		}{
			SiteRoot: wwwDir,
			Domain:   domain,
			WithSSL:  false,
		}

		if err := writeTpl(confTpl, domainConfPath, data); err != nil {
			if os.IsExist(err) {
				LogoNum(0).Warnf("%s conf: %v", domainConfPath, err)
				continue
			} else {
				Fatalf("%s conf: %v", domain, err)
			}
		}

		domains = append(domains, domain)

		if err := MkDirAll(domainPublicDir, 0755); err != nil {
			Fatalf("%s public: %v", domain, err)
		}

		if err := writeTpl(indexTpl, domainIndexPath, data); err != nil {
			LogoNum(0).Infof("%s index: %v", domain, err)
		}
	}
	return domains
}

func prepareDir() {
	if err := MkDirAll(configDir, 0700); err != nil {
		Fatalf("config dir: %v", err)
	}
	if err := MkDirAll(confDir, 0700); err != nil {
		Fatalf("site conf dir: %v", err)
	}
	if err := MkDirAll(wwwDir, 0755); err != nil {
		Fatalf("site root dir: %v", err)
	}
}

func RunNew(args []string) error {
	LogoNum(0).Warn(AcmeURL)

	if len(args) == 0 {
		return errors.New("args = 0")
	}

	prepareDir()
	confTpl, indexTpl := LoadTpl()
	domains := arg2dm(args, confTpl, indexTpl)

	var client *acme.Client

	if len(domains) > 0 {
		if err := NginxReload(); err != nil {
			Fatalf("nginx: %v", err)
		}
		time.Sleep(time.Second * 3)

		accountKey, err := anyKey(filepath.Join(configDir, accountKeyFile))
		if err != nil {
			Fatalf("account key: %v", err)
		}

		client = &acme.Client{
			Key:          accountKey,
			DirectoryURL: AcmeURL,
		}

		if _, err := readConfig(); os.IsNotExist(err) {
			if err := register(client); err != nil {
				Fatalf("register: %v", err)
			}
		}
	}

	certExpiry := 365 * 24 * time.Hour

	for _, domain := range domains {
		ipArr, err := net.LookupIP(domain)
		LogoNum(0).Warn(domain, " ", ipArr)

		domainConfPath := filepath.Join(confDir, domain+".conf")
		domainRootDir := filepath.Join(wwwDir, domain)
		domainPublicDir := filepath.Join(domainRootDir, "public")

		data := struct {
			SiteRoot string
			Domain   string
			WithSSL  bool
		}{
			SiteRoot: wwwDir,
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
		fmt.Printf("%#v\n", conf)
		if err := writeResource(conf.SslDHParam); err != nil {
			LogoNum(0).Warnf("dhparam: %v", err)
		}
		if err := writeResource(conf.SslSessionTicketKey); err != nil {
			LogoNum(0).Warnf("ticket: %v", err)
		}
		if err := writeResource(conf.SslTrustedCertificate); err != nil {
			LogoNum(0).Warnf("ocsp %s: %v", domain, err)
		}

		req := &x509.CertificateRequest{
			Subject: pkix.Name{CommonName: domain},
		}

		dnsNames := []string{domain}

		for _, cert := range conf.Certificates {
			if err := sameDir(cert.privkey, 0700); err != nil {
				Fatalf("dir: %v", err)
			}

			privateKey, err := anyKey(cert.privkey)
			if err != nil {
				Fatalf("private.pem: %v", err)
			}

			req.DNSNames = dnsNames

			csr, err := x509.CreateCertificateRequest(rand.Reader, req, privateKey)
			if err != nil {
				Fatalf("csr: %v", err)
			}

			for _, dnsName := range dnsNames {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)

				if err := authz(ctx, client, domainPublicDir, dnsName); err != nil {
					fmt.Printf("%+v\n", err.Error())
					Fatalf("%s", dnsName)
				}
				cancel()
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
			defer cancel()
			certs, _, err := client.CreateCert(ctx, csr, certExpiry, true)

			if err != nil {
				Fatalf("cert: %v", err)
			}

			// w fullchain.pem
			var fullBody []byte
			for _, b := range certs {
				b = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: b})
				fullBody = append(fullBody, b...)
			}
			if err := ioutil.WriteFile(cert.fullchain, fullBody, 0644); err != nil {
				Fatalf("cert: %v", err)
			}
			LogoNum(0).Info("gen ", cert.fullchain)
		}

		if err := NginxReload(); err != nil {
			Fatalf("nginx: %v", err)
		}
	}
	return nil
}
