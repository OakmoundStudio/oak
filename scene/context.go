package scene

import (
	"context"

	"github.com/oakmound/oak/v2/collision"
	"github.com/oakmound/oak/v2/event"
	"github.com/oakmound/oak/v2/render"
	"github.com/oakmound/oak/v2/window"
)

type Context struct {
	// This context will be canceled when the scene ends
	context.Context

	PreviousScene string
	SceneInput    interface{}
	Window        window.Window

	DrawStack     *render.DrawStack
	EventHandler  event.Handler
	CallerMap     *event.CallerMap
	MouseTree     *collision.Tree
	CollisionTree *collision.Tree
	// todo: ...
}