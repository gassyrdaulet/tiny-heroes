package base

type TileMap struct {
	Width, Height int
	Tiles         [][]int
}

func (m *TileMap) IsSolid(tx, ty int) bool {
	if ty < 0 || ty >= len(m.Tiles) {
		return false
	}
	if tx < 0 || tx >= len(m.Tiles[0]) {
		return false
	}
	return m.Tiles[ty][tx] == 1
}