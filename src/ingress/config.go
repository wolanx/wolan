package ingress

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"golang.org/x/crypto/acme"
)

const (
	accountFile    = "account.json"
	accountKeyFile = "account.ecdsa.pem"
	siteConfFile   = "site.conf"
	siteIndexFile  = "index.html"

	rsaPrivateKey = "RSA PRIVATE KEY"
	ecPrivateKey  = "EC PRIVATE KEY"
)

var (
	allowRenewDays = 30
	AcmeURL        string
	configDir      = "/root/ngxpkg"
	tplDir         = "/go/src/github.com/zx5435/wolan/tpl/ingress/rc/"
	confDir        = "/etc/nginx/conf.d"
	wwwDir         = "/usr/share/nginx/html"
)

type userConfig struct {
	acme.Account
}

func readConfig() (*userConfig, error) {
	b, err := ioutil.ReadFile(filepath.Join(configDir, accountFile))
	if err != nil {
		return nil, err
	}
	uc := &userConfig{}
	if err := json.Unmarshal(b, uc); err != nil {
		return nil, err
	}

	return uc, nil
}

func writeConfig(uc *userConfig) error {
	b, err := json.MarshalIndent(uc, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(configDir, accountFile), b, 0600)
}
