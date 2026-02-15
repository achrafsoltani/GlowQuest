package world

// DungeonID identifies a dungeon.
type DungeonID int

const (
	DungeonTailCave    DungeonID = 1
	DungeonBottleGrotto DungeonID = 2
	DungeonKeyCavern   DungeonID = 3
	DungeonAnglersTunnel DungeonID = 4
	DungeonCatfishsMaw DungeonID = 5
	DungeonFaceShrine  DungeonID = 6
	DungeonEaglesTower DungeonID = 7
	DungeonTurtleRock  DungeonID = 8
	DungeonColour      DungeonID = 9
)

// Dungeon represents a multi-room dungeon.
type Dungeon struct {
	ID          DungeonID
	Name        string
	GridW       int
	GridH       int
	Rooms       map[[2]int]*DungeonRoom
	EntrancePos [2]int
	BossRoomPos [2]int

	// Per-save dungeon state
	CurrentRoom [2]int
	SmallKeys   int
	HasMap      bool
	HasCompass  bool
	HasStoneBeak bool
	HasNightmareKey bool
	HasDungeonItem  bool
	VisitedRooms  map[[2]int]bool
	OpenedChests  map[string]bool
	ClearedRooms  map[[2]int]bool
	OpenedDoors   map[string]bool
}

// NewDungeon creates an empty dungeon structure.
func NewDungeon(id DungeonID, name string, w, h int) *Dungeon {
	return &Dungeon{
		ID:           id,
		Name:         name,
		GridW:        w,
		GridH:        h,
		Rooms:        make(map[[2]int]*DungeonRoom),
		VisitedRooms: make(map[[2]int]bool),
		OpenedChests: make(map[string]bool),
		ClearedRooms: make(map[[2]int]bool),
		OpenedDoors:  make(map[string]bool),
	}
}

// RoomAt returns the room at the given grid position, or nil.
func (d *Dungeon) RoomAt(x, y int) *DungeonRoom {
	return d.Rooms[[2]int{x, y}]
}

// CurrentDungeonRoom returns the room the player is currently in.
func (d *Dungeon) CurrentDungeonRoom() *DungeonRoom {
	return d.Rooms[d.CurrentRoom]
}
