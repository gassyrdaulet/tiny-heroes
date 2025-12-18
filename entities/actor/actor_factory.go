package actor

import (
	b "github.com/gassyrdaulet/go-fighting-game/base"
	p "github.com/gassyrdaulet/go-fighting-game/base/physics"
	c "github.com/gassyrdaulet/go-fighting-game/characters"
)

func NewActor(x, y float64, characterID string, controller b.Controller, direction int) *Actor {
	char := c.Characters[characterID]

	return &Actor{
		Body: &p.Body{
			X: x,
			Y: y,
			Width: char.Width,
			Height: char.Height,
			Weight: char.Weight,
		},
		StateMachine: b.NewStateMachine(Idle),
		Animator: b.NewAnimator(char.Animations, string(Idle)),
		Character:   char,
		Controller: controller,
		MaxHp:          char.MaxHP,
		Hp:          char.MaxHP,
		Speed:          char.Speed,
		JumpForce:          char.JumpForce,
		Direction:   direction,
		ChargingJumpTicksMax: char.ChargingJumpTicks,
	}
}
