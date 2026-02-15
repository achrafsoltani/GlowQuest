package game

// QuestState tracks all persistent quest progress.
type QuestState struct {
	Flags              map[string]bool
	DungeonsCompleted  [9]bool
	TradingItem        int
	SeashellsCollected map[string]bool
	HeartPieces        map[string]bool
	WarpPoints         map[string]bool
}

// NewQuestState creates a fresh quest state.
func NewQuestState() *QuestState {
	return &QuestState{
		Flags:              make(map[string]bool),
		SeashellsCollected: make(map[string]bool),
		HeartPieces:        make(map[string]bool),
		WarpPoints:         make(map[string]bool),
	}
}

// SetFlag sets a quest flag.
func (q *QuestState) SetFlag(key string) {
	q.Flags[key] = true
}

// HasFlag checks if a quest flag is set.
func (q *QuestState) HasFlag(key string) bool {
	return q.Flags[key]
}

// IsDungeonComplete returns whether a dungeon has been completed.
func (q *QuestState) IsDungeonComplete(dungeonNum int) bool {
	if dungeonNum < 1 || dungeonNum > 9 {
		return false
	}
	return q.DungeonsCompleted[dungeonNum-1]
}

// CompleteDungeon marks a dungeon as completed.
func (q *QuestState) CompleteDungeon(dungeonNum int) {
	if dungeonNum >= 1 && dungeonNum <= 9 {
		q.DungeonsCompleted[dungeonNum-1] = true
	}
}
