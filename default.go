package oak

import (
	"image"
	"time"

	"github.com/oakmound/oak/v2/alg/intgeom"
	"github.com/oakmound/oak/v2/event"
	"github.com/oakmound/oak/v2/render"
	"github.com/oakmound/oak/v2/scene"
)

var defaultController = NewController()

func Init(scene string) {
	defaultController.DrawStack = render.GlobalDrawStack
	defaultController.logicHandler = event.DefaultBus
	defaultController.Init(scene)
}

func AddCommand(command string, fn func([]string)) error {
	return defaultController.AddCommand(command, fn)
}

func AddScene(name string, sc scene.Scene) error {
	return defaultController.AddScene(name, sc)
}

func IsDown(key string) bool {
	return defaultController.IsDown(key)
}

func IsHeld(key string) (bool, time.Duration) {
	return defaultController.IsHeld(key)
}

func SetUp(key string) {
	defaultController.SetUp(key)
}

func SetDown(key string) {
	defaultController.SetDown(key)
}

func SetViewportBounds(rect intgeom.Rect2) {
	defaultController.SetViewportBounds(rect)
}

func ShiftScreen(x, y int) {
	defaultController.ShiftScreen(x, y)
}

func SetScreen(x, y int) {
	defaultController.SetScreen(x, y)
}

func SetFullScreen(fs bool) error {
	return defaultController.SetFullScreen(fs)
}

func SetBorderless(bs bool) error {
	return defaultController.SetBorderless(bs)
}

func ScreenShot() *image.RGBA {
	return defaultController.ScreenShot()
}

func SetLoadingRenderable(r render.Renderable) {
	defaultController.SetLoadingRenderable(r)
}

func GetBackgroundColor() image.Image {
	return defaultController.GetBackgroundColor()
}

func SetBackgroundColor(img image.Image) {
	defaultController.SetBackgroundColor(img)
}

func Width() int {
	return defaultController.Width()
}

func Height() int {
	return defaultController.Height()
}