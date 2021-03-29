package move

import (
	"github.com/oakmound/oak/v2/alg/floatgeom"
	"github.com/oakmound/oak/v2/key"
	"github.com/oakmound/oak/v2/physics"
	"github.com/oakmound/oak/v2/scene"
)

// WASD moves the given mover based on its speed as W,A,S, and D are pressed
func WASD(ctx *scene.Context, mvr Mover) {
	TopDown(ctx, mvr, key.W, key.S, key.A, key.D)
}

// Arrows moves the given mover based on its speed as the arrow keys are pressed
func Arrows(ctx *scene.Context, mvr Mover) {
	TopDown(ctx, mvr, key.UpArrow, key.DownArrow, key.LeftArrow, key.RightAlt)
}

// TopDown moves the given mover based on its speed as the given keys are pressed
func TopDown(ctx *scene.Context, mvr Mover, up, down, left, right string) {
	delta := mvr.GetDelta()
	vec := mvr.Vec()
	spd := mvr.GetSpeed()

	delta.Zero()
	if ctx.Window.IsDown(up) {
		delta.Add(physics.NewVector(0, -spd.Y()))
	}
	if ctx.Window.IsDown(down) {
		delta.Add(physics.NewVector(0, spd.Y()))
	}
	if ctx.Window.IsDown(left) {
		delta.Add(physics.NewVector(-spd.X(), 0))
	}
	if ctx.Window.IsDown(right) {
		delta.Add(physics.NewVector(spd.X(), 0))
	}
	vec.Add(delta)
	mvr.GetRenderable().SetPos(vec.X(), vec.Y())
	sp := mvr.GetSpace()
	sp.Update(vec.X(), vec.Y(), sp.GetW(), sp.GetH())
}

// CenterScreenOn will cause the screen to center on the given mover, obeying
// viewport limits if they have been set previously
func CenterScreenOn(ctx *scene.Context, mvr Mover) {
	vec := mvr.Vec()
	ctx.Window.SetScreen(
		int(vec.X())-ctx.Window.Width()/2,
		int(vec.Y())-ctx.Window.Height()/2,
	)
}

// Limit restricts the movement of the mover to stay within a given rectangle
func Limit(mvr Mover, rect floatgeom.Rect2) {
	vec := mvr.Vec()
	w, h := mvr.GetRenderable().GetDims()
	wf := float64(w)
	hf := float64(h)
	if vec.X() < rect.Min.X() {
		vec.SetX(rect.Min.X())
	} else if vec.X() > rect.Max.X()-wf {
		vec.SetX(rect.Max.X() - wf)
	}
	if vec.Y() < rect.Min.Y() {
		vec.SetY(rect.Min.Y())
	} else if vec.Y() > rect.Max.Y()-hf {
		vec.SetY(rect.Max.Y() - hf)
	}
}
