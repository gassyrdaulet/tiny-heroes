package characters

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gassyrdaulet/go-fighting-game/base"
)

type Character struct {
	ID                   string
	Name                 string
	Speed                float64
	JumpForce            float64
	Weight        		 float64
	Width, Height        float64
	Damage               int
	MaxHP                int
    AttackTicks          int
    AttackCooldownTicks  int
	ChargingJumpTicks 	 int
	DyingTicks       	 int
	HurtingTicks       	 int
	AttackRange       	 float64
	Animations           map[string]*base.Animation
	AnimationsConfigs    []base.AnimationConfig
}

type CharactersJSON struct {
    Characters []CharacterJSON `json:"characters"`
}

type CharacterJSON struct {
    ID                  string           `json:"id"`
    Name                string           `json:"name"`
    Speed               float64          `json:"speed"`
    JumpForce           float64          `json:"jumpForce"`
    Damage              int              `json:"damage"`
    MaxHP               int              `json:"maxHP"`
    Weight              float64          `json:"weight"`
    Width               float64          `json:"width"`
    Height              float64          `json:"height"`
    AttackTicks         int              `json:"attackTicks"`
    AttackCooldownTicks int              `json:"attackCooldownTicks"`
    ChargingJumpTicks   int              `json:"chargingJumpTicks"`
	DyingTicks       	int              `json:"dyingTicks"`
    HurtingTicks        int              `json:"hurtingTicks"`
	AttackRange       	float64          `json:"attackRange"`
    Animations          []AnimationJSON  `json:"animations"`
}

type AnimationJSON struct {
    Name        string 	   `json:"name"`
    Group       string    `json:"group"`
    Image       string 	   `json:"image"`
    FrameWidth  int    	   `json:"frameWidth"`
    FrameHeight int    	   `json:"frameHeight"`
    Frames      int    	   `json:"frames"`
    StartX      int    	   `json:"startX"`
    StartY      int    	   `json:"startY"`
    Speed     	int    	   `json:"speed"`
    OffsetX     float64    `json:"offsetX"`
    OffsetY     float64    `json:"offsetY"`
    Loop        bool       `json:"loop"`
}

func LoadCharacters(path string) (map[string]*Character, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    var jsonData CharactersJSON
    if err := json.Unmarshal(data, &jsonData); err != nil {
        return nil, err
    }

    result := make(map[string]*Character)

    for _, c := range jsonData.Characters {
        char := &Character{
            ID:                  c.ID,
            Name:                c.Name,
            Speed:               c.Speed,
            JumpForce:           c.JumpForce,
            Damage:              c.Damage,
            MaxHP:               c.MaxHP,
            Weight:              c.Weight,
            Width:               c.Width,
            Height:              c.Height,
            AttackTicks:         c.AttackTicks,
            AttackCooldownTicks: c.AttackCooldownTicks,
            ChargingJumpTicks:   c.ChargingJumpTicks,
            DyingTicks:          c.DyingTicks,
            HurtingTicks:        c.HurtingTicks,
            AttackRange:         c.AttackRange,
            Animations:          make(map[string]*base.Animation),
            AnimationsConfigs:   []base.AnimationConfig{},
        }

        for _, a := range c.Animations {
            cfg := base.Anim(
                a.Name,
                a.Image,
                a.Group,
                a.FrameWidth,
                a.FrameHeight,
                a.Frames,
                a.StartX,
                a.StartY,
                a.Speed,
                a.OffsetX,
                a.OffsetY,
                a.Loop,
            )

            char.AnimationsConfigs = append(char.AnimationsConfigs, cfg)
        }

		animations, err := base.LoadAnimations(char.AnimationsConfigs)
		if err != nil {
			return nil, fmt.Errorf("character %s: %w", char.ID, err)
		}
		char.Animations = animations

        result[char.ID] = char
    }

    return result, nil
}