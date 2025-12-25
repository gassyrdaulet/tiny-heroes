package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/gassyrdaulet/go-fighting-game/base"
	"github.com/gassyrdaulet/go-fighting-game/base/physics"
	"github.com/gassyrdaulet/go-fighting-game/characters"
	"github.com/gassyrdaulet/go-fighting-game/constants"
	"github.com/gassyrdaulet/go-fighting-game/controllers"
	"github.com/gassyrdaulet/go-fighting-game/entities/actor"
	"github.com/gassyrdaulet/go-fighting-game/levels"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	state       	*base.StateMachine
	introTimer  	int
	world       	*physics.World
	tileMap     	*base.TileMap
	camera      	*base.Camera
	bg          	*base.Background
	players     	[]*actor.Actor
	playersChars	map[string]*characters.Character
	controllers 	[]base.Controller
	levelName   	string
	input       	*Input
	full_screen 	bool
}

type Level struct {
	World   *physics.World
	TileMap *base.TileMap
	BG      *base.Background
	Players []*actor.Actor
}

type Input struct {
	prev map[ebiten.Key]bool
}

func NewInput() *Input {
	return &Input{
		prev: make(map[ebiten.Key]bool),
	}
}

func (i *Input) JustPressed(key ebiten.Key) bool {
	pressed := ebiten.IsKeyPressed(key)
	wasPressed := i.prev[key]

	i.prev[key] = pressed

	return pressed && !wasPressed
}

const (
	StateIntro    base.State = "intro"
	StateMainMenu base.State = "main_menu"
	StatePlaying  base.State = "playing"
	StatePaused   base.State = "paused"
)

func (g *Game) Update() error {
	if g.input.JustPressed(ebiten.KeyF11) {
		g.setFullScreen(!g.full_screen)
	}

	switch g.state.CurrentState {

	case StateIntro:
		g.updateIntro()

	case StateMainMenu:
		g.updateMenu()

	case StatePlaying:
		g.updateGame()

	case StatePaused:
		g.updatePause()
	}

	return nil
}

func (g *Game) updateMenu() {
	if g.input.JustPressed(ebiten.Key1) {
		g.levelName = "ai-arena"
		g.state.ChangeState(StatePlaying)
	}

	if g.input.JustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
}

func (g *Game) updatePause() {
	if g.input.JustPressed(ebiten.KeyEscape) {
		g.state.ChangeState(StatePlaying)
	}

	if g.input.JustPressed(ebiten.KeyM) {
		g.state.ChangeState(StateMainMenu)
	}
}

func (g *Game) updateIntro() {
	g.introTimer++

	if g.introTimer > 300 {
		g.state.ChangeState(StateMainMenu)
	}

	if g.input.JustPressed(ebiten.KeyEnter) {
		g.state.ChangeState(StateMainMenu)
	}
}

func (g *Game) updateGame() {
	for _, a := range g.players {
		a.Update(g.world)
	}

	playersPos := make([]base.PlayerPosition, 0, len(g.players))
	for _, p := range g.players {
		playersPos = append(playersPos, p)
	}

	g.camera.UpdateFromPlayers(playersPos, g.world.Width, g.world.Height)
	g.world.UpdateVirtualBounds(g.camera)

	if g.input.JustPressed(ebiten.KeyEscape) {
		g.state.ChangeState(StatePaused)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state.CurrentState {

	case StateIntro:
		g.drawIntro(screen)

	case StateMainMenu:
		g.drawControllersAmount(screen)
		g.drawMainMenu(screen)

	case StatePlaying:
		g.drawWorld(screen)

	case StatePaused:
		g.drawWorld(screen)
		g.drawPauseOverlay(screen)
		g.drawPauseMenu(screen)
	}

	g.drawDebug(screen)
}

func (g *Game) drawWorld(screen *ebiten.Image) {
	if g.bg != nil {
		g.bg.Draw(screen, g.camera)
	}

	if g.tileMap != nil {
		g.tileMap.Draw(screen, g.camera)
	}

	for _, a := range g.players {
		sx, sy := g.camera.WorldToScreen(a.X, a.Y)
		a.Draw(screen, sx, sy)
	}
}

func (g *Game) drawIntro(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(
		screen,
		"https://github.com/gassyrdaulet/tiny-heroes",
		constants.ScreenW/2-120,
		constants.ScreenH/2,
	)
}

func (g *Game) drawMainMenu(screen *ebiten.Image) {
	text := "MAIN MENU\n\n[1] Start AI Battle\n[Esc] Exit"

	ebitenutil.DebugPrintAt(
		screen,
		text,
		constants.ScreenW/2-100,
		constants.ScreenH/2,
	)
}

func (g *Game) drawPauseOverlay(screen *ebiten.Image) {
	overlay := ebiten.NewImage(constants.ScreenW, constants.ScreenH)
	overlay.Fill(color.RGBA{0, 0, 0, 120})
	screen.DrawImage(overlay, nil)
}

func (g *Game) drawPauseMenu(screen *ebiten.Image) {
	text := "PAUSED\n\n[Esc] Resume\n[M] Main Menu"

	ebitenutil.DebugPrintAt(
		screen,
		text,
		constants.ScreenW/2-80,
		constants.ScreenH/2,
	)
}

func (g *Game) drawControllersAmount(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("Controllers: %d", len(g.controllers)),
		10,
		10,
	)
}

func (g *Game) drawDebug(screen *ebiten.Image) {
	if g.camera == nil {
		ebitenutil.DebugPrintAt(
			screen,
			fmt.Sprintf("FPS: %.0f | cam(nil)", ebiten.ActualFPS()),
			10,
			constants.ScreenH-20,
		)
		return
	}
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf(
			"FPS: %.0f | cam(%.0f, %.0f)",
			ebiten.ActualFPS(),
			g.camera.X,
			g.camera.Y,
		),
		10,
		constants.ScreenH-20,
	)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return constants.ScreenW, constants.ScreenH
}

func NewGame() *Game {
	g := &Game{
		state:       base.NewStateMachine(StateIntro),
		input:       NewInput(),
		full_screen: false,
	}
	g.state.OnChange = g.onStateChange
	return g
}

func (g *Game) onStateChange(prev, current base.State) {
	switch current {
	case StateIntro:
		g.initIntro()

	case StateMainMenu:
		g.mainMenu()

	case StatePlaying:
		if g.world == nil {
			g.initializeNewGame(g.levelName)
		}

	case StatePaused:
	}
}

func (g *Game) initIntro() {
	g.introTimer = 0
}

func (g *Game) initializeNewGame(initialLevelName string) {
	g.controllers = []base.Controller{
		&controllers.KeyboardController{
			Left: ebiten.KeyLeft,
			Right: ebiten.KeyRight,
			Down: ebiten.KeyDown,
			Up: ebiten.KeyUp,
		},
		&controllers.KeyboardController{
			Left: ebiten.KeyA,
			Right: ebiten.KeyD,
			Down: ebiten.KeyS,
			Up: ebiten.KeyW,
		},
		&controllers.KeyboardController{
			Left: ebiten.KeyJ,
			Right: ebiten.KeyL,
			Down: ebiten.KeyK,
			Up: ebiten.KeyI,
		},
	}
	playersChars, err := characters.LoadCharacters("characters/players.json")
	if err != nil {
		panic(err)
	}
	g.playersChars = playersChars
	g.camera = &base.Camera{
		Width:  constants.ScreenW,
		Height: constants.ScreenH,
	}
	g.loadLevel(initialLevelName)
}

func (g *Game) loadLevel(levelName string) {
	g.players = nil
	g.tileMap = nil
	if g.world != nil {
		g.world.Clear()
	}

	level, err := levels.LoadLevel(levelName); if err != nil {
		log.Fatal(err)
	}

	g.tileMap = levels.BuildTileMapFromLines(level.TileMap)

	g.bg = levels.BuildBackground(level.Background, float64(g.tileMap.Height * constants.TileSize))

	g.world = physics.NewWorld(g.tileMap)

	g.players = levels.SpawnPlayers(g.controllers, g.playersChars, level.Spawns)

	g.levelName = levelName
}

func (g *Game) mainMenu() {
	g.playersChars = nil
	g.players = nil
	g.world = nil
	g.tileMap = nil
	g.bg = nil
	g.camera = nil
	g.levelName = ""
}

func (g *Game) setFullScreen(value bool) {
	g.full_screen = value
	ebiten.SetFullscreen(value)
}

func main() {
	ebiten.SetTPS(60)
	ebiten.SetWindowSize(constants.WindowW, constants.WindowH)
	ebiten.SetWindowTitle("Tiny Heroes")

	game := NewGame()
	game.setFullScreen(false)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
