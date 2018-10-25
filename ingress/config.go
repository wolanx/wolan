package ingress

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"golang.org/x/crypto/acme"
	"log"
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
	configDir = os.Getenv("NGX_CONFIG")
	directoryURL = os.Getenv("NGX_DIRECTORY_URL")
	resourceURL = os.Getenv("NGX_RESOURCE")
	siteConfDir = os.Getenv("NGX_SITE_CONFIG")
	siteRootDir = os.Getenv("NGX_SITE_ROOT")

	if configDir == "" {
		configDir = "/root/ngxpkg"
	}

	if directoryURL == "" {
		directoryURL = "https://acme-v01.api.letsencrypt.org/directory"
		directoryURL = "https://acme-staging.api.letsencrypt.org/directory"
		log.Println(directoryURL)
	}

	if resourceURL == "" {
		resourceURL = "/go/src/github.com/zx5435/wolan/tpl/ingress/rc/"
	}

	if siteConfDir == "" {
		siteConfDir = "/etc/nginx/conf.d"
	}

	if siteRootDir == "" {
		siteRootDir = "/usr/share/nginx/html"
	}

	allowRenewDays, err := strconv.Atoi(os.Getenv("NGX_ALLOW_RENEW_DAYS"))

	if err != nil {
		allowRenewDays = 30
	}

	if allowRenewDays < 7 {
		allowRenewDays = 7
	}

	if allowRenewDays > 30 {
		allowRenewDays = 30
	}
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
