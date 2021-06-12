package debugstream

import (
	"io"
	"sync"

	"github.com/oakmound/oak/v3/window"
)

var (
	// DefaultCommands to attach to. TODO: init should be lazy.
	DefaultCommands *ScopedCommands
	defaultsOnce    sync.Once
)

func checkOrCreateDefaults() {
	defaultsOnce.Do(func() {
		DefaultCommands = NewScopedCommands()
	})
}

// AddCommand to the default command set.
// See ScopedCommands' AddComand.
func AddCommand(s string, usageFn func([]string) string, fn func([]string) string) error {
	checkOrCreateDefaults()
	return DefaultCommands.AddCommand(s, usageFn, fn)
}

// AddScopedCommand to the default command set.
// See ScopedCommands' AddScopedCommand.
func AddScopedCommand(scopeID int32, s string, usageFn func([]string) string, fn func([]string) string) error {
	checkOrCreateDefaults()
	return DefaultCommands.AddScopedCommand(scopeID, s, usageFn, fn)
}

// AttachToStream if possible to start consuming the stream
// and executing commands per the stored infomraiton in the ScopeCommands.
func AttachToStream(input io.Reader, output io.Writer) {
	checkOrCreateDefaults()
	DefaultCommands.AttachToStream(input, output)
}

// AddDefaultsForScope for debugging.
func AddDefaultsForScope(scopeID int32, controller interface{}) {
	checkOrCreateDefaults()
	if c, ok := controller.(window.Window); ok {
		DefaultCommands.AddDefaultsForScope(scopeID, c)
	}
}