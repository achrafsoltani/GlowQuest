package entity

// BossID identifies a specific boss.
type BossID int

const (
	BossNone      BossID = 0
	BossMoldorm   BossID = 1 // Dungeon 1
	BossGenie     BossID = 2 // Dungeon 2
	BossSlimeEye  BossID = 3 // Dungeon 3
	BossAnglerFish BossID = 4 // Dungeon 4
	BossSlimeEel  BossID = 5 // Dungeon 5
	BossFacade    BossID = 6 // Dungeon 6
	BossEvilEagle BossID = 7 // Dungeon 7
	BossHotHead   BossID = 8 // Dungeon 8
	BossShadow    BossID = 9 // Final boss
)

// Boss holds boss-specific state beyond the base Enemy struct.
type Boss struct {
	ID            BossID
	Phase         int     // current boss phase (0-indexed)
	MaxPhases     int     // total phases
	PhaseHP       []int   // HP threshold for each phase
	Vulnerable    bool    // currently vulnerable to damage
	PatternTimer  float64 // timer for current attack pattern
	PatternIndex  int     // which pattern in the sequence
}

// NewBossData creates boss-specific data for a given boss ID.
func NewBossData(id BossID) *Boss {
	switch id {
	case BossMoldorm:
		return &Boss{
			ID:         BossMoldorm,
			MaxPhases:  1,
			PhaseHP:    []int{8},
			Vulnerable: true,
		}
	default:
		return &Boss{
			ID:         id,
			MaxPhases:  1,
			PhaseHP:    []int{10},
			Vulnerable: true,
		}
	}
}
