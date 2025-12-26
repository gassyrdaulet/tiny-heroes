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
	paused 				bool
}

func (w *World) Step(p PhysicalBody) {
	x, y := p.Position()
	vx, vy := p.Velocity()
	wid, h := p.Size()
	weight := p.GetWeight()

	vy += constants.Gravity * weight

	newX := x + vx
	newY := y + vy

	if newX+wid/2 < 0 || newX-wid/2 > w.Width || newY - h < 0 || newY + h > w.Height {
		p.Die()
	}

	if vx != 0 {
		if ((newX-wid/2 < w.VirtualBorderLeftX || newX+wid/2 > w.VirtualBorderRightX) && constants.VirtualBorders) ||
			(newX-wid/2 < 0 || newX+wid/2 > w.Width) {
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

	var onGround bool
	newY, onGround = resolveVerticalCollision(
		w, y, newX, newY, wid, h, &vy,
	)

	p.SetOnGround(onGround)
	p.SetPosition(newX, newY)
	p.SetVelocity(vx, vy)
}

func resolveVerticalCollision(w *World, oldY, newX, newY, wWidth, wHeight float64, vy *float64) (float64, bool) {
	onGround := false

	if *vy == 0 {
		return newY, onGround
	}

	stepSign := 1.0
	if *vy < 0 {
		stepSign = -1.0
	}

	tileX1 := int((newX - wWidth/2) / constants.TileSize)
	tileX2 := int((newX + wWidth/2) / constants.TileSize)

	if stepSign > 0 {
		tileYStart := int((oldY + wHeight) / constants.TileSize)
		tileYEnd   := int((newY + wHeight) / constants.TileSize)
		for ty := tileYStart; ty <= tileYEnd; ty++ {
			for tx := tileX1; tx <= tileX2; tx++ {
				if w.Tiles.IsSolid(tx, ty) {
					newY = float64(ty)*constants.TileSize - wHeight
					*vy = 0
					onGround = true
					return newY, onGround
				}

				if w.Tiles.IsPlatform(tx, ty) {
					tileTop := float64(ty * constants.TileSize)
					if oldY+wHeight <= tileTop {
						newY = tileTop - wHeight
						*vy = 0
						onGround = true
						return newY, onGround
					}
				}
			}
		}
	} else {
		tileYStart := int(newY / constants.TileSize)
		tileYEnd := int(oldY / constants.TileSize)
		for ty := tileYEnd; ty >= tileYStart; ty-- {
			for tx := tileX1; tx <= tileX2; tx++ {
				if w.Tiles.IsSolid(tx, ty){
					newY = float64(ty+1) * constants.TileSize
					*vy = 0
					return newY, onGround
				}
			}
		}
	}

	return newY, onGround
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

func (w *World) Clear() {
	w.Tiles = nil
	w.VirtualBorderLeftX = 0
	w.VirtualBorderRightX = 0
	w.Width = 0
	w.Height = 0
}