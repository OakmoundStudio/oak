package render

import (
	"image"
	"path/filepath"
	"sync"

	"github.com/oakmound/oak/v2/fileutil"
)

var (
	// Form ...main/core.go/assets/images,
	// the image directory.
	wd, _ = fileutil.Getwd()
	dir   = filepath.Join(
		wd,
		"assets",
		"images")

	imageLock    = sync.RWMutex{}
	loadedImages = make(map[string]*image.RGBA)

	sheetLock    = sync.RWMutex{}
	loadedSheets = make(map[string]*Sheet)

	// move to some batch load settings
	defaultPad = 0
)

func subImage(rgba *image.RGBA, x, y, w, h int) *image.RGBA {
	out := image.NewRGBA(image.Rect(0, 0, w, h))
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			out.Set(i, j, rgba.At(x+i, y+j))
		}
	}
	return out
}

// SetAssetPaths sets the directories that files are loaded from when using
// the LoadSprite utility (and others). Oak will call this with SetupConfig.Assets
// joined with SetupConfig.Images after Init.
func SetAssetPaths(imagedir string) {
	dir = imagedir
}

// UnloadAll resets the cached set of loaded sprites and sheets to empty.
func UnloadAll() {
	imageLock.Lock()
	sheetLock.Lock()
	loadedImages = make(map[string]*image.RGBA)
	loadedSheets = make(map[string]*Sheet)
	sheetLock.Unlock()
	imageLock.Unlock()
}
