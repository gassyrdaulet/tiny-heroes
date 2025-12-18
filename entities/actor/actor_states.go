package actor

import (
	"github.com/gassyrdaulet/go-fighting-game/base"
)

const (
	Idle       base.State = "idle"
	Run        base.State = "run"
	Jump       base.State = "jump"
	Fall       base.State = "fall"
	Attack     base.State = "attack"
	ChargeJump base.State = "charging_jump"
)