package world

import (
	"strings"

	"github.com/AchrafSoltani/GlowQuest/config"
)

type EnemySpawn struct {
	Type  int // maps to entity.EnemyType (0=Octorok, 1=Moblin, 2=Stalfos)
	TileX int
	TileY int
}

type ItemSpawn struct {
	Type  int // maps to entity.ItemType
	TileX int
	TileY int
}

type NPCSpawn struct {
	TileX    int
	TileY    int
	Dir      int // maps to entity.Direction
	Name     string
	Dialogue []string
}

type Screen struct {
	Tiles       [config.ScreenGridH][config.ScreenGridW]TileType
	EnemySpawns []EnemySpawn
	ItemSpawns  []ItemSpawn
	NPCSpawns   []NPCSpawn
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

// VillageScreen is the starting screen at (1,1) centre of the overworld.
func VillageScreen() *Screen {
	s := &Screen{}
	s.LoadFromString(
		"......WWWW......\n" +
			".WWWW......WWWW.\n" +
			".WFFW.SSSS.WFFW.\n" +
			".WFFW.SSSS.WFFW.\n" +
			".WOOW......WOOW.\n" +
			"......S~~S......\n" +
			"..SSS.S~~S.SSS..\n" +
			"......SSSS......\n" +
			".WWWW......WWWW.\n" +
			".WFFW......WFFW.\n" +
			".WOOW......WOOW.\n" +
			"......SSSS......")

	// Sword near the well
	s.ItemSpawns = []ItemSpawn{
		{Type: 3, TileX: 8, TileY: 7}, // ItemSword
	}

	// NPCs
	s.NPCSpawns = []NPCSpawn{
		{TileX: 3, TileY: 7, Dir: 0, Name: "Old Man",
			Dialogue: []string{
				"It's dangerous to go",
				"alone! Take the sword",
				"by the well.",
			}},
		{TileX: 12, TileY: 3, Dir: 0, Name: "Merchant",
			Dialogue: []string{
				"Welcome to our village.",
				"The ruins to the south-",
				"east hold great treasure.",
			}},
	}

	return s
}

// ForestScreen is at (0,0) — dense trees with narrow grass paths.
func ForestScreen() *Screen {
	s := &Screen{}
	s.LoadFromString(
		"WWWWWWWWWWWWWWWW\n" +
			"W.TT.TT.TT.TT..\n" +
			"W...............\n" +
			"W.TT.TT..TT.TT.\n" +
			"W......TT.......\n" +
			"W.TT.......TT...\n" +
			"W.TT..TT...TT..\n" +
			"W......TT.......\n" +
			"W.TT.......TT...\n" +
			"W...............\n" +
			"W.TT..TT..TT...\n" +
			"......TT........")

	s.EnemySpawns = []EnemySpawn{
		{Type: 0, TileX: 5, TileY: 4},  // Octorok
		{Type: 0, TileX: 10, TileY: 8}, // Octorok
	}
	s.ItemSpawns = []ItemSpawn{
		{Type: 1, TileX: 8, TileY: 2},  // Rupee
		{Type: 0, TileX: 13, TileY: 5}, // Heart
	}
	return s
}

// ForestPathScreen is at (1,0) — tree-lined east-west path.
func ForestPathScreen() *Screen {
	s := &Screen{}
	s.LoadFromString(
		"WWWWWWWWWWWWWWWW\n" +
			"..TT..TT..TT...\n" +
			"................\n" +
			".TT...TT...TT..\n" +
			"................\n" +
			"SSSSSSSSSSSSSSSS\n" +
			"SSSSSSSSSSSSSSSS\n" +
			"................\n" +
			".TT...TT...TT..\n" +
			"................\n" +
			"..TT..TT..TT...\n" +
			"......SSSS......")

	s.EnemySpawns = []EnemySpawn{
		{Type: 1, TileX: 4, TileY: 2},  // Moblin
		{Type: 0, TileX: 12, TileY: 9}, // Octorok
	}
	s.NPCSpawns = []NPCSpawn{
		{TileX: 8, TileY: 8, Dir: 0, Name: "Traveller",
			Dialogue: []string{
				"Beware the mountain.",
				"Many Moblins lurk",
				"there.",
			}},
	}
	return s
}

// MountainScreen is at (2,0) — walls forming mountain shapes, narrow passages.
func MountainScreen() *Screen {
	s := &Screen{}
	s.LoadFromString(
		"WWWWWWWWWWWWWWWW\n" +
			"....WWWW..WWWW.W\n" +
			"..WWW..W....WW.W\n" +
			"..W.......WWW..W\n" +
			"........WW.....W\n" +
			"..WWW.......WW.W\n" +
			"..W...WWWW..WW.W\n" +
			".......W.......W\n" +
			"..WW...W..WWW..W\n" +
			"..WW..........WW\n" +
			"..........WW...W\n" +
			"......SSSS....WW")

	s.EnemySpawns = []EnemySpawn{
		{Type: 1, TileX: 5, TileY: 3},  // Moblin
		{Type: 1, TileX: 10, TileY: 7}, // Moblin
		{Type: 2, TileX: 3, TileY: 9},  // Stalfos
	}
	s.ItemSpawns = []ItemSpawn{
		{Type: 2, TileX: 8, TileY: 4}, // Key
	}
	return s
}

// LakeShoreScreen is at (0,1) — large water body on the right, sand beach.
func LakeShoreScreen() *Screen {
	s := &Screen{}
	s.LoadFromString(
		"W.....TT........\n" +
			"W..........SS~~~\n" +
			"W.TT.....SSS~~~~\n" +
			"W........SS~~~~~\n" +
			"W...TT..SS~~~~~~\n" +
			"W.......SS~~~~~~\n" +
			"W......SSS~~~~~~\n" +
			"W.......SS~~~~~~\n" +
			"W..TT...SS~~~~~~\n" +
			"W.......SSS~~~~~\n" +
			"W..........SS~~~\n" +
			"W.....SS........")

	s.EnemySpawns = []EnemySpawn{
		{Type: 0, TileX: 3, TileY: 3}, // Octorok
		{Type: 0, TileX: 5, TileY: 8}, // Octorok
	}
	s.ItemSpawns = []ItemSpawn{
		{Type: 4, TileX: 3, TileY: 6}, // HeartContainer
	}
	return s
}

// EastFieldScreen is at (2,1) — open grass with scattered trees and rocks.
func EastFieldScreen() *Screen {
	s := &Screen{}
	s.LoadFromString(
		"......SSSS.....W\n" +
			"...............W\n" +
			"..TT.......TT..W\n" +
			"...............W\n" +
			".......WW......W\n" +
			"...............W\n" +
			"..TT.......TT..W\n" +
			"...............W\n" +
			".....WW........W\n" +
			"...............W\n" +
			"...TT.....TT...W\n" +
			"......SSSS.....W")

	s.EnemySpawns = []EnemySpawn{
		{Type: 1, TileX: 6, TileY: 3},  // Moblin
		{Type: 0, TileX: 10, TileY: 8}, // Octorok
	}
	return s
}

// SwampScreen is at (0,2) — mix of water and grass in irregular patterns.
func SwampScreen() *Screen {
	s := &Screen{}
	s.LoadFromString(
		"W.....SS........\n" +
			"W..~~......~~...\n" +
			"W.~~~~..~.~~~~..\n" +
			"W..~~..~~~.~~...\n" +
			"W......~~~......\n" +
			"W..~........~...\n" +
			"W.~~~..~~..~~~..\n" +
			"W..~..~~~~..~...\n" +
			"W......~~.......\n" +
			"W..~~......~~...\n" +
			"W...............\n" +
			"WWWWWWWWWWWWWWWW")

	s.EnemySpawns = []EnemySpawn{
		{Type: 2, TileX: 5, TileY: 5},  // Stalfos
		{Type: 2, TileX: 11, TileY: 9}, // Stalfos
	}
	return s
}

// SouthFieldScreen is at (1,2) — open field with sand paths.
func SouthFieldScreen() *Screen {
	s := &Screen{}
	s.LoadFromString(
		"......SSSS......\n" +
			"......SSSS......\n" +
			"..TT..SSSS..TT..\n" +
			"......SSSS......\n" +
			"SSSSSSSSSSSSSSSS\n" +
			"................\n" +
			"................\n" +
			"SSSSSSSSSSSSSSSS\n" +
			"......SSSS......\n" +
			"..TT..SSSS..TT..\n" +
			"......SSSS......\n" +
			"WWWWWWWWWWWWWWWW")

	s.EnemySpawns = []EnemySpawn{
		{Type: 1, TileX: 4, TileY: 5}, // Moblin
	}
	return s
}

// RuinsScreen is at (2,2) — broken wall formations, floor tiles, stairs.
func RuinsScreen() *Screen {
	s := &Screen{}
	s.LoadFromString(
		"......SSSS.....W\n" +
			"..WWWW...WWWW..W\n" +
			"..WFFW...WFFW..W\n" +
			"..WFFW...WFFW..W\n" +
			"..W..W...W..W..W\n" +
			"...............W\n" +
			"..FFF..>..FFF..W\n" +
			"..FFF.....FFF..W\n" +
			"...............W\n" +
			"..WW..WWW..WW..W\n" +
			"..WW..WFW..WW..W\n" +
			"WWWWWWWWWWWWWWWW")

	s.EnemySpawns = []EnemySpawn{
		{Type: 2, TileX: 5, TileY: 5},  // Stalfos
		{Type: 2, TileX: 10, TileY: 5}, // Stalfos
		{Type: 1, TileX: 8, TileY: 8},  // Moblin
	}
	s.ItemSpawns = []ItemSpawn{
		{Type: 2, TileX: 4, TileY: 7},  // Key
		{Type: 1, TileX: 11, TileY: 7}, // Rupee
		{Type: 1, TileX: 12, TileY: 6}, // Rupee
	}
	s.NPCSpawns = []NPCSpawn{
		{TileX: 8, TileY: 6, Dir: 0, Name: "Ghost",
			Dialogue: []string{
				"You've found the",
				"ancient ruins. The",
				"stairs lead deeper...",
			}},
	}
	return s
}
