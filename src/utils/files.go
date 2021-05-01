package utils

import (
	"errors"
	"io/ioutil"
	"os"
)

func IsDirectoryExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func SaveFile(fileName string, data []byte) error {
	err := ioutil.WriteFile(fileName, data, 0644)
	return err
}

func LoadFile(fileName string) ([]byte, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		err = errors.New("could not load configuration file")
	}
	return data, err
}
