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
	XO, YO     float64
	Loop       bool
}

type AnimationConfig struct {
	Name        string
	FilePath    string
	Width       int
	Height      int
	Count       int
	StartX      int
	StartY      int
	LeftOffset  int
	RightOffset int
	FrameSpeed  int
	XO, YO      float64
	Loop        bool
}

func Anim(
	name, path string,
	w, h, count, x, y, speed int,
	xOffset, yOffset float64,
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
		XO:         xOffset,
		YO:         yOffset,
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
		sx1 := sx0 + cfg.Width
		sy1 := sy0 + cfg.Height

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
			XO:         cfg.XO,
			YO:         cfg.YO,
		}
	}

	return animations, nil
}

type Animator struct {
	CurrentAnimation           string
	Animations                 map[string]*Animation
	SpriteScaleX, SpriteScaleY float64
}

func NewAnimator(animations map[string]*Animation, current string) *Animator {
	return &Animator{
		Animations:       animations,
		CurrentAnimation: current,
		SpriteScaleX:     1,
		SpriteScaleY:     1,
	}
}

func (a *Animator) UpdateFrame(state string) {
	if a.CurrentAnimation != state {
		a.CurrentAnimation = state
		anim := a.Animations[a.CurrentAnimation]
		anim.FrameIndex = 0
		anim.FrameTick = 0
	}

	anim := a.Animations[a.CurrentAnimation]
	anim.FrameTick++
	if anim.FrameTick >= anim.FrameSpeed {
		anim.FrameTick = 0
		if anim.Loop {
			anim.FrameIndex = (anim.FrameIndex + 1) % len(anim.Frames)
		} else if anim.FrameIndex < len(anim.Frames)-1 {
			anim.FrameIndex++
		}
	}
}

func (a *Animator) DrawFrame(screen *ebiten.Image, x, y float64, flip bool) {
	anim := a.Animations[a.CurrentAnimation]
	if len(anim.Frames) == 0 {
		return
	}
	currentFrame := anim.Frames[anim.FrameIndex]
	frameWidth := float64(currentFrame.Bounds().Dx())

	op := &ebiten.DrawImageOptions{}
	if flip {
		op.GeoM.Scale(-a.SpriteScaleX, a.SpriteScaleY)
		op.GeoM.Translate(x+(frameWidth/2)*a.SpriteScaleX-anim.XO, y-anim.YO)
	} else {
		op.GeoM.Scale(a.SpriteScaleX, a.SpriteScaleY)
		op.GeoM.Translate(x-(frameWidth/2)*a.SpriteScaleX+anim.XO, y-anim.YO)
	}
	screen.DrawImage(currentFrame, op)
}
