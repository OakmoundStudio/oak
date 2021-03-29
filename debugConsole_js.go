//+build js

package oak

import (
	"io"
)

var (
	viewportLocked = false
)

func (c* Controller) AddCommand(s string, fn func([]string)) error {
	return nil
}
func defaultDebugConsole()                                         {}
func debugConsole(resetCh, skipScene chan bool, input io.Reader)   {}
