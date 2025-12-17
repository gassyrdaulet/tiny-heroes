package controllers

import (
	"github.com/gassyrdaulet/go-fighting-game/base"
	"github.com/hajimehoshi/ebiten/v2"
)
type KeyboardController struct {
	Left  ebiten.Key
	Right ebiten.Key
	Jump  ebiten.Key
}

func (c *KeyboardController) GetInput() base.Input {
	var input base.Input

	if ebiten.IsKeyPressed(c.Left) {
		input.MoveX = -1
	}
	if ebiten.IsKeyPressed(c.Right) {
		input.MoveX = 1
	}
	if ebiten.IsKeyPressed(c.Jump) {
		input.Jump = true
	}

	return input
}