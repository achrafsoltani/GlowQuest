package world

import "github.com/AchrafSoltani/GlowQuest/config"

// VillageHouseInterior1 — small room with an NPC inside (top-left house in village).
func VillageHouseInterior1() *InteriorDef {
	s := &Screen{}
	s.LoadFromString(
		"WWWWWWWWWWWWWWWW\n" +
			"WWWWWWWWWWWWWWWW\n" +
			"WWWWWWWWWWWWWWWW\n" +
			"WWWWWWWWWWWWWWWW\n" +
			"WWWWWWFFWWWWWWWW\n" +
			"WWWWWFFFFFFWWWWW\n" +
			"WWWWWFFFFFFWWWWW\n" +
			"WWWWWFFFFFFWWWWW\n" +
			"WWWWWFFFFFFWWWWW\n" +
			"WWWWWFFFFFFWWWWW\n" +
			"WWWWWWFOFWWWWWWW\n" +
			"WWWWWWWWWWWWWWWW")

	return &InteriorDef{
		ID:     "village_house_1",
		Screen: s,
		NPCSpawns: []NPCSpawn{
			{TileX: 8, TileY: 6, Dir: 0, Name: "Villager",
				Dialogue: []string{
					"Please make yourself",
					"at home. The village",
					"is peaceful... for now.",
				}},
		},
	}
}

// VillageHouseInterior2 — small room with different NPC (top-right house).
func VillageHouseInterior2() *InteriorDef {
	s := &Screen{}
	s.LoadFromString(
		"WWWWWWWWWWWWWWWW\n" +
			"WWWWWWWWWWWWWWWW\n" +
			"WWWWWWWWWWWWWWWW\n" +
			"WWWWWWWWWWWWWWWW\n" +
			"WWWWWWFFWWWWWWWW\n" +
			"WWWWWFFFFFFWWWWW\n" +
			"WWWWWFFFFFFWWWWW\n" +
			"WWWWWFFFFFFWWWWW\n" +
			"WWWWWFFFFFFWWWWW\n" +
			"WWWWWFFFFFFWWWWW\n" +
			"WWWWWWFOFWWWWWWW\n" +
			"WWWWWWWWWWWWWWWW")

	return &InteriorDef{
		ID:     "village_house_2",
		Screen: s,
		NPCSpawns: []NPCSpawn{
			{TileX: 8, TileY: 6, Dir: 0, Name: "Scholar",
				Dialogue: []string{
					"I've been studying the",
					"ruins. Ancient power",
					"sleeps beneath them.",
				}},
		},
	}
}

// ForestCaveInterior — narrow cave with enemies.
func ForestCaveInterior() *InteriorDef {
	s := &Screen{}
	s.LoadFromString(
		"WWWWWWWWWWWWWWWW\n" +
			"WWWWWWWWWWWWWWWW\n" +
			"WWWWWWWWWWWWWWWW\n" +
			"WWWWFFFFFFFFFFWW\n" +
			"WWWWF........FWW\n" +
			"WWWWF........FWW\n" +
			"WWWWF........FWW\n" +
			"WWWWF........FWW\n" +
			"WWWWFFFFFFFFFFWW\n" +
			"WWWWWWWWWWWWWWWW\n" +
			"WWWWWWWOFWWWWWWW\n" +
			"WWWWWWWWWWWWWWWW")

	// Replace dots with Floor in the cave
	for y := 4; y <= 7; y++ {
		for x := 5; x <= 12; x++ {
			s.Tiles[y][x] = TileFloor
		}
	}

	return &InteriorDef{
		ID:     "forest_cave",
		Screen: s,
		EnemySpawns: []EnemySpawn{
			{Type: 0, TileX: 7, TileY: 4},  // Octorok
			{Type: 0, TileX: 10, TileY: 6}, // Octorok
		},
		ItemSpawns: []ItemSpawn{
			{Type: 1, TileX: 6, TileY: 5}, // Rupee
		},
	}
}

// RuinsDungeonInterior — larger room with Stalfos and treasure.
func RuinsDungeonInterior() *InteriorDef {
	s := &Screen{}
	s.LoadFromString(
		"WWWWWWWWWWWWWWWW\n" +
			"WWWWWWWWWWWWWWWW\n" +
			"WWWFFFFFFFFFFFFFFW\n" +
			"WWWF..........FWW\n" +
			"WWWF..........FWW\n" +
			"WWWF..........FWW\n" +
			"WWWF..........FWW\n" +
			"WWWF..........FWW\n" +
			"WWWF..........FWW\n" +
			"WWWFFFFFFFFFFFFFFW\n" +
			"WWWWWWWOFWWWWWWWW\n" +
			"WWWWWWWWWWWWWWWW")

	// Manually set interior floor
	for y := 3; y <= 8; y++ {
		for x := 4; x <= 13; x++ {
			s.Tiles[y][x] = TileFloor
		}
	}
	// Walls around
	for x := 3; x <= 14; x++ {
		s.Tiles[2][x] = TileFloor
		s.Tiles[9][x] = TileFloor
	}
	for y := 2; y <= 9; y++ {
		s.Tiles[y][3] = TileFloor
		s.Tiles[y][14] = TileFloor
	}
	// Outer walls
	for x := 2; x <= 14; x++ {
		s.Tiles[1][x] = TileWall
		s.Tiles[10][x] = TileWall
	}
	for y := 1; y <= 10; y++ {
		s.Tiles[y][2] = TileWall
	}
	// Fix exit door
	s.Tiles[10][7] = TileDoorOpen
	s.Tiles[10][8] = TileFloor

	// Stairs deeper — leading to boss room
	s.Tiles[3][8] = TileStairs

	ts := float64(config.TileSize)
	return &InteriorDef{
		ID:     "ruins_dungeon",
		Screen: s,
		EnemySpawns: []EnemySpawn{
			{Type: 2, TileX: 6, TileY: 4},  // Stalfos
			{Type: 2, TileX: 11, TileY: 7}, // Stalfos
		},
		ItemSpawns: []ItemSpawn{
			{Type: 1, TileX: 9, TileY: 5},  // Rupee
			{Type: 1, TileX: 10, TileY: 5}, // Rupee
			{Type: 4, TileX: 9, TileY: 4},  // HeartContainer
		},
		DoorLinks: []DoorLink{
			{
				DoorTileX:  8,
				DoorTileY:  3,
				InteriorID: "boss_room",
				SpawnX:     7 * ts,
				SpawnY:     9 * ts,
				ExitX:      8 * ts,
				ExitY:      4 * ts,
			},
		},
	}
}

// BossRoomInterior — boss arena with wall pillars for cover.
func BossRoomInterior() *InteriorDef {
	s := &Screen{}
	// Start with all walls, carve out arena
	for y := 0; y < config.ScreenGridH; y++ {
		for x := 0; x < config.ScreenGridW; x++ {
			s.Tiles[y][x] = TileWall
		}
	}
	// Floor arena (rows 2-9, cols 2-13)
	for y := 2; y <= 9; y++ {
		for x := 2; x <= 13; x++ {
			s.Tiles[y][x] = TileFloor
		}
	}
	// 4 wall pillars for cover
	s.Tiles[4][5] = TileWall
	s.Tiles[4][10] = TileWall
	s.Tiles[7][5] = TileWall
	s.Tiles[7][10] = TileWall

	// Exit door at bottom
	s.Tiles[10][7] = TileDoorOpen

	return &InteriorDef{
		ID:     "boss_room",
		Screen: s,
		EnemySpawns: []EnemySpawn{
			{Type: 3, TileX: 7, TileY: 3}, // Boss
		},
	}
}

// BuildDoorLinks returns all door links in the overworld.
func BuildDoorLinks() []DoorLink {
	ts := float64(config.TileSize)
	return []DoorLink{
		// Village top-left house door (tile 2,4 is the 'O' door)
		{
			ScreenX: 1, ScreenY: 1,
			DoorTileX: 2, DoorTileY: 4,
			InteriorID: "village_house_1",
			SpawnX: 7 * ts, SpawnY: 9 * ts,
			ExitX: 2*ts + 1, ExitY: 4*ts + 2,
		},
		// Village top-right house door
		{
			ScreenX: 1, ScreenY: 1,
			DoorTileX: 12, DoorTileY: 4,
			InteriorID: "village_house_2",
			SpawnX: 7 * ts, SpawnY: 9 * ts,
			ExitX: 12*ts + 1, ExitY: 4*ts + 2,
		},
		// Village bottom-left house door
		{
			ScreenX: 1, ScreenY: 1,
			DoorTileX: 2, DoorTileY: 10,
			InteriorID: "village_house_1",
			SpawnX: 7 * ts, SpawnY: 9 * ts,
			ExitX: 2*ts + 1, ExitY: 10*ts + 2,
		},
		// Village bottom-right house door
		{
			ScreenX: 1, ScreenY: 1,
			DoorTileX: 12, DoorTileY: 10,
			InteriorID: "village_house_2",
			SpawnX: 7 * ts, SpawnY: 9 * ts,
			ExitX: 12*ts + 1, ExitY: 10*ts + 2,
		},
		// Ruins stairs → dungeon
		{
			ScreenX: 2, ScreenY: 2,
			DoorTileX: 7, DoorTileY: 6,
			InteriorID: "ruins_dungeon",
			SpawnX: 7 * ts, SpawnY: 8 * ts,
			ExitX: 7*ts + 1, ExitY: 6*ts + 2,
		},
	}
}

// BuildInteriors returns all interior definitions keyed by ID.
func BuildInteriors() map[string]*InteriorDef {
	return map[string]*InteriorDef{
		"village_house_1": VillageHouseInterior1(),
		"village_house_2": VillageHouseInterior2(),
		"forest_cave":     ForestCaveInterior(),
		"ruins_dungeon":   RuinsDungeonInterior(),
		"boss_room":       BossRoomInterior(),
	}
}
