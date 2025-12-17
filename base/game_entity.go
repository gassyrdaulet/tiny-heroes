package base

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameEntity interface {
	Update(entities []GameEntity) error
	Draw(screen *ebiten.Image)
}
