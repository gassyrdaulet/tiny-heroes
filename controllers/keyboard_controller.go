package controllers

import (
	"github.com/gassyrdaulet/go-fighting-game/base"
	"github.com/hajimehoshi/ebiten/v2"
)

type KeyboardController struct {
	// клавиши
	Left, Right, Up, Down, Attack ebiten.Key

	prevKeys map[ebiten.Key]bool
}

func NewKeyboardController(left, right, up, down, attack ebiten.Key) *KeyboardController {
	return &KeyboardController{
		Left: left, Right: right, Up: up, Down: down, Attack: attack,
		prevKeys: make(map[ebiten.Key]bool),
	}
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

	attackPressed := ebiten.IsKeyPressed(c.Attack)
	if attackPressed && !c.prevKeys[c.Attack] {
		input.Attack = true
	}
	c.prevKeys[c.Attack] = attackPressed

	return input
}
