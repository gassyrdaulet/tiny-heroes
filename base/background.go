package base

import (
	"math"

	c "github.com/gassyrdaulet/go-fighting-game/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

type BackgroundLayer struct {
	Image  *ebiten.Image
	ScrollX float64
	ScrollY  float64
	StretchY bool
}

type Background struct {
	Layers []*BackgroundLayer
	BaseY    float64
}

func (bg *Background) Draw(screen *ebiten.Image, cam *Camera) {
	camX, camY := cam.TopLeft()

	screenW := float64(c.ScreenW)
	screenH := float64(c.ScreenH)

	for _, layer := range bg.Layers {
		imgW := float64(layer.Image.Bounds().Dx())
		imgH := float64(layer.Image.Bounds().Dy())

		offsetX := -camX * layer.ScrollX
		offsetY := -(camY - bg.BaseY) * layer.ScrollY

		startX := math.Mod(offsetX, imgW)
		if startX > 0 {
			startX -= imgW
		}

		scaleY := 1.0
		if layer.StretchY && imgH < screenH {
			scaleY = screenH / imgH
		}

		opBase := &ebiten.DrawImageOptions{}
		opBase.GeoM.Scale(1, scaleY)
		opBase.GeoM.Translate(0, offsetY)

		for x := startX; x < screenW; x += imgW {
			op := *opBase
			op.GeoM.Translate(x, 0)
			screen.DrawImage(layer.Image, &op)
		}
	}
}
