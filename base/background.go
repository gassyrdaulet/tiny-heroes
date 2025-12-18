package base

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type BackgroundLayer struct {
	Image  *ebiten.Image
	Scroll float64
}

type Background struct {
	Layers []*BackgroundLayer
}

func (bg *Background) Draw(screen *ebiten.Image, cam *Camera) {
	camX, _ := cam.TopLeft()

	screenW := float64(screen.Bounds().Dx())

	for _, layer := range bg.Layers {
		imgW := float64(layer.Image.Bounds().Dx())

		offsetX := -camX * (1 - layer.Scroll)
		offsetY := 0.0

		startX := math.Mod(offsetX, imgW)
		if startX > 0 {
			startX -= imgW
		}

		for x := startX; x < screenW; x += imgW {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(x, offsetY)
			screen.DrawImage(layer.Image, op)
		}
	}
}
