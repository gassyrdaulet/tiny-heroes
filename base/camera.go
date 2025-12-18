package base

import (
	"github.com/gassyrdaulet/go-fighting-game/constants"
)

type Camera struct {
	X, Y          float64
	Width, Height int
}

func (c *Camera) WorldToScreen(worldX, worldY float64) (screenX, screenY float64) {
	screenX = float64(worldX - c.X + float64(c.Width)/2)
	screenY = float64(worldY - c.Y + float64(c.Height)/2)
	return
}

func (c *Camera) UpdateFromPlayers(players []PlayerPosition) {
	if len(players) > 0 {
		minX, minY := players[0].Position()
		maxX, maxY := minX, minY

		for _, p := range players {
			pX, pY := p.Position()
			if pX < minX { minX = pX }
			if pY < minY { minY = pY }
			if pX > maxX { maxX = pX }
			if pY > maxY { maxY = pY }
		}
		
		if maxX - minX  < constants.ScreenW - constants.CameraDeadZoneX {
			c.X = (minX + maxX) / 2
		}
		if maxY - minY < constants.ScreenH - constants.CameraDeadZoneY {
			c.Y = (minY + maxY) / 2 - constants.CameraYOffset
		}
	}
}