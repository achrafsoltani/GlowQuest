package world

// DoorType defines how a dungeon room door behaves.
type DoorType int

const (
	DoorNone     DoorType = iota // no door (wall)
	DoorOpen                     // always open
	DoorLocked                   // requires small key
	DoorBoss                     // requires nightmare key
	DoorBombable                 // requires bomb to open
	DoorOneWay                   // one-way passage
	DoorKeyBlock                 // key block
	DoorEntrance                 // dungeon entrance/exit
)

// ClearCondition defines what clears a dungeon room.
type ClearCondition int

const (
	ClearNone    ClearCondition = iota // no clear needed
	ClearKillAll                       // kill all enemies
	ClearPuzzle                        // solve a puzzle
	ClearBoss                          // defeat the boss
)

// RoomDoor represents a door on one side of a room.
type RoomDoor struct {
	Type   DoorType
	Opened bool
}

// DungeonChest represents a chest in a dungeon room.
type DungeonChest struct {
	X        int
	Y        int
	Contents string // item ID
	Flag     string // save flag for persistence
	Opened   bool
}

// DungeonRoom represents a single room in a dungeon.
type DungeonRoom struct {
	X, Y           int
	Screen         *Screen
	Doors          [4]*RoomDoor // 0=North, 1=South, 2=East, 3=West
	EnemySpawns    []EnemySpawn
	Chests         []DungeonChest
	ClearCondition ClearCondition
	Cleared        bool
}

// NewDungeonRoom creates an empty dungeon room.
func NewDungeonRoom(x, y int) *DungeonRoom {
	return &DungeonRoom{
		X:      x,
		Y:      y,
		Screen: &Screen{},
	}
}
