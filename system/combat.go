package system

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
)

// CheckSwordHits returns enemies hit by the player's active sword swing.
func CheckSwordHits(p *entity.Player, enemies []*entity.Enemy) []*entity.Enemy {
	if !p.Sword.Active {
		return nil
	}

	sx, sy, sw, sh := p.Sword.HitBox(p.X, p.Y, p.Width, p.Height)
	var hit []*entity.Enemy

	for _, e := range enemies {
		if e.Dead || e.InvTimer > 0 {
			continue
		}
		if AABBOverlap(sx, sy, sw, sh, e.X, e.Y, float64(e.Width), float64(e.Height)) {
			hit = append(hit, e)
		}
	}
	return hit
}

// ApplyKnockback starts a knockback effect on an enemy, pushing away from the player.
func ApplyKnockback(e *entity.Enemy, fromX, fromY float64) {
	dx := e.CenterX() - fromX
	dy := e.CenterY() - fromY
	dist := dx*dx + dy*dy
	if dist < 0.01 {
		dx = 0
		dy = 1
	}
	// Normalise
	if dist > 0 {
		inv := 1.0 / sqrt(dist)
		dx *= inv
		dy *= inv
	}
	e.KnockbackX = dx * config.KnockbackDist / config.KnockbackTime
	e.KnockbackY = dy * config.KnockbackDist / config.KnockbackTime
	e.KnockbackTimer = config.KnockbackTime
}

// UpdateKnockback moves an enemy during knockback.
func UpdateKnockback(e *entity.Enemy, dt float64) {
	if e.KnockbackTimer <= 0 {
		return
	}
	e.X += e.KnockbackX * dt
	e.Y += e.KnockbackY * dt
	e.KnockbackTimer -= dt
	if e.KnockbackTimer <= 0 {
		e.KnockbackTimer = 0
		e.KnockbackX = 0
		e.KnockbackY = 0
	}
	// Clamp to play area
	if e.X < 0 {
		e.X = 0
	}
	if e.Y < 0 {
		e.Y = 0
	}
	maxX := float64(config.PlayAreaWidth - e.Width)
	maxY := float64(config.PlayAreaHeight - e.Height)
	if e.X > maxX {
		e.X = maxX
	}
	if e.Y > maxY {
		e.Y = maxY
	}
}

// CheckEnemyPlayerCollision checks if an enemy overlaps the player.
func CheckEnemyPlayerCollision(p *entity.Player, e *entity.Enemy) bool {
	return AABBOverlap(p.X, p.Y, float64(p.Width), float64(p.Height),
		e.X, e.Y, float64(e.Width), float64(e.Height))
}

// CheckProjectilePlayerCollision checks if a projectile overlaps the player.
func CheckProjectilePlayerCollision(p *entity.Player, proj *entity.Projectile) bool {
	return AABBOverlap(p.X, p.Y, float64(p.Width), float64(p.Height),
		proj.X, proj.Y, float64(proj.Width), float64(proj.Height))
}

// CheckProjectileSwordCollision checks if a projectile overlaps the sword hitbox.
func CheckProjectileSwordCollision(p *entity.Player, proj *entity.Projectile) bool {
	if !p.Sword.Active {
		return false
	}
	sx, sy, sw, sh := p.Sword.HitBox(p.X, p.Y, p.Width, p.Height)
	return AABBOverlap(sx, sy, sw, sh, proj.X, proj.Y, float64(proj.Width), float64(proj.Height))
}

func sqrt(x float64) float64 {
	if x <= 0 {
		return 0
	}
	// Newton's method
	z := x / 2
	for i := 0; i < 10; i++ {
		z = (z + x/z) / 2
	}
	return z
}
