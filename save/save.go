package save

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type SaveData struct {
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
	return &s
}

func Save(s *SaveData) error {
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
