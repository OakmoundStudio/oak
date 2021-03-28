//+build js

package oak

import (
	"image"
	"runtime"

	"github.com/oakmound/oak/v2/collision"
	"github.com/oakmound/oak/v2/dlog"
	"github.com/oakmound/oak/v2/event"
	omouse "github.com/oakmound/oak/v2/mouse"
	"github.com/oakmound/oak/v2/render"
	"github.com/oakmound/oak/v2/scene"
	"github.com/oakmound/oak/v2/timing"
	"github.com/oakmound/shiny/screen"
)

var (
	winBuffer     screen.Image
	screenControl screen.Screen
	windowControl screen.Window
	firstSceneJs  string
)

func lifecycleLoop(inScreen screen.Screen) {
	dlog.Info("Init Lifecycle")

	firstScene = firstSceneJs

	screenControl = inScreen
	var err error

	// The window buffer represents the subsection of the world which is available to
	// be shown in a window.
	dlog.Info("Creating window buffer")
	winBuffer, err = screenControl.NewImage(image.Point{ScreenWidth, ScreenHeight})
	if err != nil {
		dlog.Error(err)
		return
	}

	// The window controller handles incoming hardware or platform events and
	// publishes image data to the screen.\
	dlog.Info("Creating window controller")
	changeWindow(int32(conf.Screen.X), int32(conf.Screen.Y), ScreenWidth*conf.Screen.Scale, ScreenHeight*conf.Screen.Scale)

	go drawLoop()
	inputLoopInit()
	var prevScene string

	SceneMap.CurrentScene = "loading"

	result := new(scene.Result)

	dlog.Info("First Scene Start")

	drawCh <- true
	drawCh <- true

	dlog.Verb("Draw Channel Activated")

	schedCt := 0
	for {
		ViewPos = image.Point{0, 0}
		updateScreen(0, 0)
		useViewBounds = false

		dlog.Info("Scene Start", SceneMap.CurrentScene)
		go func() {
			dlog.Info("Starting scene in goroutine", SceneMap.CurrentScene)
			s, ok := SceneMap.GetCurrent()
			if !ok {
				dlog.Error("Unknown scene", SceneMap.CurrentScene)
				panic("Unknown scene")
			}
			s.Start(prevScene, result.NextSceneInput)
			transitionCh <- true
		}()
		sceneTransition(result)
		// Post transition, begin loading animation
		dlog.Info("Starting load animation")
		drawCh <- true
		dlog.Info("Getting Transition Signal")
		<-transitionCh
		dlog.Info("Resume Drawing")
		// Send a signal to resume (or begin) drawing
		drawCh <- true

		dlog.Info("Looping Scene")
		cont := true
		logicTicker := timing.NewDynamicTicker()
		logicTicker.SetTick(timing.FPSToDuration(FrameRate))
		scen, ok := SceneMap.GetCurrent()
		if !ok {
			dlog.Error("missing scene")
		}
		for cont {
			<-logicTicker.C
			logicHandler.Update()
			inputLoopSwitch()
			logicHandler.Flush()
			cont = scen.Loop()
			schedCt++
			if schedCt > 100 {
				schedCt = 0
				runtime.Gosched()
			}
		}
		dlog.Info("Scene End", SceneMap.CurrentScene)

		prevScene = SceneMap.CurrentScene

		// Send a signal to stop drawing
		drawCh <- true

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
		logicHandler.Reset()
		// We follow by clearing collision areas
		// because otherwise collision function calls
		// on non-entities (i.e. particles) can still
		// be triggered and attempt to access an entity
		dlog.Verb("Event Bus Reset")
		collision.Clear()
		omouse.Clear()
		event.ResetEntities()
		render.ResetDrawStack()
		render.PreDraw()
		dlog.Verb("Engine Reset")

		// Todo: Add in customizable loading scene between regular scenes

		SceneMap.CurrentScene, result = scen.End()
		// For convenience, we allow the user to return nil
		// but it gets translated to an empty result
		if result == nil {
			result = new(scene.Result)
		}
	}
}
