package tileset

import (
	"encoding/json"
	"os"

	"github.com/gassyrdaulet/go-fighting-game/base"
	"github.com/gassyrdaulet/go-fighting-game/constants"
	"github.com/gassyrdaulet/go-fighting-game/utils"
)

type TileSetJSON struct {
	Tiles []TileDefJSON `json:"tiles"`
}

type TileType string

const (
    TileSolid    TileType = "solid"
    TilePlatform TileType = "platform"
    TileDecor    TileType = "decor"
)

type TileDefJSON struct {
	ID    int       	`json:"id"`
	Type  TileType      `json:"type"`
	Image string    	`json:"image"`
}

func LoadTileSetFromJSON(
    tileMap *base.TileMap,
    tilesetName string,
) error {
    data, err := os.ReadFile(constants.TilesetDirectory + tilesetName + ".json")
    if err != nil {
        return err
    }

    var tileset TileSetJSON
    if err := json.Unmarshal(data, &tileset); err != nil {
        return err
    }

    for _, tile := range tileset.Tiles {
        img := utils.MustLoad(tile.Image)

        var collision base.TileType
        switch tile.Type {
        case TileSolid:
            collision = base.Solid
        case TilePlatform:
            collision = base.Platform
        default:
            collision = base.Empty
        }

        tileMap.AddTileType(tile.ID, collision, img)
    }

    return nil
}