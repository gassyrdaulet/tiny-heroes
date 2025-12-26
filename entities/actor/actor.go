package actor

import (
	"fmt"
	"image"
	"image/color"

	"github.com/gassyrdaulet/go-fighting-game/base"
	"github.com/gassyrdaulet/go-fighting-game/base/physics"
	"github.com/gassyrdaulet/go-fighting-game/characters"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type AnimationName string

const (
	debug	   bool 		 = true
	Idle       AnimationName = "idle"
	Run        AnimationName = "run"
	Jump       AnimationName = "jump"
	Fall       AnimationName = "fall"
	Attack     AnimationName = "attack"
	JumpAttack AnimationName = "jump_attack"
	FallAttack AnimationName = "fall_attack"
	RunAttack  AnimationName = "run_attack"
	ChargeJump AnimationName = "charging_jump"
	Dying	   AnimationName = "dying"
	Hurting	   AnimationName = "hurt"
	None	   AnimationName = "none"
)

type Actor struct {
	*physics.Body
	*base.Animator
	Character         		*characters.Character
	Controller        		base.Controller
	MaxHp           		int
	Hp              		int
	Dead 			  		bool
	Speed           		float64
	JumpForce       		float64
	JumpForceCurrent		float64
	Direction       		int
	ChargingJump			bool
	ChargingJumpTicksMax 	int
	ChargingJumpTicks 		int
	Attacking				bool
	AttackTicksMax			int
	AttackTicks				int
	AttackCooldownTicksMax 	int
	AttackCooldownTicks 	int
	Dying					bool
	DyingTicksMax			int
	DyingTicks				int
	Hurting					bool
	HurtingTicksMax			int
	HurtingTicks			int
	AttackRange				float64
}

func (a *Actor) GoLeft() {
	if a.OnGround {
		a.VX = -a.Speed
	} else {
		a.VX = -a.Speed * 5 / 6
	}
	a.Direction = -1
}

func (a *Actor) GoRight() {
	if a.OnGround {
		a.VX = a.Speed
	} else {
		a.VX = a.Speed * 5 / 6
	}
	a.Direction = 1
}

func (a *Actor) TakeDamage(amount int, from *Actor) {
	if a.Dying || a.Dead || a.Hurting {
		return
	}
	a.VX = 0
	a.Hurting = true
	a.HurtingTicks = a.HurtingTicksMax
	a.Hp -= amount
	if a.Hp <= 0 {
		a.Hp = 0
		a.Die()
	}
}

func (a *Actor) Die() {
	if !a.Dead && !a.Dying {
		a.VY = 0
		a.VX = 0
		a.Dying = true
		a.DyingTicks = a.DyingTicksMax
	}
}

func (a *Actor) ChargeJump() {
	if a.OnGround && !a.ChargingJump && !a.Attacking {
		a.ChargingJump = true
		a.ChargingJumpTicks = a.ChargingJumpTicksMax
		if a.VX != 0 {
			a.JumpForceCurrent = a.JumpForce * 1.1
		} else {
			a.JumpForceCurrent = a.JumpForce
		}
		a.VX = 0
	}
}

func (a *Actor) StartAttack() {
	if a.AttackCooldownTicks <= 0 && a.AttackTicks <= 0 && !a.Attacking {
		a.Attacking = true
		a.AttackTicks = a.AttackTicksMax
	}
}

func (a *Actor) Attack(others []*Actor) {
	damage := a.Character.Damage
	hitbox := a.AttackHitBox()
	for _, o := range others {
		if o == a || o.Dead {
			continue
		}
		if hitbox.Overlaps(o.HitBox()) {
			o.TakeDamage(damage, a)
		}
	}
	a.AttackCooldownTicks = a.AttackCooldownTicksMax
}

func (a *Actor) AttackHitBox() image.Rectangle {
	dir := a.Direction

	x := a.X + (a.Character.Width/2)*float64(dir)
	
	return image.Rect(
		int(x),
		int(a.Y+5),
		int(x+a.AttackRange),
		int(a.Y+a.Character.Height-5),
	)
}

func (a *Actor) HitBox() image.Rectangle {
	return image.Rect(
		int(a.X),
		int(a.Y),
		int(a.X+a.Character.Width),
		int(a.Y+a.Character.Height),
	)
}

func (a *Actor) Update(world *physics.World, players []*Actor, friendlyFire bool) {
	if !a.ChargingJump && !a.Hurting && !a.Dying && !a.Dead {
		input := a.Controller.GetInput()
		if input.Left {
			a.GoLeft()
		} else if input.Right {
			a.GoRight()
		} else {
			a.VX = 0
		}
		if input.Attack {
			a.StartAttack()
		}
		if input.Up {
			a.ChargeJump()
		}
	}

	if a.AttackCooldownTicks > 0 {
		a.AttackCooldownTicks--
	}
	if a.AttackTicks > 0 {
		a.AttackTicks--
	} else {
		if a.Attacking {
			if friendlyFire {
				a.Attack(players)
			} 
			a.AttackTicks = 0
			a.Attacking = false
		}
	}
	if a.HurtingTicks > 0 {
		a.HurtingTicks--
	} else {
		if a.Hurting {
			a.HurtingTicks = 0 
			a.Hurting = false
		}
	}
	if a.DyingTicks > 0 {
		a.DyingTicks--
	} else {
		if a.Dying {
			a.Dying = false
			a.DyingTicks = 0
			a.Dead = true
		}
	}
	if a.ChargingJumpTicks > 0 {
		a.ChargingJumpTicks--
	} else {
		if a.OnGround && a.ChargingJump && a.ChargingJumpTicks <= 0 {
			a.ChargingJumpTicks = 0
			a.VY = a.JumpForceCurrent
			a.ChargingJump = false
		}
	}

	world.Step(a)

	a.UpdateFrame(string(a.UpdateAnimation()))
}

func (a *Actor) UpdateAnimation() AnimationName {
	if a.Dead {
		return None
	}
	if a.Dying {
		return Dying
	}
	if a.Hurting {
		return Hurting
	}
	if a.ChargingJumpTicks > 0 {
		return ChargeJump
	}

	if a.Attacking {
		switch a.CurrentAnimation {
		case string(Attack), string(RunAttack), string(JumpAttack):
			return AnimationName(a.CurrentAnimation)
		default:
			if a.OnGround {
				if a.VX == 0 {
					return Attack
				}
				return RunAttack
			} else {
				return JumpAttack
			}
		}
	}

	if a.OnGround {
		if a.VX == 0 {
			return Idle
		}
		return Run
	}

	if a.VY < 0 {
		return Jump
	}
	return Fall
}

 
func (a *Actor) Draw(screen *ebiten.Image, sx, sy float64) {
	if debug {
		a.DrawDebug(screen, sx, sy)
	}
	a.DrawHPBar(screen, sx, sy)
	a.DrawFrame(screen, sx, sy, a.Direction == -1)
}

func (a *Actor) IsAlive() bool{
	return !a.Dead
}

func (a *Actor) DrawHPBar(screen *ebiten.Image, sx, sy float64) {
	w := 30
	h := 4

	ratio := float64(a.Hp) / float64(a.MaxHp)
	newWidth := int(float64(w)*ratio)
	if newWidth > 0 {
		bar := ebiten.NewImage(newWidth, h)
		bar.Fill(color.RGBA{200, 45, 45, 255})
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(sx-float64(w/2), sy-10)

		screen.DrawImage(bar, op)
	}
}

func (a *Actor) DrawDebug(screen *ebiten.Image, sx, sy float64) {
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("Animation & Attacking: (%s, %t)", a.CurrentAnimation, a.Attacking),
		int(sx-30),
		int(sy-30),
	)
}