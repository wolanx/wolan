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
	"net"

	"golang.org/x/crypto/acme"
	"errors"
	"github.com/zx5435/wolan/src/log"
)

func LoadTpl() (*template.Template, *template.Template) {
	conf, err := fetchResource(siteConfFile)
	if err != nil {
		log.Fatalf("read conf: %v", err)
	}
	index, err := fetchResource(siteIndexFile)
	if err != nil {
		log.Fatalf("read index: %v", err)
	}
	confTpl, err := template.New("siteConf").Parse(string(conf))
	if err != nil {
		log.Fatalf("parse conf: %v", err)
	}
	indexTpl, err := template.New("siteIndex").Parse(string(index))
	if err != nil {
		log.Fatalf("parse index: %v", err)
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
			log.Fatalf("%s root: %v", domainConfPath, err)
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
				log.Warnf("%s conf: %v", domainConfPath, err)
				continue
			} else {
				log.Fatalf("%s conf: %v", domain, err)
			}
		}

		domains = append(domains, domain)

		if err := MkDirAll(domainPublicDir, 0755); err != nil {
			log.Fatalf("%s public: %v", domain, err)
		}

		if err := writeTpl(indexTpl, domainIndexPath, data); err != nil {
			log.Infof("%s index: %v", domain, err)
		}
	}
	return domains
}

func mkDir() {
	if err := MkDirAll(configDir, 0700); err != nil {
		log.Fatalf("config dir: %v", err)
	}
	if err := MkDirAll(confDir, 0700); err != nil {
		log.Fatalf("site conf dir: %v", err)
	}
	if err := MkDirAll(wwwDir, 0755); err != nil {
		log.Fatalf("site root dir: %v", err)
	}
}

func RunNew(args []string) error {
	log.Debug(AcmeURL)
	log.Debug(configDir)

	if len(args) == 0 {
		return errors.New("args = 0")
	}

	return nil

	mkDir()
	confTpl, indexTpl := LoadTpl()
	domains := arg2dm(args, confTpl, indexTpl)

	var client *acme.Client

	if len(domains) > 0 {
		if err := NginxReload(); err != nil {
			log.Fatalf("nginx: %v", err)
		}
		time.Sleep(time.Second * 3)

		accountKey, err := anyKey(filepath.Join(configDir, accountKeyFile))
		if err != nil {
			log.Fatalf("account key: %v", err)
		}

		client = &acme.Client{
			Key:          accountKey,
			DirectoryURL: AcmeURL,
		}

		if _, err := readConfig(); os.IsNotExist(err) {
			if err := register(client); err != nil {
				log.Fatalf("register: %v", err)
			}
		}
	}

	certExpiry := 365 * 24 * time.Hour

	for _, domain := range domains {
		ipArr, err := net.LookupIP(domain)
		log.Debug(domain, ipArr)

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
			log.Fatalf("%s conf: %v", domain, err)
		}

		conf, err := parseSiteConf(domainConfPath)
		if err != nil {
			log.Fatalf("%s conf: %v", domain, err)
		}
		log.Debug("%#v\n", conf)
		if err := writeResource(conf.SslDHParam); err != nil {
			log.Warnf("dhparam: %v", err)
		}
		if err := writeResource(conf.SslSessionTicketKey); err != nil {
			log.Warnf("ticket: %v", err)
		}
		if err := writeResource(conf.SslTrustedCertificate); err != nil {
			log.Warnf("ocsp %s: %v", domain, err)
		}

		req := &x509.CertificateRequest{
			Subject: pkix.Name{CommonName: domain},
		}

		dnsNames := []string{domain}
		//dnsNames := []string{domain, "www.zx5435.com", "x.test.zx5435.com"}

		for _, cert := range conf.Certificates {
			if err := sameDir(cert.privkey, 0700); err != nil {
				log.Fatalf("dir: %v", err)
			}

			privateKey, err := anyKey(cert.privkey)
			if err != nil {
				log.Fatalf("private.pem: %v", err)
			}

			req.DNSNames = dnsNames

			csr, err := x509.CreateCertificateRequest(rand.Reader, req, privateKey)
			if err != nil {
				log.Fatalf("csr: %v", err)
			}

			for _, dnsName := range dnsNames {
				ctx, cxl := context.WithTimeout(context.Background(), 10*time.Minute)

				if err := authz(ctx, client, domainPublicDir, dnsName); err != nil {
					log.Debugf("%+v\n", err.Error())
					log.Fatalf("%s", dnsName)
				}
				cxl()
			}

			ctx, cxl := context.WithTimeout(context.Background(), 30*time.Minute)
			defer cxl()
			certs, _, err := client.CreateCert(ctx, csr, certExpiry, true)
			if err != nil {
				log.Fatalf("cert: %v", err)
			}

			// w fullchain.pem
			var fullBody []byte
			for _, b := range certs {
				b = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: b})
				fullBody = append(fullBody, b...)
			}
			if err := ioutil.WriteFile(cert.fullchain, fullBody, 0644); err != nil {
				log.Fatalf("cert: %v", err)
			}
			log.Info("gen ", cert.fullchain)
		}

		if err := NginxReload(); err != nil {
			log.Fatalf("nginx: %v", err)
		}
	}
	return nil
}
