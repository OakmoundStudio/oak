//+build js

package oak

import (
	"io"
)

var (
	viewportLocked = false
)

func AddCommand(s string, fn func([]string))                     {}
func defaultDebugConsole()                                       {}
func debugConsole(resetCh, skipScene chan bool, input io.Reader) {}
