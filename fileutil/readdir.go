package fileutil

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/oakmound/oak/v2/dlog"
)

// ReadDir replaces ioutil.ReadDir, trying to use the BinaryDir if it exists.
func ReadDir(file string) ([]os.FileInfo, error) {
	var fis []os.FileInfo
	var err error
	var rel string
	var strs []string
	if BindataDir != nil {
		dlog.Verb("Bindata not nil, reading directory", file)
		rel, err = filepath.Rel(wd, file)
		if err != nil {
			dlog.Warn(err)
			// Just try the relative path by itself if we can't form
			// an absolute path.
			rel = file
		}
		strs, err = BindataDir(rel)
		if err == nil {
			fis = make([]os.FileInfo, len(strs))
			for i, s := range strs {
				// If the data does not contain a period, we consider it
				// a directory
				// todo: can we supply a function that will tell us this
				// so we don't make this (bad) assumption?
				fis[i] = dummyfileinfo{s, !strings.ContainsRune(s, '.')}
				dlog.Verb("Creating dummy file into for", s, fis[i])
			}
			return fis, nil
		}
		dlog.Warn(err)
	}
	return ReadDirOS(file)
}
