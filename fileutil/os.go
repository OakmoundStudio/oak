//+build !js

package fileutil

import (
	"io/ioutil"
	"os"
)

func OpenOS(s string) (File, error) {
	return os.Open(s)
}

func ReadDirOS(dir string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dir)
}

func Getwd() (string, error) {
	return os.Getwd()
}
