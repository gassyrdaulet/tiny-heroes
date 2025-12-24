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
	"github.com/gassyrdaulet/go-fighting-game/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	state 	*base.StateMachine
	introTimer int
	world   *physics.World
	tileMap *base.TileMap
	camera  *base.Camera
	bg      *base.Background
	players []*actor.Actor
	levelName string
	input *Input
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
    if g.input.JustPressed(ebiten.KeyEnter) {
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
		g.drawMainMenu(screen)

	case StatePlaying:
		g.drawWorld(screen)
		g.drawHUD(screen)

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
		"TINY HEROES\n\nPress any key",
		constants.ScreenW/2-80,
		constants.ScreenH/2,
	)
}

func (g *Game) drawMainMenu(screen *ebiten.Image) {
	text := "MAIN MENU\n\n[Enter] Start Game\n[Esc] Exit"

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

func (g *Game) drawHUD(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("Players: %d", len(g.players)),
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
		state: base.NewStateMachine(StateIntro),
		input: NewInput(),
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
			g.initializeNewGame()
		}

	case StatePaused:
		g.pause()
	}
}

func (g *Game) initIntro() {
	g.introTimer = 0
}

func (g *Game) initializeNewGame() {
	// Камера
	g.camera = &base.Camera{
		Width:  constants.ScreenW,
		Height: constants.ScreenH,
	}
	// Карта
	tileMap := base.NewTileMap(100, 300, constants.TileSize)

	tile1, _ := utils.LoadImage("assets/tiles/tile1.png")
	tile2, _ := utils.LoadImage("assets/tiles/tile2.png")
	tile3, _ := utils.LoadImage("assets/tiles/tile3.png")
	tile8, _ := utils.LoadImage("assets/tiles/tile8.png")
	tile9, _ := utils.LoadImage("assets/tiles/tile9.png")
	tile10, _ := utils.LoadImage("assets/tiles/tile10.png")
	tile12, _ := utils.LoadImage("assets/tiles/tile12.png")
	tile13, _ := utils.LoadImage("assets/tiles/tile13.png")
	tile22, _ := utils.LoadImage("assets/tiles/tile22.png")
	tile23, _ := utils.LoadImage("assets/tiles/tile23.png")
	tile24, _ := utils.LoadImage("assets/tiles/tile24.png")
	tile25, _ := utils.LoadImage("assets/tiles/tile25.png")
	tile26, _ := utils.LoadImage("assets/tiles/tile26.png")
	tile27, _ := utils.LoadImage("assets/tiles/tile27.png")
	tile39, _ := utils.LoadImage("assets/tiles/tile39.png")
	tile40, _ := utils.LoadImage("assets/tiles/tile40.png")
	tile41, _ := utils.LoadImage("assets/tiles/tile41.png")
	tile42, _ := utils.LoadImage("assets/tiles/tile42.png")

	tileMap.AddTileType(1, base.Solid, tile1)
	tileMap.AddTileType(2, base.Solid, tile2)
	tileMap.AddTileType(3, base.Solid, tile3)
	tileMap.AddTileType(8, base.Solid, tile8)
	tileMap.AddTileType(9, base.Solid, tile9)
	tileMap.AddTileType(10, base.Solid, tile10)
	tileMap.AddTileType(12, base.Solid, tile12)
	tileMap.AddTileType(13, base.Solid, tile13)
	tileMap.AddTileType(22, base.Solid, tile22)
	tileMap.AddTileType(22, base.Solid, tile22)
	tileMap.AddTileType(23, base.Solid, tile23)
	tileMap.AddTileType(24, base.Solid, tile24)
	tileMap.AddTileType(25, base.Solid, tile25)
	tileMap.AddTileType(26, base.Solid, tile26)
	tileMap.AddTileType(27, base.Solid, tile27)
	tileMap.AddTileType(39, base.Solid, tile39)
	tileMap.AddTileType(40, base.Solid, tile40)
	tileMap.AddTileType(41, base.Solid, tile41)
	tileMap.AddTileType(42, base.Solid, tile42)

	for _, x := range []int{1, 2, 3, 4, 9, 10, 11, 12} {
		tileMap.SetTile(x, 17, 2)
	}
	for x := range 60 {
		tileMap.SetTile(x+13, 17, 2)
		tileMap.SetTile(x+13, 18, 9)
	}
	for _, x := range []int{1, 2, 3, 4, 9, 10, 11, 12} {
		tileMap.SetTile(x, 18, 9)
	}
	for x := range 72 {
		tileMap.SetTile(x+1, 19, 9)
	}
	for x := range 99 {
		tileMap.SetTile(x+1, 199, 2)
	}
	for x := range 40 {
		tileMap.SetTile(x, 298, 2)
	}

	tileMap.SetTile(5, 17, 3)
	tileMap.SetTile(8, 17, 1)
	tileMap.SetTile(5, 18, 12)
	tileMap.SetTile(6, 18, 2)
	tileMap.SetTile(7, 18, 2)
	tileMap.SetTile(8, 18, 13)

	tileMap.SetTile(16, 15, 22)
	tileMap.SetTile(17, 15, 24)
	tileMap.SetTile(19, 13, 22)
	tileMap.SetTile(20, 13, 23)
	tileMap.SetTile(21, 13, 24)
	tileMap.SetTile(16, 11, 24)
	tileMap.SetTile(15, 11, 22)
	tileMap.SetTile(14, 9, 25)
	tileMap.SetTile(15, 6, 22)
	tileMap.SetTile(16, 6, 24)
	
	GroundY := float64(tileMap.Height * constants.TileSize)

	g.bg = &base.Background{
		Layers: []*base.BackgroundLayer{
			// 8 — небо
			{
				Image:    utils.MustLoad("assets/backgrounds/1/8.png"),
				ScrollX:  0.0,
				ScrollY:  0.0,
				BaseY:    GroundY,
				StretchY: true,
			},

			// 7 — облака
			{
				Image:   utils.MustLoad("assets/backgrounds/1/7.png"),
				ScrollX: 0.05,
				ScrollY: 0.03,
				BaseY:   GroundY,
			},

			// 6 — дальние горы
			{
				Image:   utils.MustLoad("assets/backgrounds/1/6.png"),
				ScrollX: 0.10,
				ScrollY: 0.05,
				BaseY:   GroundY,
			},

			// 5 — дальний горизонт
			{
				Image:   utils.MustLoad("assets/backgrounds/1/5.png"),
				ScrollX: 0.20,
				ScrollY: 0.06,
				BaseY:   GroundY,
			},

			// 4 — средний горизонт
			{
				Image:   utils.MustLoad("assets/backgrounds/1/4.png"),
				ScrollX: 0.30,
				ScrollY: 0.07,
				BaseY:   GroundY,
			},

			// 3 — ближний горизонт
			{
				Image:   utils.MustLoad("assets/backgrounds/1/3.png"),
				ScrollX: 0.45,
				ScrollY: 0.08,
				BaseY:   GroundY,
			},

			// 2 — очень близкий горизонт
			{
				Image:   utils.MustLoad("assets/backgrounds/1/2.png"),
				ScrollX: 0.65,
				ScrollY: 0.09,
				BaseY:   GroundY,
			},

			// 1 — трава (передний план)
			{
				Image:   utils.MustLoad("assets/backgrounds/1/1.png"),
				ScrollX: 0.90,
				ScrollY: 1,
				BaseY:   GroundY,
			},
		},
	}
	g.tileMap = tileMap

	// Мир
	g.world = physics.NewWorld(g.tileMap)

	// Игроки
	g.players = []*actor.Actor{
		actor.NewActor(
			100,
			100,
			"blue",
			&controllers.KeyboardController{
				Left:  ebiten.KeyLeft,
				Right: ebiten.KeyRight,
				Up:    ebiten.KeyUp,
				Down:  ebiten.KeyDown,
			},
			1,
		),
		// actor.NewActor(
		// 	150,
		// 	100,
		// 	"pink",
		// 	&controllers.KeyboardController{
		// 		Left:  ebiten.KeyA,
		// 		Right: ebiten.KeyD,
		// 		Up:    ebiten.KeyW,
		// 		Down:  ebiten.KeyD,
		// 	},
		// 	1,
		// ),
	}

	g.levelName = "level1"
}

func (g *Game) loadLevel(level string) {
	// очищаем старое
	g.players = nil
	g.tileMap = nil
	g.world.Clear()

	// грузим новое
	g.tileMap = base.NewTileMap(100, 200, constants.TileSize)

	g.levelName = level
}

func (g *Game) pause() {
	if g.world != nil {
		g.world.SetPaused(true)
	}
}

func (g *Game) mainMenu() {
	g.players = nil
	g.world = nil
	g.tileMap = nil
	g.bg = nil
	g.camera = nil
}

func main() {
	ebiten.SetTPS(60)
	// ebiten.SetWindowSize(constants.WindowW, constants.WindowH)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Tiny Heroes")

	if err := characters.LoadCharacterAnimations(); err != nil {
		log.Fatal(err)
	}

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
