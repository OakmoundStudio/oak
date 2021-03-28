//+build js

package oak

import "github.com/oakmound/shiny/screen"

var (
	drawLoopPublishDef = func(c *Controller, tx screen.Texture) {
		c.windowControl.Upload(zeroPoint, c.winBuffer, c.winBuffer.Bounds())
	}
)
