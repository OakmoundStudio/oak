//+build js

package fileutil

import (
	"errors"
	"os"
)

func ReadDirOS(string) ([]os.FileInfo, error) {
	return []os.FileInfo{}, errors.New("ReadDir unsupported on JS without Bindata")
}
