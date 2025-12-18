package base

import "github.com/hajimehoshi/ebiten/v2"

type TileType int

const (
	Empty TileType = iota
	Solid
	Platform
	Hazard
)

type Tile struct {
	ID    int
	Type  TileType
	Image *ebiten.Image
}

type TileMap struct {
	Width, Height int
	Tiles         [][]*Tile
	Textures      map[int]*ebiten.Image
	TileTypes     map[int]TileType
	TileSize      int
}

func NewTileMap(width, height, tileSize int) *TileMap {
	tiles := make([][]*Tile, height)
	for y := range tiles {
		tiles[y] = make([]*Tile, width)
		for x := range tiles[y] {
			tiles[y][x] = &Tile{ID: 0, Type: Empty, Image: nil}
		}
	}

	return &TileMap{
		Width:     width,
		Height:    height,
		Tiles:     tiles,
		Textures:  make(map[int]*ebiten.Image),
		TileTypes: make(map[int]TileType),
		TileSize:  tileSize,
	}
}

func (m *TileMap) AddTileType(id int, t TileType, img *ebiten.Image) {
	m.Textures[id] = img
	m.TileTypes[id] = t
}

func (m *TileMap) SetTile(x, y, id int) {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return
	}
	m.Tiles[y][x].ID = id
	m.Tiles[y][x].Type = m.TileTypes[id]
	m.Tiles[y][x].Image = m.Textures[id]
}

func (m *TileMap) IsSolid(tx, ty int) bool {
	if ty < 0 || ty >= m.Height || tx < 0 || tx >= m.Width {
		return false
	}
	tile := m.Tiles[ty][tx]
	return tile.Type == Solid
}

func (m *TileMap) Draw(screen *ebiten.Image, cam *Camera) {
	camX, camY := cam.TopLeft()

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			tile := m.Tiles[y][x]
			if tile.ID == 0 || tile.Image == nil {
				continue
			}

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(
				float64(x*m.TileSize)-camX,
				float64(y*m.TileSize)-camY,
			)
			screen.DrawImage(tile.Image, op)
		}
	}
}
