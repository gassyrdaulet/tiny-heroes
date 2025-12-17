package characters

import (
	"fmt"

	"github.com/gassyrdaulet/go-fighting-game/base"
)

type Character struct {
	ID                   string
	Name                 string
	Speed                float64
	SpeedInAir           float64
	Jump                 float64
	GravityResist        float64
	InitialWidth, InitialHeight        float64
	Damage               int
	MaxHP                int
	ChargingJumpTicksMax int // need uneven number
	Animations           map[string]*base.Animation
	AnimationsConfigs    []base.AnimationConfig
	ScaleX, ScaleY float64
}

var Characters = map[string]*Character{
	"blue": {
		ID:                   "blue",
		Name:                 "Blue",
		Speed:                4.0,
		Jump:                 -11.0,
		Damage:               10,
		MaxHP:                100,
		GravityResist:        0,
		InitialWidth:                60,
		InitialHeight:               75,
		ChargingJumpTicksMax: 20,
		ScaleX: 2.5,
		ScaleY: 2.5,
		AnimationsConfigs: []base.AnimationConfig{
			base.Anim("idle", "assets/sprites/3/Idle.png", 42, 30, 4, 0, 12, 6, true),
			base.Anim("run", "assets/sprites/3/Run.png", 42, 30, 6, 0, 12, 6, true),
			base.Anim("charging_jump", "assets/sprites/3/Jump.png", 42, 30, 3, 0, 12, 6, false),
			base.Anim("jump_up", "assets/sprites/3/Jump.png", 42, 30, 1, 126, 12, 6, false),
			base.Anim("jump_down", "assets/sprites/3/Jump.png", 42, 30, 1, 252, 12, 6, false),
		},
	},
	"pink": {
		ID:                   "pink",
		Name:                 "Pink",
		Speed:                5.0,
		Damage:               11,
		Jump:                 -10.0,
		MaxHP:                75,
		GravityResist:        0.1,
		InitialWidth:                60,
		InitialHeight:               75,
		ChargingJumpTicksMax: 15,
		ScaleX: 2.5,
		ScaleY: 2.5,
		AnimationsConfigs: []base.AnimationConfig{
			base.Anim("idle", "assets/sprites/1/Idle.png", 42, 30, 4, 0, 12, 6, true),
			base.Anim("run", "assets/sprites/1/Run.png", 42, 30, 6, 0, 12, 6, true),
			base.Anim("charging_jump", "assets/sprites/1/Jump.png", 42, 30, 3, 0, 12, 6, false),
			base.Anim("jump_up", "assets/sprites/1/Jump.png", 42, 30, 1, 126, 12, 6, false),
			base.Anim("jump_down", "assets/sprites/1/Jump.png", 42, 30, 1, 252, 12, 6, false),
		},
	},
	"white": {
		ID:                   "white",
		Name:                 "White",
		Speed:                3.4,
		Damage:               15,
		Jump:                 -12.0,
		MaxHP:                140,
		GravityResist:        -0.2,
		InitialWidth:                60,
		InitialHeight:               75,
		ChargingJumpTicksMax: 25,
		ScaleX: 2.5,
		ScaleY: 2.5,
		AnimationsConfigs: []base.AnimationConfig{
			base.Anim("idle", "assets/sprites/2/Idle.png", 42, 30, 4, 0, 12, 6, true),
			base.Anim("run", "assets/sprites/2/Run.png", 42, 30, 6, 0, 12, 6, true),
			base.Anim("charging_jump", "assets/sprites/2/Jump.png", 42, 30, 3, 0, 12, 6, false),
			base.Anim("jump_up", "assets/sprites/2/Jump.png", 42, 30, 1, 126, 12, 6, false),
			base.Anim("jump_down", "assets/sprites/2/Jump.png", 42, 30, 1, 252, 12, 6, false),
		},
	},
}

func LoadCharacterAnimations() error {
	for id, char := range Characters {
		animations, err := base.LoadAnimations(char.AnimationsConfigs)
		if err != nil {
			return fmt.Errorf("character %s: %w", id, err)
		}
		char.Animations = animations
	}
	return nil
}
