package fileutil

import (
	"bytes"
	"io"
	"path/filepath"

	"github.com/oakmound/oak/v2/dlog"
)

// ReadFile replaces ioutil.ReadFile, trying to use the BinaryFn if it exists.
func ReadFile(file string) ([]byte, error) {
	if BindataFn != nil {
		rel, err := filepath.Rel(wd, file)
		if err != nil {
			dlog.Warn(err)
			// Try the relative path by itself when we can't form an absolute path
			rel = file
		}
		return BindataFn(rel)
	}
	f, err := OpenOS(file)
	if err != nil {
		return nil, err
	}
	// It's a good but not certain bet that FileInfo will tell us exactly how much to
	// read, so let's try it but be prepared for the answer to be wrong.
	var n int64

	if fi, err := f.Stat(); err == nil {
		// Don't preallocate a huge buffer, just in case.
		if size := fi.Size(); size < 1e9 {
			n = size
		}
	}
	// As initial capacity for readAll, use n + a little extra in case Size is zero,
	// and to avoid another allocation after Read has filled the buffer. The readAll
	// call will read into its allocated internal buffer cheaply. If the size was
	// wrong, we'll either waste some space off the end or reallocate as needed, but
	// in the overwhelmingly common case we'll get it just right.
	return readAll(f, n+bytes.MinRead)
}

func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, capacity))
	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}
