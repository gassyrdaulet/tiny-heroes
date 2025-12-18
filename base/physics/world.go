package physics

import (
	"math"

	"github.com/gassyrdaulet/go-fighting-game/base"
	"github.com/gassyrdaulet/go-fighting-game/constants"
)

type World struct {
	Tiles               *base.TileMap
	VirtualBorderLeftX  float64
	VirtualBorderRightX float64
	Width, Height       float64
}

func (w *World) Step(p PhysicalBody) {
	x, y := p.Position()
	vx, vy := p.Velocity()
	wid, h := p.Size()
	weight := p.GetWeight()

	vy += constants.Gravity * weight

	newX := x + vx
	newY := y + vy

	if vx != 0 {
		if newX-wid/2 < w.VirtualBorderLeftX || newX+wid/2 > w.VirtualBorderRightX {
			newX = x
		} else {
			tileX := int((newX + math.Copysign(wid/2, vx)) / constants.TileSize)
			tileY1 := int(y / constants.TileSize)
			tileY2 := int((y + h - 1) / constants.TileSize)

			for ty := tileY1; ty <= tileY2; ty++ {
				if w.Tiles.IsSolid(tileX, ty) {
					newX = x
					break
				}
			}
		}
	}

	onGround := false

	if vy != 0 {
		if vy > 0 {
			tileY := int((newY + h) / constants.TileSize)
			tileX1 := int((newX - wid/2) / constants.TileSize)
			tileX2 := int((newX + wid/2) / constants.TileSize)

			for tx := tileX1; tx <= tileX2; tx++ {
				if w.Tiles.IsSolid(tx, tileY) {
					newY = float64(tileY)*constants.TileSize - h
					vy = 0
					onGround = true
					break
				}
			}
		} else {
			tileY := int(newY / constants.TileSize)
			tileX1 := int((newX - wid/2) / constants.TileSize)
			tileX2 := int((newX + wid/2) / constants.TileSize)

			for tx := tileX1; tx <= tileX2; tx++ {
				if w.Tiles.IsSolid(tx, tileY) {
					newY = float64(tileY+1) * constants.TileSize
					vy = 0
					break
				}
			}
		}
	}

	p.SetOnGround(onGround)
	p.SetPosition(newX, newY)
	p.SetVelocity(vx, vy)
}

func (w *World) UpdateVirtualBounds(cam *base.Camera) {
	w.VirtualBorderLeftX = cam.X - float64(cam.Width)/2
	w.VirtualBorderRightX = cam.X + float64(cam.Width)/2
}

func NewWorld(tiles *base.TileMap) *World {
	w := &World{
		Tiles: tiles,
	}
	w.Width = float64(tiles.Width) * constants.TileSize
	w.Height = float64(tiles.Height) * constants.TileSize
	return w
}
