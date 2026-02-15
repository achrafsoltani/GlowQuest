package world

type TileType int

const (
	TileGrass      TileType = 0  // '.' passable grass
	TileWall       TileType = 1  // 'W' solid wall
	TileWater      TileType = 2  // '~' deep water (needs Flippers)
	TileTree       TileType = 3  // 'T' tree (solid)
	TileSand       TileType = 4  // 'S' sand (passable)
	TileFloor      TileType = 5  // 'F' interior floor
	TileStairs     TileType = 6  // '>' stairs/warp
	TileDoorLocked TileType = 7  // 'D' locked door
	TileDoorOpen   TileType = 8  // 'O' open door

	// Terrain
	TileShallowWater TileType = 9  // shallow water (needs Flippers)
	TileCliffN       TileType = 10 // one-way ledge (jump down south)
	TileCliffS       TileType = 11 // one-way ledge (jump down north)
	TileCliffE       TileType = 12 // one-way ledge (jump down west)
	TileCliffW       TileType = 13 // one-way ledge (jump down east)
	TileBridge       TileType = 14 // bridge over water
	TilePit          TileType = 15 // pit (fall down, needs Roc's Feather to cross)

	// Interactive
	TileBush       TileType = 16 // cuttable bush
	TileRock       TileType = 17 // liftable rock (Power Bracelet L1)
	TileHeavyRock  TileType = 18 // heavy rock (Power Bracelet L2)
	TilePot        TileType = 19 // liftable pot
	TileSignpost   TileType = 20 // readable sign
	TileChest      TileType = 21 // unopened chest
	TileChestOpen  TileType = 22 // opened chest
	TileOwlStatue  TileType = 23 // owl statue (hints)
	TileKeyBlock   TileType = 24 // key block (uses small key)

	// Dungeon
	TileCrackedWall TileType = 25 // bombable cracked wall
	TileConveyorN   TileType = 26 // conveyor belt north
	TileConveyorS   TileType = 27 // conveyor belt south
	TileConveyorE   TileType = 28 // conveyor belt east
	TileConveyorW   TileType = 29 // conveyor belt west
	TileSpikes      TileType = 30 // floor spikes (damaging)
	TileLava        TileType = 31 // lava (damaging, impassable)
	TileIce         TileType = 32 // ice floor (slippery)
	TileSwitchOff   TileType = 33 // floor switch (unpressed)
	TileSwitchOn    TileType = 34 // floor switch (pressed)
	TileWarpTile    TileType = 35 // warp tile
	TileBossLocked  TileType = 36 // boss locked door
	TileBombable    TileType = 37 // bombable floor/wall
	TileTorch       TileType = 38 // unlit torch
	TileTorchLit    TileType = 39 // lit torch

	// Decoration
	TileGrassFlower TileType = 40 // grass with flowers
	TilePathH       TileType = 41 // horizontal path
	TilePathV       TileType = 42 // vertical path
	TileHouseFront  TileType = 43 // house front wall
	TileRoof        TileType = 44 // roof
	TileWindow      TileType = 45 // window (solid)
	TileFenceH      TileType = 46 // horizontal fence
	TileFenceV      TileType = 47 // vertical fence

	TileCount TileType = 48
)

func TileFromChar(c byte) TileType {
	switch c {
	case '.':
		return TileGrass
	case 'W':
		return TileWall
	case '~':
		return TileWater
	case 'T':
		return TileTree
	case 'S':
		return TileSand
	case 'F':
		return TileFloor
	case '>':
		return TileStairs
	case 'D':
		return TileDoorLocked
	case 'O':
		return TileDoorOpen
	default:
		return TileGrass
	}
}
