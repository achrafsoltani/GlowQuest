package system

import (
	"github.com/AchrafSoltani/GlowQuest/entity"
	"github.com/AchrafSoltani/GlowQuest/world"
)

// UpdateBossAI dispatches to boss-specific AI based on the BossID.
// Returns a projectile if the boss fires one.
// This is a framework — individual boss AIs will be added in dungeon phases.
func UpdateBossAI(boss *entity.Boss, e *entity.Enemy, p *entity.Player, screen *world.Screen, dt float64, rng *SimpleRNG) *entity.Projectile {
	if boss == nil {
		return nil
	}

	switch boss.ID {
	case entity.BossMoldorm:
		return updateMoldormAI(boss, e, p, screen, dt, rng)
	default:
		// Default: use the existing boss AI from enemy_ai.go
		return nil
	}
}

// updateMoldormAI is a placeholder for Phase 8 Moldorm boss.
func updateMoldormAI(boss *entity.Boss, e *entity.Enemy, p *entity.Player, screen *world.Screen, dt float64, rng *SimpleRNG) *entity.Projectile {
	// Placeholder — will be fully implemented in Phase 8
	return nil
}
