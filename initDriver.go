// +build !js

package oak

import (
	"os"

	"github.com/oakmound/oak/v2/dlog"
)

func (c *Controller) initDriver(firstScene, imageDir, audioDir string) {
	dlog.Info("Init Scene Loop")
	go c.sceneLoop(firstScene, conf.TrackInputChanges, conf.DisableDebugConsole)
	dlog.Info("Init Console")
	if !conf.DisableDebugConsole {
		dlog.Info("Init Console")
		go c.debugConsole(c.debugResetCh, c.skipSceneCh, os.Stdin)
	}
	dlog.Info("Init Main Driver")
	c.Driver(c.lifecycleLoop)
}
