package characters

import (
	"fmt"

	"github.com/gassyrdaulet/go-fighting-game/base"
)

var Characters = map[string]*Character{
	"blue": {
		ID:                "blue",
		Name:              "Blue",
		Speed:             3.8,
		JumpForce:         -9.2,
		Damage:            10,
		MaxHP:             100,
		Weight:            1,
		Width:             12,
		Height:            26,
		ChargingJumpTicks: 23,
		AnimationsConfigs: []base.AnimationConfig{
			base.Anim("idle", "assets/sprites/3/Idle.png", 42, 42, 4, 0, 0, 6, 5, 16, true),
			base.Anim("run", "assets/sprites/3/Run.png", 42, 42, 6, 0, 0, 6, 5, 16, true),
			base.Anim("charging_jump", "assets/sprites/3/Jump.png", 42, 42, 3, 0, 0, 6, 5, 16, false),
			base.Anim("jump", "assets/sprites/3/Jump.png", 42, 42, 1, 126, 0, 6, 5, 16, false),
			base.Anim("fall", "assets/sprites/3/Jump.png", 42, 42, 1, 252, 0, 6, 5, 16, false),
		},
	},
	"pink": {
		ID:                "pink",
		Name:              "Pink",
		Speed:             4.2,
		Damage:            11,
		JumpForce:         -9.5,
		MaxHP:             75,
		Weight:            0.9,
		Width:             12,
		Height:            26,
		ChargingJumpTicks: 26,
		AnimationsConfigs: []base.AnimationConfig{
			base.Anim("idle", "assets/sprites/1/Idle.png", 42, 42, 4, 0, 0, 6, 5, 16, true),
			base.Anim("run", "assets/sprites/1/Run.png", 42, 42, 6, 0, 0, 6, 5, 16, true),
			base.Anim("charging_jump", "assets/sprites/1/Jump.png", 42, 42, 3, 0, 0, 6, 5, 16, false),
			base.Anim("jump", "assets/sprites/1/Jump.png", 42, 42, 1, 126, 0, 6, 5, 16, false),
			base.Anim("fall", "assets/sprites/1/Jump.png", 42, 42, 1, 252, 0, 6, 5, 16, false),
		},
	},
	"white": {
		ID:                "white",
		Name:              "White",
		Speed:             3.6,
		Damage:            15,
		JumpForce:         -9.0,
		MaxHP:             140,
		Weight:            1.1,
		Width:             12,
		Height:            26,
		ChargingJumpTicks: 20,
		AnimationsConfigs: []base.AnimationConfig{
			base.Anim("idle", "assets/sprites/2/Idle.png", 42, 42, 4, 0, 0, 6, 5, 16, true),
			base.Anim("run", "assets/sprites/2/Run.png", 42, 42, 6, 0, 0, 6, 5, 16, true),
			base.Anim("charging_jump", "assets/sprites/2/Jump.png", 42, 42, 3, 0, 0, 6, 5, 16, false),
			base.Anim("jump", "assets/sprites/2/Jump.png", 42, 42, 1, 126, 0, 6, 5, 16, false),
			base.Anim("fall", "assets/sprites/2/Jump.png", 42, 42, 1, 252, 0, 6, 5, 16, false),
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
