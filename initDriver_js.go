//+build js

package oak

import "github.com/oakmound/oak/v2/dlog"

func (c *Controller) initDriver(firstScene, imageDir, audioDir string) {
	dlog.Info("Init JS Driver")
	firstSceneJs = firstScene
	c.Driver(c.lifecycleLoop)
}
