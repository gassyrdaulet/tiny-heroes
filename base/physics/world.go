package physics

import (
	"github.com/gassyrdaulet/go-fighting-game/base"
	"github.com/gassyrdaulet/go-fighting-game/constants"
	"github.com/gassyrdaulet/go-fighting-game/entities"
)

type World struct {
	Tiles base.TileMap
}

func (w *World) StepControllable(p *entities.Controllable, in base.Input) {
	body := p.Body
	char := p.Character

	// ===== INPUT =====
	if p.IsJumpCharging && p.ChargingJumpTicks > 0 {
		p.ChargingJumpTicks--
	} else {
		if p.OnGround {
			body.VX = in.MoveX * char.Speed
		} else {
			body.VX = in.MoveX*char.Speed - char.Speed/3
		}
	}

	if in.Jump && !p.IsJumpCharging && body.OnGround {
		p.ChargingJumpTicks = p.Character.ChargingJumpTicksMax
		p.IsJumpCharging = true
	}

	if p.IsJumpCharging && body.OnGround && p.ChargingJumpTicks <= 0 {
		p.vy = char.Jump
		p.InAir = 1
		p.ChargingJumpTicks = 0
		p.IsJumpCharging = false
	}

	// ===== GRAVITY =====
	body.VY += constants.Gravity - char.GravityResist

	// ===== MOVE X =====
	newX := body.X + body.VX
	if body.VX != 0 {
		if !w.collidesAt(newX, body.Y, body) {
			body.X = newX
		} else {
			body.VX = 0
		}
	}

	// ===== MOVE Y =====
	newY := body.Y + body.VY
	if body.VY != 0 {
		if !w.collidesAt(body.X, newY, body) {
			body.Y = newY
			body.OnGround = false
		} else {
			if body.VY > 0 {
				body.OnGround = true
			}
			body.VY = 0
		}
	}
}

func (w *World) collidesAt(x, y float64, body *entities.Body) bool {
	left := worldToTile(x)
	right := worldToTile(x + body.Width - 1)
	top := worldToTile(y)
	bottom := worldToTile(y + body.Height - 1)

	for ty := top; ty <= bottom; ty++ {
		for tx := left; tx <= right; tx++ {
			if w.Tiles.IsSolid(tx, ty) {
				return true
			}
		}
	}

	return false
}

func worldToTile(v float64) int {
	return int(v) / constants.TileSize
}
