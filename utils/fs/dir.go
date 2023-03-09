package fs

import (
	"errors"
	"os"
)

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	var exist bool
	for _, v := range dirs {
		exist, err = PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return err
}
