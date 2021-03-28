package render

import "sync"

var (
	debugMap  = make(map[string]Renderable)
	debugLock = sync.RWMutex{}
)

// UpdateDebugMap stores a renderable under a name in a package global map.
// this is used by some built in debugConsole helper functions.
func UpdateDebugMap(rName string, r Renderable) {
	debugLock.Lock()
	debugMap[rName] = r
	debugLock.Unlock()
}

// GetDebugRenderable returns whatever renderable is stored under the input
// string, if any.
func GetDebugRenderable(rName string) (Renderable, bool) {
	debugLock.RLock()
	r, ok := debugMap[rName]
	debugLock.RUnlock()	
	if r == nil {
		return nil, false
	}
	return r.(Renderable), ok
}
