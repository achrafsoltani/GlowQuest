package save

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// QuestSaveData holds serializable quest state.
type QuestSaveData struct {
	Flags             map[string]bool `json:"flags,omitempty"`
	DungeonsCompleted [9]bool         `json:"dungeons_completed"`
	TradingItem       int             `json:"trading_item"`
}

type SaveData struct {
	Version        int             `json:"version"`
	HasSword       bool            `json:"has_sword"`
	MaxHP          int             `json:"max_hp"`
	HP             int             `json:"hp"`
	Rupees         int             `json:"rupees"`
	Keys           int             `json:"keys"`
	CollectedItems map[string]bool `json:"collected_items"`
	UnlockedDoors  map[string]bool `json:"unlocked_doors"`
	ScreenX        int             `json:"screen_x"`
	ScreenY        int             `json:"screen_y"`
	PlayerX        float64         `json:"player_x"`
	PlayerY        float64         `json:"player_y"`
	InInterior     bool            `json:"in_interior"`
	InteriorID     string          `json:"interior_id,omitempty"`
	BossDefeated   bool            `json:"boss_defeated"`

	// V2 fields
	Bombs         int             `json:"bombs,omitempty"`
	Arrows        int             `json:"arrows,omitempty"`
	SwordLevel    int             `json:"sword_level,omitempty"`
	ShieldLevel   int             `json:"shield_level,omitempty"`
	BraceletLevel int             `json:"bracelet_level,omitempty"`
	ButtonA       int             `json:"button_a,omitempty"`
	ButtonB       int             `json:"button_b,omitempty"`
	OwnedItems    []int           `json:"owned_items,omitempty"`
	LocationType  int             `json:"location_type,omitempty"`
	DungeonID     string          `json:"dungeon_id,omitempty"`
	DungeonRoomX  int             `json:"dungeon_room_x,omitempty"`
	DungeonRoomY  int             `json:"dungeon_room_y,omitempty"`
	Quest         *QuestSaveData  `json:"quest,omitempty"`
}

func savePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "glowquest", "save.json")
}

func Exists() bool {
	_, err := os.Stat(savePath())
	return err == nil
}

func Load() *SaveData {
	data, err := os.ReadFile(savePath())
	if err != nil {
		return nil
	}
	var s SaveData
	if err := json.Unmarshal(data, &s); err != nil {
		return nil
	}
	// Migrate V1 saves
	if s.Version == 0 {
		s.Version = 1
	}
	if s.Version == 1 {
		migrateV1(&s)
	}
	return &s
}

func Save(s *SaveData) error {
	if s.Version == 0 {
		s.Version = 2
	}
	p := savePath()
	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p, data, 0644)
}

// migrateV1 upgrades a V1 save to V2 format.
func migrateV1(s *SaveData) {
	s.Version = 2
	// If player had sword in V1, set it as owned and assign to A button
	if s.HasSword {
		s.SwordLevel = 1
		s.OwnedItems = append(s.OwnedItems, 1) // EquipSword = 1
		if s.ButtonA == 0 {
			s.ButtonA = 1 // EquipSword
		}
	}
}
