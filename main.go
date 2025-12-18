package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/gassyrdaulet/go-fighting-game/base"
	"github.com/gassyrdaulet/go-fighting-game/base/physics"
	"github.com/gassyrdaulet/go-fighting-game/characters"
	"github.com/gassyrdaulet/go-fighting-game/constants"
	"github.com/gassyrdaulet/go-fighting-game/controllers"
	"github.com/gassyrdaulet/go-fighting-game/entities/actor"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	world   *physics.World
	tileMap *base.TileMap
	camera *base.Camera

	players []*actor.Actor
}

func (g *Game) Update() error {
	for _, a := range g.players {
		a.Update(g.world)
	}

	playersPos := make([]base.PlayerPosition, 0, len(g.players))
	for _, p := range g.players {
		playersPos = append(playersPos, p)
	}
	g.camera.UpdateFromPlayers(playersPos)

	g.world.UpdateVirtualBounds(g.camera)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{30, 30, 50, 255})

	for y := 0; y < g.tileMap.Height; y++ {
		for x := 0; x < g.tileMap.Width; x++ {
			if g.tileMap.Tiles[y][x] == 1 {
				screenX, screenY := g.camera.WorldToScreen(
					float64(x*constants.TileSize), 
					float64(y*constants.TileSize),
				)
				vector.FillRect(
					screen,
					float32(screenX),
					float32(screenY),
					constants.TileSize,
					constants.TileSize,
					color.RGBA{80, 80, 80, 255},
					false,
				)
			}
		}
	}

	for _, a := range g.players {
		sx, sy := g.camera.WorldToScreen(a.X, a.Y)
		a.Draw(screen, sx, sy)
	}

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("FPS: %.0f X: %.0f Y: %.0f", ebiten.ActualFPS(), g.players[0].X, g.players[0].Y),
	)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return constants.ScreenW, constants.ScreenH
}

func main() {
	ebiten.SetTPS(60)
	ebiten.SetWindowSize(constants.WindowW, constants.WindowH)
	ebiten.SetWindowTitle("Tiny Heroes")

	if err := characters.LoadCharacterAnimations(); err != nil {
		log.Fatal(err)
	}

	tileMap := &base.TileMap{
		Width:  100,
		Height: 20,
		Tiles:  make([][]int, 20),
	}

	for y := range 20 {
		tileMap.Tiles[y] = make([]int, 100)
	}

	for x := range 100 {
		tileMap.Tiles[18][x] = 1
	}

	for y := range 10 {
		tileMap.Tiles[18-y][1] = 1
	}

	tileMap.Tiles[15][19] = 1

	tileMap.Tiles[17][14] = 1

	world := physics.NewWorld(tileMap)

	actor1 := actor.NewActor(
		100,
		100,
		"blue",
		&controllers.KeyboardController{
			Left: ebiten.KeyLeft,
			Right: ebiten.KeyRight,
			Up: ebiten.KeyUp,
			Down: ebiten.KeyDown,
		},
		1,
	)
	actor2 := actor.NewActor(
		150,
		100,
		"pink",
		&controllers.KeyboardController{
			Left: ebiten.KeyA,
			Right: ebiten.KeyD,
			Up: ebiten.KeyW,
			Down: ebiten.KeyD,
		},
		1,
	)

	game := &Game{
		world:   world,
		tileMap: tileMap,
		camera: &base.Camera{
			Width: constants.ScreenW,
			Height: constants.ScreenH,
		},
		players: []*actor.Actor{
			actor1,
			actor2,
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
