package world

import (
	"strings"

	"github.com/AchrafSoltani/GlowQuest/config"
)

type Screen struct {
	Tiles [config.ScreenGridH][config.ScreenGridW]TileType
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
// Houses (wall rectangles), paths (sand), central well (water).
// Open passages on all 4 edges (facing other screens).
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
	return s
}

// ForestScreen is at (0,0) — dense trees with narrow grass paths.
// Open on right edge (faces ForestPath) and bottom edge (faces LakeShore).
// Walls on top and left (world boundary).
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
	return s
}

// ForestPathScreen is at (1,0) — tree-lined east-west path.
// Open on left (faces Forest), right (faces Mountain), bottom (faces Village).
// Wall on top (world boundary).
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
	return s
}

// MountainScreen is at (2,0) — walls forming mountain shapes, narrow passages.
// Open on left (faces ForestPath), bottom (faces EastField).
// Walls on top and right (world boundary).
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
	return s
}

// LakeShoreScreen is at (0,1) — large water body on the right, sand beach.
// Open on top (faces Forest), right (faces Village), bottom (faces Swamp).
// Wall on left (world boundary).
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
	return s
}

// EastFieldScreen is at (2,1) — open grass with scattered trees and rocks.
// Open on left (faces Village), top (faces Mountain), bottom (faces Ruins).
// Wall on right (world boundary).
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
	return s
}

// SwampScreen is at (0,2) — mix of water and grass in irregular patterns.
// Open on top (faces LakeShore), right (faces SouthField).
// Walls on left and bottom (world boundary).
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
	return s
}

// SouthFieldScreen is at (1,2) — open field with sand paths.
// Open on top (faces Village), left (faces Swamp), right (faces Ruins).
// Wall on bottom (world boundary).
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
	return s
}

// RuinsScreen is at (2,2) — broken wall formations, floor tiles, stairs.
// Open on top (faces EastField), left (faces SouthField).
// Walls on right and bottom (world boundary).
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
	return s
}
