package fileutil

var (
	// BindataFn is a function to access binary data outside of os.Open
	BindataFn func(string) ([]byte, error)
	// BindataDir is a function to access directory representations alike to ioutil.ReadDir
	BindataDir func(string) ([]string, error)
	wd, _      = Getwd()
)
