package linux

import (
	"os"
)

func CreateMissingDirs(dirs ...string) error {
	for _, dir := range dirs {
		err := CreateMissingDir(dir)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateMissingDir(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
