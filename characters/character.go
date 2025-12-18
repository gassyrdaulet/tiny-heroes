package characters

import (
	"github.com/gassyrdaulet/go-fighting-game/base"
)

type Character struct {
	ID                   string
	Name                 string
	Speed                float64
	JumpForce                 float64
	Weight        float64
	Width, Height        float64
	Damage               int
	MaxHP                int
	ChargingJumpTicks int
	Animations           map[string]*base.Animation
	AnimationsConfigs    []base.AnimationConfig
}