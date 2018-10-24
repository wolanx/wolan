package main

import (
	"html/template"
	"os"
)

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
