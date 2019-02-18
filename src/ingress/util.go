package ingress

import (
	"os"
	"path/filepath"
	"html/template"
	"os/exec"
	"flag"
	"github.com/zx5435/wolan/src/log"
	"fmt"
)

func MkDirAll(dir string, perm os.FileMode) error {
	return os.MkdirAll(dir, perm)
}

func sameDir(filename string, perm os.FileMode) error {
	dir := filepath.Dir(filename)
	return MkDirAll(dir, perm)
}

func NginxReload() error {
	cmd := exec.Command("/bin/sh", "-c", "nginx -s reload")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func writeTpl(tpl *template.Template, fp string, data interface{}) error {
	log.Debug("WriteTpl ", fp)
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		fn, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer fn.Close()

		return tpl.Execute(fn, data)
	}

	return os.ErrExist
}

func editTpl(tpl *template.Template, fp string, data interface{}) error {
	fn, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer fn.Close()

	return tpl.Execute(fn, data)
}

func UsageAndExit(msg string) {
	if msg != "" {
		log.Error(msg)
	}
	fmt.Println()
	flag.Usage()
	fmt.Println(`Demo:
  wolan-ingress -s new -d www.test.com
  wolan-ingress -env=prod -s=new -d zx5435.com`)
	os.Exit(1)
}
