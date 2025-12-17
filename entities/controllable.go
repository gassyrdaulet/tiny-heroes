package entities

import (
	"github.com/gassyrdaulet/go-fighting-game/base"
	"github.com/gassyrdaulet/go-fighting-game/characters"
	"github.com/gassyrdaulet/go-fighting-game/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

type Controllable struct {
	*base.Entity
	*base.Body
	Hp                int
	Direction         int
	Controller         base.Controller
	Character         *characters.Character
	CurrentAnim       string
	PrevAnim          string
	IsJumpCharging bool
	ChargingJumpTicks int
}

func (p *Controllable) Update([]base.GameEntity) error {
	char := p.Character
	p.vx = 0

	newSpeed := char.Speed - float64(p.InAir)*(char.Speed/3)

	if  p.IsJumpCharging && p.ChargingJumpTicks > 0{
		p.ChargingJumpTicks -= 1
	} else {
		if ebiten.IsKeyPressed(p.InputKeys.LeftKey) {
			p.vx = -newSpeed
			p.Direction = -1
		}
		if ebiten.IsKeyPressed(p.InputKeys.RightKey) {
			p.vx = newSpeed
			p.Direction = 1
		}
	}
	if !p.IsJumpCharging && p.InAir == 0 && ebiten.IsKeyPressed(p.InputKeys.UpKey) {
		p.IsJumpCharging = true
		p.ChargingJumpTicks = p.Character.ChargingJumpTicksMax
	}

	if p.IsJumpCharging && p.InAir == 0 && p.ChargingJumpTicks <= 0 {
		p.vy = char.Jump
		p.InAir = 1
		p.ChargingJumpTicks = 0
		p.IsJumpCharging = false
	}

	p.vy += constants.Gravity - char.GravityResist

	p.X = p.X + p.vx
	p.Y = p.Y + p.vy

	if p.Y+p.Height >= constants.ScreenH-constants.GroundHeight {
		p.Y = constants.ScreenH - constants.GroundHeight - p.Height
		p.vy = 0
		p.InAir = 0
	}

	if p.X-p.Width/2 < 0{
		p.X = 0+p.Width/2
	}
	if p.X+p.Width/2 > constants.ScreenW {
		p.X = constants.ScreenW-p.Width/2
	}

	if p.InAir == 0 {
		if p.ChargingJumpTicks > 0 {
			p.ChangeAnim("charging_jump")
		} else {
			if p.vx != 0 {
				p.ChangeAnim("run")
			} else {
				p.ChangeAnim("idle")
			}
		}
	} else {
		if p.vy < 0 {
			p.ChangeAnim("jump_up")
		} else {
			p.ChangeAnim("jump_down")
		}
	}

	anim := char.Animations[p.CurrentAnim]

	if len(anim.Frames) > 0 {
		anim.FrameTick++
		if anim.FrameTick >= anim.FrameSpeed {
			anim.FrameTick = 0

			if anim.Loop {
				anim.FrameIndex = (anim.FrameIndex + 1) % len(anim.Frames)
			} else {
				if anim.FrameIndex < len(anim.Frames)-1 {
					anim.FrameIndex++
				}
			}
		}
	}

	return nil
}

func (p *Controllable) Draw(screen *ebiten.Image) {
	anim := p.Character.Animations[p.CurrentAnim]
	if len(anim.Frames) == 0 {
		return
	}
	currentFrame := anim.Frames[anim.FrameIndex]
	char := p.Character

	// x := p.X - p.Width/2
	// y := p.Y
	// w := p.Width
	// h := p.Height
	// vector.StrokeRect(screen, float32(x), float32(y), float32(w), float32(h), 1, color.RGBA{255, 0, 0, 255}, false)

	frameWidth := float64(currentFrame.Bounds().Dx())
	op := &ebiten.DrawImageOptions{}

	if p.Direction < 0 {
		op.GeoM.Scale(-char.ScaleX, char.ScaleY)
		op.GeoM.Translate(p.X+(frameWidth/2)*char.ScaleX, p.Y)
	} else {
		op.GeoM.Scale(char.ScaleX, char.ScaleY)
		op.GeoM.Translate(p.X-(frameWidth/2)*char.ScaleX, p.Y)
	}
	screen.DrawImage(currentFrame, op)
}

func (p *Controllable) ChangeAnim(animationName string) {
	if p.CurrentAnim != animationName {
		anim := p.Character.Animations[p.CurrentAnim]
		anim.FrameIndex = 0
		anim.FrameTick = 0
		p.PrevAnim = p.CurrentAnim
		p.CurrentAnim = animationName
	}
}