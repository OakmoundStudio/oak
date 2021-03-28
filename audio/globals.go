//+build !js

package audio

import (
	"sync"

	"github.com/200sc/klangsynthese/font"
)

var (
	loadedLock sync.RWMutex
	loaded     = make(map[string]Data)
	// DefFont is the font used for default functions. It can be publicly
	// modified to apply a default font to generated audios through def
	// methods. If it is not modified, it is a font of zero filters.
	DefFont = font.New()
)
