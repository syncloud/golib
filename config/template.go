package config

import (
	"errors"
	"path/filepath"
	"fmt"
	"os"
	"text/template"
"io/fs"
)

func Generate(input, output string, data interface{}) error {
	_, err := os.Stat(output)
	if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(output, 0755)
		if err != nil {
			return err
		}
	}
	err = filepath.Walk(input, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return os.MkdirAll(path, 0755)
		}
 t, err := template.ParseFiles(path)
if err != nil {
			return err
		}
		return write(path, t, data)

	})
	
	return err
}

func write(output string, t *template.Template, data interface{}) error {
	f, err := os.Create(fmt.Sprintf("%s/%s", output, t.Name()))
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, data)

}
