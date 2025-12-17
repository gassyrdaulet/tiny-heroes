package base

import (
	"image"

	u "github.com/gassyrdaulet/go-fighting-game/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	Frames     []*ebiten.Image
	FrameIndex int
	FrameTick  int
	FrameSpeed int
	Loop       bool
}

type AnimationConfig struct {
	Name       string
	FilePath   string
	Width      int
	Height     int
	Count      int
	StartX     int
	StartY     int
	FrameSpeed int
	Loop       bool
}

func Anim(
	name, path string,
	w, h, count, x, y, speed int,
	loop bool,
) AnimationConfig {
	return AnimationConfig{
		Name:       name,
		FilePath:   path,
		Width:      w,
		Height:     h,
		Count:      count,
		StartX:     x,
		StartY:     y,
		FrameSpeed: speed,
		Loop:       loop,
	}
}

func loadFrames(cfg *AnimationConfig) ([]*ebiten.Image, error) {
	img, err := u.LoadImage(cfg.FilePath)
	if err != nil {
		return nil, err
	}

	frames := make([]*ebiten.Image, cfg.Count)
	for i := 0; i < cfg.Count; i++ {
		sx0 := i*cfg.Width + cfg.StartX
		sy0 := cfg.StartY
		sx1 := sx0 + cfg.Height
		sy1 := cfg.Height + cfg.Height

		sub := img.SubImage(image.Rect(sx0, sy0, sx1, sy1)).(*ebiten.Image)
		frames[i] = sub
	}
	return frames, nil
}

func LoadAnimations(cfgs []AnimationConfig) (map[string]*Animation, error) {
	animations := map[string]*Animation{}

	for _, cfg := range cfgs {
		frames, err := loadFrames(&cfg)
		if err != nil {
			return nil, err
		}

		animations[cfg.Name] = &Animation{
			Frames:     frames,
			FrameSpeed: cfg.FrameSpeed,
			Loop:       cfg.Loop,
		}
	}

	return animations, nil
}
