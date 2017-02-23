package render

import (
	"image"
	"image/color"
	"image/draw"
)

// The Compound type is intended for use to easily swap between multiple
// renderables that are drawn at the same position on the same layer.
// A common use case for this would be a character entitiy who switches
// their animation based on how they are moving or what they are doing.
//
// The Compound type removes the need to repeatedly draw and undraw elements
// of a character, which has a tendency to leave nothing drawn for a draw frame.
type Compound struct {
	LayeredPoint
	subRenderables map[string]Modifiable
	curRenderable  string
}

func NewCompound(start string, m map[string]Modifiable) *Compound {
	return &Compound{
		subRenderables: m,
		curRenderable:  start,
	}
}

func (c *Compound) Add(k string, v Modifiable) {
	c.subRenderables[k] = v
}

func (c *Compound) Set(k string) {
	if _, ok := c.subRenderables[k]; !ok {
		panic("Unknown renderable for string " + k + " on compound")
	}
	c.curRenderable = k
}
func (c *Compound) GetSub(s string) Modifiable {
	return c.subRenderables[s]
}
func (c *Compound) Get() string {
	return c.curRenderable
}

func (c *Compound) IsInterruptable() bool {
	switch t := c.subRenderables[c.curRenderable].(type) {
	case *Animation:
		return t.Interruptable
	case *Sequence:
		return t.Interruptable
	case *Reverting:
		return t.IsInterruptable()
	case *Compound:
		return t.IsInterruptable()
	}
	return true
}

func (c *Compound) IsStatic() bool {
	switch c.subRenderables[c.curRenderable].(type) {
	case *Animation, *Sequence:
		return false
	case *Reverting:
		return c.subRenderables[c.curRenderable].(*Reverting).IsStatic()
	case *Compound:
		return c.subRenderables[c.curRenderable].(*Compound).IsStatic()
	}
	return true
}

func (c *Compound) SetOffsets(k string, offsets Point) {
	if r, ok := c.subRenderables[k]; ok {
		r.SetPos(offsets.X, offsets.Y)
	}
}

func (c *Compound) Copy() Modifiable {
	newC := new(Compound)
	*newC = *c
	newSubRenderables := make(map[string]Modifiable)
	for k, v := range c.subRenderables {
		newSubRenderables[k] = v.Copy()
	}
	newC.subRenderables = newSubRenderables
	return newC
}

func (c *Compound) GetRGBA() *image.RGBA {
	return c.subRenderables[c.curRenderable].GetRGBA()
}

func (c *Compound) ApplyColor(co color.Color) Modifiable {
	for _, rend := range c.subRenderables {
		rend.ApplyColor(co)
	}
	return c
}

func (c *Compound) FillMask(img image.RGBA) Modifiable {
	for _, rend := range c.subRenderables {
		rend.FillMask(img)
	}
	return c
}

func (c *Compound) ApplyMask(img image.RGBA) Modifiable {
	for _, rend := range c.subRenderables {
		rend.ApplyMask(img)
	}
	return c
}

func (c *Compound) Rotate(degrees int) Modifiable {
	for _, rend := range c.subRenderables {
		rend.Rotate(degrees)
	}
	return c
}

func (c *Compound) Scale(xRatio float64, yRatio float64) Modifiable {
	for _, rend := range c.subRenderables {
		rend.Scale(xRatio, yRatio)
	}
	return c
}

func (c *Compound) FlipX() Modifiable {
	for _, rend := range c.subRenderables {
		rend.FlipX()
	}
	return c
}

func (c *Compound) FlipY() Modifiable {
	for _, rend := range c.subRenderables {
		rend.FlipY()
	}
	return c
}
func (c *Compound) Fade(alpha int) Modifiable {
	for _, rend := range c.subRenderables {
		rend.Fade(alpha)
	}
	return c
}

func (c *Compound) DrawOffset(buff draw.Image, xOff float64, yOff float64) {
	c.subRenderables[c.curRenderable].DrawOffset(buff, c.X+xOff, c.Y+yOff)
}

func (c *Compound) Draw(buff draw.Image) {
	c.subRenderables[c.curRenderable].DrawOffset(buff, c.X, c.Y)
}

func (c *Compound) ShiftPos(x, y float64) {
	c.SetPos(c.X+x, c.Y+y)
}

func (c *Compound) ShiftY(y float64) {
	c.SetPos(c.X, c.Y+y)
}

func (c *Compound) ShiftX(x float64) {
	c.SetPos(c.X+x, c.Y)
}

func (c *Compound) SetPos(x, y float64) {
	c.LayeredPoint.SetPos(x, y)
}

func (c *Compound) Pause() {
	switch t := c.subRenderables[c.curRenderable].(type) {
	case *Animation:
		t.Pause()
	case *Sequence:
		t.Pause()
	case *Reverting:
		t.Pause()
	}
}

func (c *Compound) Unpause() {
	switch t := c.subRenderables[c.curRenderable].(type) {
	case *Animation:
		t.Unpause()
	case *Sequence:
		t.Unpause()
	case *Reverting:
		t.Unpause()
	}
}

func (c *Compound) Revert(mod int) {
	for _, v := range c.subRenderables {
		switch t := v.(type) {
		case *Reverting:
			t.Revert(mod)
		}
	}
}
