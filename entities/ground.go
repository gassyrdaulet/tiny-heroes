package entities

import (
	"image/color"

	"github.com/gassyrdaulet/go-fighting-game/base"
	"github.com/gassyrdaulet/go-fighting-game/constants"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ground struct {
	*base.Entity
}

func (g *Ground) Draw(screen *ebiten.Image) {
	vector.FillRect(
		screen,
		0,
		float32(constants.ScreenH-constants.GroundHeight),
		float32(constants.ScreenW),
		float32(constants.GroundHeight),
		color.RGBA{60, 120, 60, 255},
		false,
	)
}
