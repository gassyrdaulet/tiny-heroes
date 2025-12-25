package actor

import (
	b "github.com/gassyrdaulet/go-fighting-game/base"
	p "github.com/gassyrdaulet/go-fighting-game/base/physics"
	"github.com/gassyrdaulet/go-fighting-game/characters"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewActor(x, y float64, controller b.Controller, direction int, char *characters.Character) *Actor {
	animCopy := make(map[string]*b.Animation)
    for k, v := range char.Animations {
        framesCopy := make([]*ebiten.Image, len(v.Frames))
        copy(framesCopy, v.Frames)
        animCopy[k] = &b.Animation{
            Frames:     framesCopy,
            FrameSpeed: v.FrameSpeed,
            Loop:       v.Loop,
            XO:         v.XO,
            YO:         v.YO,
        }
    }

	return &Actor{
		Body: &p.Body{
			X: x,
			Y: y,
			Width: char.Width,
			Height: char.Height,
			Weight: char.Weight,
		},
		StateMachine: b.NewStateMachine(Idle),
		Animator: b.NewAnimator(animCopy, string(Idle)),
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
