package base

type TileMap interface {
	IsSolid(tx, ty int) bool
}
