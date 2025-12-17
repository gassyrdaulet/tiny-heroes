package factories

import (
	b "github.com/gassyrdaulet/go-fighting-game/base"
	c "github.com/gassyrdaulet/go-fighting-game/characters"
	e "github.com/gassyrdaulet/go-fighting-game/entities"
)

func NewPlayer(x, y float64, characterID string, controller *b.Controller, direction int) *e.Controllable {
	char := c.Characters[characterID]

	return &e.Controllable{
		Entity: &b.Entity{},
		Body: &b.Body{
			X: x,
			Y: y,
			Width: char.InitialWidth,
			Height: char.InitialHeight,
		},
		Character:   char,
		Hp:          char.MaxHP,
		Direction:   direction,
		CurrentAnim: "idle",
	}
}
