//+build !js

package oak

import (
	"image"

	"github.com/oakmound/oak/v2/dlog"

	"github.com/oakmound/shiny/screen"
)

func (c *Controller) lifecycleLoop(s screen.Screen) {
	dlog.Info("Init Lifecycle")

	c.screenControl = s
	var err error

	// The window buffer represents the subsection of the world which is available to
	// be shown in a window.
	dlog.Info("Creating window buffer")
	c.winBuffer, err = c.screenControl.NewImage(image.Point{c.ScreenWidth, c.ScreenHeight})
	if err != nil {
		dlog.Error(err)
		return
	}

	// Next time:
	// Right here, query the backing scale factor of the physical screen
	// Apply that factor to the scale

	dlog.Info("Creating window controller")
	c.newWindow(int32(conf.Screen.X), int32(conf.Screen.Y), c.ScreenWidth*conf.Screen.Scale, c.ScreenHeight*conf.Screen.Scale)

	dlog.Info("Starting draw loop")
	go c.drawLoop()
	dlog.Info("Starting input loop")
	go c.inputLoop()

	// The quit channel represents a signal
	// for the engine to stop.
	<-c.quitCh
}
