package world

type DoorLink struct {
	ScreenX    int
	ScreenY    int
	DoorTileX  int
	DoorTileY  int
	InteriorID string
	SpawnX     float64
	SpawnY     float64
	ExitX      float64
	ExitY      float64
}

type InteriorDef struct {
	ID          string
	Screen      *Screen
	EnemySpawns []EnemySpawn
	ItemSpawns  []ItemSpawn
	NPCSpawns   []NPCSpawn
	DoorLinks   []DoorLink
}
