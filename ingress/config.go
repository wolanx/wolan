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
	configDir      string
	directoryURL   string
	allowRenewDays int
	resourceURL    string
	siteConfDir    string
	siteRootDir    string
)

func init() {
	configDir = "/root/ngxpkg"
	resourceURL = "/go/src/github.com/zx5435/wolan/tpl/ingress/rc/"
	siteConfDir = "/etc/nginx/conf.d"
	siteRootDir = "/usr/share/nginx/html"

	//directoryURL = "https://acme-v01.api.letsencrypt.org/directory"
	directoryURL = "https://acme-staging.api.letsencrypt.org/directory"

	allowRenewDays = 30
}

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
