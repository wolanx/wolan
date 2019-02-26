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
	"github.com/zx5435/wolan/src/log"
)

func RunNew(domains []string) {
	log.Debug(AcmeURL, domains)

	if len(domains) == 0 {
		log.Fatal("no domains")
	}

	MkDirAll(configDir, 0700)
	MkDirAll(confDir, 0700)
	MkDirAll(wwwDir, 0755)

	confTpl, err := template.New("siteConf").Parse(getFile("site.conf"))
	if err != nil {
		log.Fatalf("parse conf: %v", err)
	}
	indexTpl, err := template.New("siteIndex").Parse(getFile("index.html"))
	if err != nil {
		log.Fatalf("parse index: %v", err)
	}
	add2(domains, confTpl, indexTpl)

	NginxReload()
	time.Sleep(time.Second * 2)

	accountKey, err := anyKey(filepath.Join(configDir, accountKeyFile))
	if err != nil {
		log.Fatalf("account key: %v", err)
	}

	client := &acme.Client{
		Key:          accountKey,
		DirectoryURL: AcmeURL,
	}

	if _, err := readConfig(); os.IsNotExist(err) {
		if err := register(client); err != nil {
			log.Fatalf("register: %v", err)
		}
	}

	certExpiry := 365 * 24 * time.Hour

	for _, dm := range domains {
		ipArr, err := net.LookupIP(dm)
		log.Debug(dm, ipArr)

		ngConf := filepath.Join(confDir, dm+".conf")
		dmRoot := filepath.Join(wwwDir, dm)
		dmCk := filepath.Join(dmRoot, "public")

		data := struct {
			SiteRoot string
			Domain   string
			WithSSL  bool
		}{
			SiteRoot: wwwDir,
			Domain:   dm,
			WithSSL:  true,
		}

		if err := fileEdit(confTpl, ngConf, data); err != nil {
			log.Fatalf("%s conf: %v", dm, err)
		}

		conf, err := parseSiteConf(ngConf)
		if err != nil {
			log.Fatalf("%s conf: %v", dm, err)
		}
		log.Debugf("%#v\n", conf)

		// no use
		if err := writeResource(conf.SslDHParam); err != nil {
			log.Warnf("dhparam: %v", err)
		}
		if err := writeResource(conf.SslSessionTicketKey); err != nil {
			log.Warnf("ticket: %v", err)
		}
		if err := writeResource(conf.SslTrustedCertificate); err != nil {
			log.Warnf("ocsp %s: %v", dm, err)
		}

		req := &x509.CertificateRequest{
			Subject: pkix.Name{CommonName: dm},
		}

		dnsNames := []string{dm}

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
				defer cxl()

				if err := createWellKown(ctx, client, dmCk, dnsName); err != nil {
					log.Debugf("%+v\n", err.Error())
					log.Fatalf("%s", dnsName)
				}
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

		NginxReload()
	}
	return
}

func add2(dms []string, confTpl *template.Template, indexTpl *template.Template) {
	for _, dm := range dms {
		ngConf := filepath.Join(confDir, dm+".conf")
		dmRoot := filepath.Join(wwwDir, dm)
		dmCk := filepath.Join(dmRoot, "public")
		dmCkIdx := filepath.Join(dmCk, "index.html")

		MkDirAll(dmRoot, 0755)
		MkDirAll(dmCk, 0755)

		data := struct {
			SiteRoot string
			Domain   string
			WithSSL  bool
		}{
			SiteRoot: wwwDir,
			Domain:   dm,
			WithSSL:  false,
		}

		if err := fileCreate(confTpl, ngConf, data); err != nil {
			if os.IsExist(err) {
				log.Warn(err, ngConf)
				continue
			} else {
				log.Fatal(err, dm)
			}
		}

		if err := fileCreate(indexTpl, dmCkIdx, data); err != nil {
			log.Infof("%s index: %v", dm, err)
		}
	}
}
