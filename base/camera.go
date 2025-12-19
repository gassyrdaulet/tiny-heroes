package base

import (
	"github.com/gassyrdaulet/go-fighting-game/constants"
	u "github.com/gassyrdaulet/go-fighting-game/utils"
)

type Camera struct {
	X, Y          float64
	Width, Height int
}

func (c *Camera) TopLeft() (x, y float64) {
	return c.X - float64(c.Width)/2,
		c.Y - float64(c.Height)/2
}

func (c *Camera) WorldToScreen(worldX, worldY float64) (screenX, screenY float64) {
	tlx, tly := c.TopLeft()
	return worldX - tlx, worldY - tly
}

func (c *Camera) UpdateFromPlayers(players []PlayerPosition, worldW, worldH float64) {
	if len(players) == 0 {
		return
	}

	minX, minY := players[0].Position()
	maxX, maxY := minX, minY

	for _, p := range players {
		pX, pY := p.Position()

		if pX < minX {
			minX = pX
		}
		if pY < minY {
			minY = pY
		}
		if pX > maxX {
			maxX = pX
		}
		if pY > maxY {
			maxY = pY
		}
	}

	targetX, targetY := c.X, c.Y

	if maxX-minX < constants.ScreenW-constants.CameraMaxOffsetX {
		targetX = (minX + maxX) / 2
	}
	if maxY-minY < constants.ScreenH-constants.CameraMaxOffsetY {
		targetY = (minY + maxY) / 2
	}

	halfW := float64(constants.ScreenW) / 2
	halfH := float64(constants.ScreenH) / 2

	targetX = u.Clamp(targetX, halfW, worldW-halfW)
	targetY = u.Clamp(targetY, halfH, worldH-halfH)

	const smooth = 0.45
	c.X = u.Lerp(c.X, targetX, smooth)
	c.Y = u.Lerp(c.Y, targetY, smooth)
}
