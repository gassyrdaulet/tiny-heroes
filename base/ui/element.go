package ui

import "github.com/hajimehoshi/ebiten/v2"

type Element interface {
	Update()
	Draw(screen *ebiten.Image)
}