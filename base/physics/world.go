package physics

import (
	"github.com/gassyrdaulet/go-fighting-game/base"
	"github.com/gassyrdaulet/go-fighting-game/constants"
	"github.com/gassyrdaulet/go-fighting-game/entities"
)

type World struct {}

func (w *World) Step(p *entities.Controllable, in base.Input) {
	body := p.Body
	char := p.Character

	body.VX = in.MoveX * char.Speed

	if in.Jump && body.OnGround {
		body.VY = char.Jump
		body.OnGround = false
	}

	body.VY += constants.Gravity - char.GravityResist

	body.X += body.VX
	body.Y += body.VY

	if body.Y+body.Height >= constants.ScreenH-constants.GroundHeight {
		body.Y = constants.ScreenH - constants.GroundHeight - body.Height
		body.VY = 0
		body.OnGround = true
	}
}