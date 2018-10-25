package ingress

import (
	"os"
	"path/filepath"
	"html/template"
	"os/exec"
	"fmt"
	"flag"
	"sync"
	"runtime"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var (
	exitMu     sync.Mutex
	exitStatus = 0
)

func MkdirAll(dir string, perm os.FileMode) error {
	LogInfoNum(2)("mkdir -p ", dir)
	return os.MkdirAll(dir, perm)
}

func sameDir(filename string, perm os.FileMode) error {
	dir := filepath.Dir(filename)
	return MkdirAll(dir, perm)
}

func NginxReload() error {
	cmd := exec.Command("/bin/sh", "-c", "nginx -s reload")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func writeTpl(tpl *template.Template, fp string, data interface{}) error {

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
		fmt.Fprintf(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func SetExitStatus(n int) {
	exitMu.Lock()
	if exitStatus < n {
		exitStatus = n
	}
	exitMu.Unlock()
}

func Exit() {
	os.Exit(exitStatus)
}

func CutName(a string) string {
	return strings.Replace(a, "/go/src/github.com/zx5435/wolan/", "", 1)
}

func LogInfo(args ...interface{}) {
	_, file, no, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"file": CutName(file) + ":" + strconv.Itoa(no),
	}).Info(args...)
}

func LogWarn(args ...interface{}) {
	_, file, no, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"file": CutName(file) + ":" + strconv.Itoa(no),
	}).Warn(args...)
}

func LogInfoNum(n int) func(args ...interface{}) {
	_, file, no, _ := runtime.Caller(n)
	return log.WithFields(log.Fields{
		"file": CutName(file) + ":" + strconv.Itoa(no),
	}).Info
}

func LogInfof(a string, args ...interface{}) {
	_, file, no, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"file": CutName(file) + ":" + strconv.Itoa(no),
	}).Infof(a, args...)
}

func Errorf(format string, args ...interface{}) {
	_, file, no, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"file": CutName(file) + ":" + strconv.Itoa(no),
	}).Errorf(format, args...)
	SetExitStatus(1)
}

func Fatalf(format string, args ...interface{}) {
	_, file, no, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"file": CutName(file) + ":" + strconv.Itoa(no),
	}).Errorf(format, args...)
	Exit()
}
