package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func fetchResource(filename string) ([]byte, error) {
	if filename == "" {
		return nil, errors.New("ngx: empty resource name")
	}

	dir := filepath.Join(configDir, "rc")

	if err := MkdirAll(dir, 0700); err != nil {
		return nil, err
	}

	fp := filepath.Join(dir, filename)

	if _, err := os.Stat(fp); os.IsNotExist(err) {

		file, _ := os.Open(resourceURL + filename)

		bytes, err := ioutil.ReadAll(file)

		if err != nil {
			return nil, err
		}

		if err := ioutil.WriteFile(fp, bytes, 0600); err != nil {
			return nil, err
		}

		return bytes, nil
	}

	bytes, err := ioutil.ReadFile(fp)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func writeResource(filename string) error {
	if filename == "" {
		return errors.New("ngx: empty resource name")
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := sameDir(filename, 0700); err != nil {
			return err
		}

		bytes, err := fetchResource(filepath.Base(filename))

		if err != nil {
			return err
		}

		return ioutil.WriteFile(filename, bytes, 0600)

	}

	return nil
}
