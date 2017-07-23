package render

import (
	"image"
	"image/color"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine(t *testing.T) {
	l := NewLine(0, 0, 10, 10, color.RGBA{255, 255, 255, 255})
	rgba := l.GetRGBA()
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if x == y {
				assert.Equal(t, rgba.At(x, y), color.RGBA{255, 255, 255, 255})
			} else {
				assert.Equal(t, rgba.At(x, y), color.RGBA{0, 0, 0, 0})
			}
		}
	}
	l = NewLine(0, 0, 0, 0, color.RGBA{255, 255, 255, 255})
	rgba = l.GetRGBA()
	rgba2 := image.NewRGBA(image.Rect(0, 0, 1, 1))
	rgba2.Set(0, 0, color.RGBA{255, 255, 255, 255})
	assert.Equal(t, rgba, rgba2)

	l = NewLine(0, 0, 0, 5, color.RGBA{255, 255, 255, 255})
	rgba = l.GetRGBA()
	rgba2 = image.NewRGBA(image.Rect(0, 0, 1, 5))
	for y := 0; y < 5; y++ {
		rgba2.Set(0, y, color.RGBA{255, 255, 255, 255})
	}
	assert.Equal(t, rgba, rgba2)
}

func TestThickLine(t *testing.T) {
	l := NewThickLine(0, 0, 10, 10, color.RGBA{255, 255, 255, 255}, 1)
	rgba := l.GetRGBA()
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if math.Abs(float64(x)-float64(y)) <= 2 {
				assert.Equal(t, rgba.At(x, y), color.RGBA{255, 255, 255, 255})
			} else {
				assert.Equal(t, rgba.At(x, y), color.RGBA{0, 0, 0, 0})
			}
		}
	}
}

func TestDrawLineOnto(t *testing.T) {
	l := NewLine(0, 0, 10, 10, color.RGBA{255, 255, 255, 255})
	rgba := l.GetRGBA()
	// See height addition in line
	rgba2 := image.NewRGBA(image.Rect(0, 0, 10, 11))
	DrawLineOnto(rgba2, 0, 0, 10, 10, color.RGBA{255, 255, 255, 255})
	assert.Equal(t, rgba, rgba2)
}
