package config

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func Generate(input, output string, data interface{}) error {
	return filepath.Walk(input, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		target := strings.Replace(path, input, output, 1)
		if info.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		t, err := template.ParseFiles(path)
		if err != nil {
			return err
		}
		return write(target, t, data)
	})
}

func write(output string, t *template.Template, data interface{}) error {
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()
	return t.Execute(f, data)
}
