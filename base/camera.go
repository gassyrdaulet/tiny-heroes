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

	anchorX, anchorY := players[0].Position()

	minX, minY := anchorX, anchorY
	maxX, maxY := anchorX, anchorY

	for i := 1; i < len(players); i++ {
		x, y := players[i].Position()

		if x < minX {
			minX = x
		}
		if y < minY {
			minY = y
		}
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	groupCenterX := (minX + maxX) / 2
	groupCenterY := (minY + maxY) / 2

	targetX := u.Lerp(groupCenterX, anchorX, constants.CameraAnchorWeight)
	targetY := u.Lerp(groupCenterY, anchorY, constants.CameraAnchorWeight)

	halfW := float64(constants.ScreenW) / 2
	halfH := float64(constants.ScreenH) / 2

	targetX = u.Clamp(targetX, halfW, worldW-halfW)
	targetY = u.Clamp(targetY, halfH, worldH-halfH)

	const smooth = constants.CameraSmoothness
	c.X = u.Lerp(c.X, targetX, smooth)
	c.Y = u.Lerp(c.Y, targetY, smooth)
}
