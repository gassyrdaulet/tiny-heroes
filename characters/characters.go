package characters

import (
	"fmt"

	"github.com/gassyrdaulet/go-fighting-game/base"
)

var Characters = map[string]*Character{
	"blue": {
		ID:                   "blue",
		Name:                 "Blue",
		Speed:                4.0,
		JumpForce:                 -11.0,
		Damage:               10,
		MaxHP:                100,
		Weight:        1,
		Width:                24,
		Height:               30,
		ChargingJumpTicks: 20,
		AnimationsConfigs: []base.AnimationConfig{
			base.Anim("idle", "assets/sprites/3/Idle.png", 42, 30, 4, 0, 12, 6, true),
			base.Anim("run", "assets/sprites/3/Run.png", 42, 30, 6, 0, 12, 6, true),
			base.Anim("charging_jump", "assets/sprites/3/Jump.png", 42, 30, 3, 0, 12, 6, false),
			base.Anim("jump", "assets/sprites/3/Jump.png", 42, 30, 1, 126, 12, 6, false),
			base.Anim("fall", "assets/sprites/3/Jump.png", 42, 30, 1, 252, 12, 6, false),
		},
	},
	"pink": {
		ID:                   "pink",
		Name:                 "Pink",
		Speed:                5.0,
		Damage:               11,
		JumpForce:                 -10.0,
		MaxHP:                75,
		Weight:        0.9,
		Width:                24,
		Height:               30,
		ChargingJumpTicks: 15,
		AnimationsConfigs: []base.AnimationConfig{
			base.Anim("idle", "assets/sprites/1/Idle.png", 42, 30, 4, 0, 12, 6, true),
			base.Anim("run", "assets/sprites/1/Run.png", 42, 30, 6, 0, 12, 6, true),
			base.Anim("charging_jump", "assets/sprites/1/Jump.png", 42, 30, 3, 0, 12, 6, false),
			base.Anim("jump", "assets/sprites/1/Jump.png", 42, 30, 1, 126, 12, 6, false),
			base.Anim("fall", "assets/sprites/1/Jump.png", 42, 30, 1, 252, 12, 6, false),
		},
	},
	"white": {
		ID:                   "white",
		Name:                 "White",
		Speed:                3.4,
		Damage:               15,
		JumpForce:                 -12.0,
		MaxHP:                140,
		Weight:        1.1,
		Width:                24,
		Height:               30,
		ChargingJumpTicks: 25,
		AnimationsConfigs: []base.AnimationConfig{
			base.Anim("idle", "assets/sprites/2/Idle.png", 42, 30, 4, 0, 12, 6, true),
			base.Anim("run", "assets/sprites/2/Run.png", 42, 30, 6, 0, 12, 6, true),
			base.Anim("charging_jump", "assets/sprites/2/Jump.png", 42, 30, 3, 0, 12, 6, false),
			base.Anim("jump", "assets/sprites/2/Jump.png", 42, 30, 1, 126, 12, 6, false),
			base.Anim("fall", "assets/sprites/2/Jump.png", 42, 30, 1, 252, 12, 6, false),
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
