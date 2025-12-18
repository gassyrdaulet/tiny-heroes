package controllers

import (
	"github.com/gassyrdaulet/go-fighting-game/base"
	"github.com/hajimehoshi/ebiten/v2"
)
type KeyboardController struct {
	Left  ebiten.Key
	Right ebiten.Key
	Up  ebiten.Key
	Down  ebiten.Key
}

func (c *KeyboardController) GetInput() base.Input {
	var input base.Input

	if ebiten.IsKeyPressed(c.Left) {
		input.Left = true
	}
	if ebiten.IsKeyPressed(c.Right) {
		input.Right = true
	}
	if ebiten.IsKeyPressed(c.Up) {
		input.Up = true
	}
	if ebiten.IsKeyPressed(c.Down) {
		input.Down = true
	}

	return input
}