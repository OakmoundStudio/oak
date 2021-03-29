//+build js

package oak

import (
	"image"

	"github.com/oakmound/oak/v2/collision"
	"github.com/oakmound/oak/v2/alg/intgeom"
	"github.com/oakmound/oak/v2/dlog"
	"github.com/oakmound/oak/v2/event"
	omouse "github.com/oakmound/oak/v2/mouse"
	"github.com/oakmound/oak/v2/render"
	"github.com/oakmound/oak/v2/scene"
	"github.com/oakmound/oak/v2/timing"
	"github.com/oakmound/shiny/screen"
)

var (
	firstSceneJs  string
)

func (c *Controller) lifecycleLoop(inScreen screen.Screen) {
	dlog.Info("Init Lifecycle")

	c.firstScene = firstSceneJs

	c.screenControl = inScreen
	var err error

	// The window buffer represents the subsection of the world which is available to
	// be shown in a window.
	dlog.Info("Creating window buffer")
	c.winBuffer, err = c.screenControl.NewImage(image.Point{c.Width(), c.Height()})
	if err != nil {
		dlog.Error(err)
		return
	}

	// The window controller handles incoming hardware or platform events and
	// publishes image data to the screen.\
	dlog.Info("Creating window controller")
	c.newWindow(int32(conf.Screen.X), int32(conf.Screen.Y), c.Width()*conf.Screen.Scale, c.Height()*conf.Screen.Scale)

	go c.drawLoop()
	var prevScene string

	c.SceneMap.CurrentScene = "loading"

	result := new(scene.Result)

	dlog.Info("First Scene Start")

	c.drawCh <- true
	c.drawCh <- true

	dlog.Verb("Draw Channel Activated")

	for {
		c.ViewPos = intgeom.Point2{0, 0}
		c.updateScreen(c.ViewPos)
		c.useViewBounds = false

		dlog.Info("Scene Start", c.SceneMap.CurrentScene)
		go func() {
			dlog.Info("Starting scene in goroutine", c.SceneMap.CurrentScene)
			s, ok := c.SceneMap.GetCurrent()
			if !ok {
				dlog.Error("Unknown scene", c.SceneMap.CurrentScene)
				panic("Unknown scene")
			}
			s.Start(&scene.Context{
				PreviousScene: prevScene,
				SceneInput:    result.NextSceneInput,
				DrawStack:     c.DrawStack,
				EventHandler:  c.logicHandler,
				CallerMap:     c.CallerMap,
				MouseTree:     c.MouseTree,
				CollisionTree: c.CollisionTree,
				Window:        c,
			})
			c.transitionCh <- true
		}()
		c.sceneTransition(result)
		// Post transition, begin loading animation
		dlog.Info("Starting load animation")
		c.drawCh <- true
		dlog.Info("Getting Transition Signal")
		<-c.transitionCh
		dlog.Info("Resume Drawing")
		// Send a signal to resume (or begin) drawing
		c.drawCh <- true

		dlog.Info("Looping Scene")
		cont := true
		logicTicker := timing.NewDynamicTicker()
		logicTicker.SetTick(timing.FPSToDuration(c.FrameRate))
		scen, ok := c.SceneMap.GetCurrent()
		if !ok {
			dlog.Error("missing scene")
		}
		for cont {
			<-logicTicker.C
			c.logicHandler.Update()
			c.inputLoopSwitch()
			c.logicHandler.Flush()
			cont = scen.Loop()
		}
		dlog.Info("Scene End", c.SceneMap.CurrentScene)

		prevScene = c.SceneMap.CurrentScene

		// Send a signal to stop drawing
		c.drawCh <- true

		// Reset any ongoing delays
	delayLabel:
		for {
			select {
			case timing.ClearDelayCh <- true:
			default:
				break delayLabel
			}
		}

		dlog.Verb("Resetting Engine")
		// Reset transient portions of the engine
		// We start by clearing the event bus to
		// remove most ongoing code
		c.logicHandler.Reset()
		// We follow by clearing collision areas
		// because otherwise collision function calls
		// on non-entities (i.e. particles) can still
		// be triggered and attempt to access an entity
		dlog.Verb("Event Bus Reset")
		collision.Clear()
		omouse.Clear()
		event.ResetCallerMap()
		render.ResetDrawStack()
		render.GlobalDrawStack.PreDraw()
		dlog.Verb("Engine Reset")

		// Todo: Add in customizable loading scene between regular scenes

		c.SceneMap.CurrentScene, result = scen.End()
		// For convenience, we allow the user to return nil
		// but it gets translated to an empty result
		if result == nil {
			result = new(scene.Result)
		}
	}
}
