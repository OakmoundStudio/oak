package main

import (
	"image/color"

	oak "github.com/oakmound/oak/v3"
	"github.com/oakmound/oak/v3/entities"
	"github.com/oakmound/oak/v3/render"
	"github.com/oakmound/oak/v3/scene"
)

func main() {
	oak.AddScene("platformer", scene.Scene{Start: func(*scene.Context) {
		char := entities.NewMoving(100, 100, 16, 32,
			render.NewColorBox(16, 32, color.RGBA{255, 0, 0, 255}),
			nil, 0, 0)

		render.Draw(char.R)
	}})
	oak.Init("platformer")
}
