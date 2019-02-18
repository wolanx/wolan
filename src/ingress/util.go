package ingress

import (
	"os"
	"path/filepath"
	"html/template"
	"os/exec"
	"flag"
	"runtime"
	log2 "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"github.com/zx5435/wolan/src/log"
	"fmt"
)

var (
	exitStatus = 0
)

func MkDirAll(dir string, perm os.FileMode) error {
	//LogoNum(1).Info("mkdir -p ", dir)
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
	LogoNum(1).Info("WriteTpl ", fp)
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
	os.Exit(1)
}

func Exit() {
	os.Exit(exitStatus)
}

func CutName(a string) string {
	return strings.Replace(a, "/go/src/github.com/zx5435/wolan/", "", 1)
}

// log2
func LogoNum(n int) *log2.Entry {
	_, file, no, _ := runtime.Caller(1 + n)
	return log2.WithFields(log2.Fields{
		"file": CutName(file) + ":" + strconv.Itoa(no),
	})
}

func Fatalf(format string, args ...interface{}) {
	_, file, no, _ := runtime.Caller(1)
	log2.WithFields(log2.Fields{
		"file": CutName(file) + ":" + strconv.Itoa(no),
	}).Errorf(format, args...)
	Exit()
}
