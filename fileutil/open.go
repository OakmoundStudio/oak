package fileutil

import (
	"bytes"
	"io"
	"path/filepath"

	"github.com/oakmound/oak/v2/dlog"
)

// Open is a wrapper around os.Open that will also check BindataFn to access
// embedded data. The intended use is to use the an embedding library to create an
// Asset function that matches this signature.
func Open(file string) (io.ReadCloser, error) {
	var err error
	var rel string
	var data []byte
	// Check bindata
	if BindataFn != nil {
		// It looks like we need to clean this output sometimes--
		// we get capitalization where we don't want it occasionally?
		rel, err = filepath.Rel(wd, file)
		if err != nil {
			// Just try the relative path by itself if we can't form
			// an absolute path.
			rel = file
		}
		data, err = BindataFn(rel)
		if err == nil {
			dlog.Verb("Found file in binary,", rel)
			// convert data to io.Reader
			return nopCloser{bytes.NewReader(data)}, err
		}
		dlog.Warn("File not found in binary", rel)
	}
	return OpenOS(file)
}
