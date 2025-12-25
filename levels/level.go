package levels

import (
	"encoding/json"
	"os"
	"sort"
	"strconv"

	"github.com/gassyrdaulet/go-fighting-game/base"
	c "github.com/gassyrdaulet/go-fighting-game/characters"
	"github.com/gassyrdaulet/go-fighting-game/constants"
	"github.com/gassyrdaulet/go-fighting-game/entities/actor"
	"github.com/gassyrdaulet/go-fighting-game/levels/tileset"
	"github.com/gassyrdaulet/go-fighting-game/utils"
)

type LevelData struct {
	TileMap    TileMapData     `json:"tilemap"`
	Background []BackgroundDef `json:"background"`
	Spawns     []SpawnPoint    `json:"spawns"`
}

type TileMapData struct {
	Width    int            `json:"width"`
	Height   int            `json:"height"`
	Tileset  string         `json:"tileset"`
	Lines    map[string]string `json:"lines"`
}

type BackgroundDef struct {
	Image    string  `json:"image"`
	ScrollX  float64 `json:"scrollX"`
	ScrollY  float64 `json:"scrollY"`
	StretchY bool    `json:"stretchY"`
}

type SpawnPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func LoadLevel(levelName string) (*LevelData, error) {
	data, err := os.ReadFile(constants.LevelsDirectory + levelName + ".json")
	if err != nil {
		return nil, err
	}

	var level LevelData
	if err := json.Unmarshal(data, &level); err != nil {
		return nil, err
	}

	return &level, nil
}

func BuildTileMapFromLines(data TileMapData) *base.TileMap {
	tileMap := base.NewTileMap(
		data.Width,
		data.Height,
		constants.TileSize,
	)

	tileset.LoadTileSetFromJSON(tileMap, data.Tileset)

	symbolToID := map[rune]int{
		'=': 2,
		'\\': 3,
		'/': 1,
		'*': 9,
		'}': 10,
		'{': 8,
		'|': 19,
		'0': 4,
		'1': 25,
		'(': 13,
		')': 12,
		'<': 22,
		'_': 23,
		'>': 24,
		'«': 36,
		'-': 37,
		'»': 38,
		'~': 21,
		' ': 0,
	}

	for yStr, line := range data.Lines {
		y, err := strconv.Atoi(yStr)
		if err != nil {
			continue
		}
		runes := []rune(line)
		for x := range runes {
			char := runes[x]
			id, ok := symbolToID[char]
			if !ok {
				id = 0
			}
			tileMap.SetTile(x, y, id)
		}
	}

	return tileMap
}

func BuildBackground(defs []BackgroundDef, groundY float64) *base.Background {
    layers := []*base.BackgroundLayer{}

    for _, bg := range defs {
        layers = append(layers, &base.BackgroundLayer{
            Image:    utils.MustLoad(bg.Image),
            ScrollX:  bg.ScrollX,
            ScrollY:  bg.ScrollY,
            StretchY: bg.StretchY,
        })
    }

    return &base.Background{Layers: layers, BaseY: groundY}
}

func SpawnPlayers(
    controllers []base.Controller,
    characters map[string]*c.Character,
    spawns []SpawnPoint,
) []*actor.Actor {
	if len(controllers) == 0 || len(characters) == 0 {
		return nil
	}

	charIDs := make([]string, 0, len(characters))
	for id := range characters {
		charIDs = append(charIDs, id)
	}

    sort.Strings(charIDs)

    players := make([]*actor.Actor, 0, len(controllers))

    for i, ctrl := range controllers {
        spawn := spawns[0]
        if i < len(spawns) {
            spawn = spawns[i]
        }

        charID := charIDs[i % len(charIDs)]

        players = append(players,
            actor.NewActor(
                spawn.X,
                spawn.Y,
                ctrl,
                1,
                characters[charID],
            ),
        )
    }

    return players
}