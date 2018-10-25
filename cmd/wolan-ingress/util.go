package main

import (
	"os"
	"path/filepath"
	"log"
	"html/template"
	"os/exec"
	"fmt"
	"flag"
	"sync"
)

var (
	exitMu     sync.Mutex
	exitStatus = 0
)

func MkdirAll(dir string, perm os.FileMode) error {
	log.Println("mkdir", dir)
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

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

var logf = log.Printf

func Errorf(format string, args ...interface{}) {
	logf(format, args...)
	SetExitStatus(1)
}

func Fatalf(format string, args ...interface{}) {
	Errorf(format, args...)
	Exit()
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
