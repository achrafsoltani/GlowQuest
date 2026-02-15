package world

// TileProperties defines behaviour for each tile type.
type TileProperties struct {
	Passable    bool
	Swimmable   bool    // needs Flippers to traverse
	Cuttable    bool    // sword destroys → becomes grass
	Liftable    int     // 0=no, 1=Power Bracelet L1, 2=Power Bracelet L2
	Bombable    bool    // can be destroyed by bombs
	Damaging    bool    // hurts the player on contact
	SlowFactor  float64 // 1.0=normal, <1.0=slower
	Slippery    bool    // ice physics
	ConveyorDir int     // 0=none, 1=N, 2=S, 3=E, 4=W
	JumpDown    bool    // one-way ledge
	JumpDir     int     // direction to jump (0=N,1=S,2=E,3=W)
}

// TileProps is the global tile property lookup table.
var TileProps [64]TileProperties

func init() {
	// Default all to impassable
	for i := range TileProps {
		TileProps[i] = TileProperties{SlowFactor: 1.0}
	}

	// Passable ground tiles
	TileProps[TileGrass] = TileProperties{Passable: true, SlowFactor: 1.0}
	TileProps[TileSand] = TileProperties{Passable: true, SlowFactor: 1.0}
	TileProps[TileFloor] = TileProperties{Passable: true, SlowFactor: 1.0}
	TileProps[TileStairs] = TileProperties{Passable: true, SlowFactor: 1.0}
	TileProps[TileDoorOpen] = TileProperties{Passable: true, SlowFactor: 1.0}
	TileProps[TileBridge] = TileProperties{Passable: true, SlowFactor: 1.0}
	TileProps[TileGrassFlower] = TileProperties{Passable: true, SlowFactor: 1.0}
	TileProps[TilePathH] = TileProperties{Passable: true, SlowFactor: 1.0}
	TileProps[TilePathV] = TileProperties{Passable: true, SlowFactor: 1.0}
	TileProps[TileChestOpen] = TileProperties{Passable: true, SlowFactor: 1.0}
	TileProps[TileWarpTile] = TileProperties{Passable: true, SlowFactor: 1.0}
	TileProps[TileSwitchOff] = TileProperties{Passable: true, SlowFactor: 1.0}
	TileProps[TileSwitchOn] = TileProperties{Passable: true, SlowFactor: 1.0}
	TileProps[TileTorchLit] = TileProperties{Passable: false, SlowFactor: 1.0}
	TileProps[TileTorch] = TileProperties{Passable: false, SlowFactor: 1.0}

	// Impassable solid tiles (default, already set)
	// TileWall, TileTree, TileDoorLocked, TileHouseFront, TileRoof, TileWindow, TileFenceH, TileFenceV

	// Water tiles — need Flippers
	TileProps[TileWater] = TileProperties{Swimmable: true, SlowFactor: 0.5}
	TileProps[TileShallowWater] = TileProperties{Swimmable: true, SlowFactor: 0.6}

	// Interactive tiles
	TileProps[TileBush] = TileProperties{Cuttable: true, SlowFactor: 1.0}
	TileProps[TileRock] = TileProperties{Liftable: 1, SlowFactor: 1.0}
	TileProps[TileHeavyRock] = TileProperties{Liftable: 2, SlowFactor: 1.0}
	TileProps[TilePot] = TileProperties{Liftable: 1, SlowFactor: 1.0}
	TileProps[TileSignpost] = TileProperties{SlowFactor: 1.0} // impassable, interactable
	TileProps[TileChest] = TileProperties{SlowFactor: 1.0}    // impassable until opened
	TileProps[TileOwlStatue] = TileProperties{SlowFactor: 1.0}
	TileProps[TileKeyBlock] = TileProperties{SlowFactor: 1.0}

	// Dungeon tiles
	TileProps[TileCrackedWall] = TileProperties{Bombable: true, SlowFactor: 1.0}
	TileProps[TileBombable] = TileProperties{Bombable: true, SlowFactor: 1.0}
	TileProps[TileConveyorN] = TileProperties{Passable: true, ConveyorDir: 1, SlowFactor: 1.0}
	TileProps[TileConveyorS] = TileProperties{Passable: true, ConveyorDir: 2, SlowFactor: 1.0}
	TileProps[TileConveyorE] = TileProperties{Passable: true, ConveyorDir: 3, SlowFactor: 1.0}
	TileProps[TileConveyorW] = TileProperties{Passable: true, ConveyorDir: 4, SlowFactor: 1.0}
	TileProps[TileSpikes] = TileProperties{Passable: true, Damaging: true, SlowFactor: 1.0}
	TileProps[TileLava] = TileProperties{Damaging: true, SlowFactor: 1.0} // impassable + damaging
	TileProps[TileIce] = TileProperties{Passable: true, Slippery: true, SlowFactor: 1.0}
	TileProps[TileBossLocked] = TileProperties{SlowFactor: 1.0} // impassable until nightmare key
	TileProps[TilePit] = TileProperties{SlowFactor: 1.0}        // impassable (jump over with Roc's Feather)

	// Cliff ledges — passable only from one direction (jump down)
	TileProps[TileCliffN] = TileProperties{JumpDown: true, JumpDir: 0, SlowFactor: 1.0}
	TileProps[TileCliffS] = TileProperties{JumpDown: true, JumpDir: 1, SlowFactor: 1.0}
	TileProps[TileCliffE] = TileProperties{JumpDown: true, JumpDir: 2, SlowFactor: 1.0}
	TileProps[TileCliffW] = TileProperties{JumpDown: true, JumpDir: 3, SlowFactor: 1.0}
}
