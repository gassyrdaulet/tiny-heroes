package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/gassyrdaulet/go-fighting-game/base"
	"github.com/gassyrdaulet/go-fighting-game/characters"
	"github.com/gassyrdaulet/go-fighting-game/constants"
	"github.com/gassyrdaulet/go-fighting-game/entities"
	"github.com/gassyrdaulet/go-fighting-game/factories"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	entities []base.GameEntity
	players  []*entities.Player
	ground   *entities.Ground
}

func (g *Game) Update() error {
	for _, entity := range g.entities {
		if err := entity.Update(g.entities); err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{30, 30, 50, 255})

	for _, entity := range g.entities {
		entity.Draw(screen)
	}

	fps := ebiten.ActualFPS()
	tps := ebiten.ActualTPS()

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("FPS: %.0f  TPS: %.0f P1 vy: %.0f", fps, tps, g.players[0].Vy()),
	)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return constants.ScreenW, constants.ScreenH
}

func main() {
	ebiten.SetTPS(60)

	if err := characters.LoadCharacterAnimations(); err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(constants.ScreenW, constants.ScreenH)
	ebiten.SetWindowTitle("Game")

	players := []*entities.Player{
		factories.NewPlayer(
			20,
			constants.ScreenH-constants.GroundHeight-400,
			"blue",
			&entities.InputKeys{
				LeftKey:  ebiten.KeyLeft,
				RightKey: ebiten.KeyRight,
				UpKey:    ebiten.KeyUp,
			},
			1,
		),
		factories.NewPlayer(
			200,
			constants.ScreenH-constants.GroundHeight-400,
			"pink",
			&entities.InputKeys{
				LeftKey:  ebiten.KeyA,
				RightKey: ebiten.KeyD,
				UpKey:    ebiten.KeyW,
			},
			-1,
		),
		factories.NewPlayer(
			400,
			constants.ScreenH-constants.GroundHeight-400,
			"white",
			&entities.InputKeys{
				LeftKey:  ebiten.KeyJ,
				RightKey: ebiten.KeyL,
				UpKey:    ebiten.KeyI,
			},
			1,
		),
	}

	ground := &entities.Ground{}

	entitiesList := []base.GameEntity{
		ground,
	}

	for _, p := range players {
		entitiesList = append(entitiesList, p)
	}

	game := &Game{
		entities: entitiesList,
		players:  players,
		ground:   ground,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
