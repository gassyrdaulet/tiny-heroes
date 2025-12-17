package base

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {}

func (e *Entity) Update(entities []GameEntity) error {
	return nil
}

func (e *Entity) Draw(screen *ebiten.Image) error {
	return nil
}
