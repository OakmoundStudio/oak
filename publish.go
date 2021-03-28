//+build !js

package oak

import (
	"image/draw"

	"github.com/oakmound/shiny/screen"
)

var (
	drawLoopPublishDef = func(c *Controller, tx screen.Texture) {
		tx.Upload(zeroPoint, c.winBuffer, c.winBuffer.Bounds())
		c.windowControl.Scale(c.windowRect, tx, tx.Bounds(), draw.Src)
		c.windowControl.Publish()
	}
)
