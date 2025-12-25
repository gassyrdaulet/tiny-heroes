package actor

import (
	"github.com/gassyrdaulet/go-fighting-game/base"
	"github.com/gassyrdaulet/go-fighting-game/base/physics"
	"github.com/gassyrdaulet/go-fighting-game/characters"
	"github.com/hajimehoshi/ebiten/v2"
)

type Actor struct {
	*physics.Body
	*base.StateMachine 
	*base.Animator
	Character         *characters.Character
	Controller        base.Controller
	MaxHp             int
	Hp                int
	Speed             float64
	JumpForce         float64
	JumpForceCurrent  float64
	Direction         int
	ChargingJumpTicksMax int
	ChargingJumpTicks int
}

func (a *Actor) GoLeft() {
	if a.OnGround {
		a.VX = -a.Speed
	} else {
		a.VX = -a.Speed * 5 / 6
	}
	a.Direction = -1
}

func (a *Actor) GoRight() {
	if a.OnGround {
		a.VX = a.Speed
	} else {
		a.VX = a.Speed * 5 / 6
	}
	a.Direction = 1
}

func (a *Actor) ChargeJump() {
	if a.OnGround && !a.Is(ChargeJump) {
		a.ChangeState(ChargeJump)
		a.ChargingJumpTicks = a.ChargingJumpTicksMax
		if a.VX != 0 {
			a.JumpForceCurrent = a.JumpForce * 1.1
		} else {
			a.JumpForceCurrent = a.JumpForce
		}
		a.VX = 0
	}
}

func (a *Actor) Update(world *physics.World) {
	input := a.Controller.GetInput()
	if a.IsNot(ChargeJump) {
		if input.Left {
			a.GoLeft()
		} else if input.Right {
			a.GoRight()
		} else {
			a.VX = 0
		}
		if input.Up {
			a.ChargeJump()
		}
	}

	if a.Is(ChargeJump) && a.ChargingJumpTicks <= 0 && a.OnGround {
		a.VY = a.JumpForceCurrent
		a.ChangeState(Jump)
	}
	if a.Is(ChargeJump) && a.ChargingJumpTicks > 0 {
		a.ChargingJumpTicks--
	}

	world.Step(a)

	a.UpdateState()
	a.UpdateFrame(string(a.CurrentState))
}

func (a *Actor) UpdateState() {
	if a.Is(ChargeJump) {
		return
	}

	if a.OnGround {
		if a.VX == 0 {
			a.ChangeState(Idle)
		} else {
			a.ChangeState(Run)
		}
	} else {
		if a.VY < 0 {
			a.ChangeState(Jump)
		} else {
			a.ChangeState(Fall)
		}
	}
}
 
func (a *Actor) Draw(screen *ebiten.Image, sx, sy float64) {
	a.DrawFrame(screen, sx, sy, a.Direction == -1)
}