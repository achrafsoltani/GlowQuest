package world

import (
	"strings"

	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
)

type EnemySpawn struct {
	Type  int // maps to entity.EnemyType (0=Octorok, 1=Moblin, 2=Stalfos, 3=Boss)
	TileX int
	TileY int
}

type ItemSpawn struct {
	Type  int // maps to entity.ItemType
	TileX int
	TileY int
}

type NPCSpawn struct {
	ID                  string
	TileX               int
	TileY               int
	Dir                 int // maps to entity.Direction
	Name                string
	Dialogue            []string
	ConditionalDialogues []entity.DialogueOption
}

// ScreenWarp defines a warp point on a screen (door, stairs, etc.).
type ScreenWarp struct {
	TileX  int
	TileY  int
	Target string  // "interior:ID" or "overworld" or "dungeon:ID"
	SpawnX float64 // spawn position in target
	SpawnY float64
	ExitX  float64 // return position when exiting
	ExitY  float64
}

type Screen struct {
	Tiles       [config.ScreenGridH][config.ScreenGridW]TileType
	EnemySpawns []EnemySpawn
	ItemSpawns  []ItemSpawn
	NPCSpawns   []NPCSpawn
	Warps       []ScreenWarp
}

func (s *Screen) LoadFromString(data string) {
	rows := strings.Split(strings.TrimSpace(data), "\n")
	for y := 0; y < config.ScreenGridH && y < len(rows); y++ {
		row := rows[y]
		for x := 0; x < config.ScreenGridW && x < len(row); x++ {
			s.Tiles[y][x] = TileFromChar(row[x])
		}
	}
}

func (s *Screen) TileAt(gx, gy int) TileType {
	if gx < 0 || gx >= config.ScreenGridW || gy < 0 || gy >= config.ScreenGridH {
		return TileWall
	}
	return s.Tiles[gy][gx]
}
