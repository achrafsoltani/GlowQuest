package system

import (
	"math"

	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
	"github.com/AchrafSoltani/GlowQuest/world"
)

// UpdateEnemyAI updates an enemy's AI behaviour and movement.
// Returns a projectile if the enemy fires one, otherwise nil.
func UpdateEnemyAI(e *entity.Enemy, p *entity.Player, screen *world.Screen, dt float64, rng *SimpleRNG) *entity.Projectile {
	if e.Dead {
		return nil
	}

	// Update invincibility
	if e.InvTimer > 0 {
		e.InvTimer -= dt
	}

	// Handle knockback — skip AI movement
	if e.KnockbackTimer > 0 {
		UpdateKnockback(e, dt)
		return nil
	}

	e.AITimer -= dt

	switch e.Type {
	case entity.EnemyOctorok:
		return updateOctorok(e, p, screen, dt, rng)
	case entity.EnemyMoblin:
		updateMoblin(e, p, screen, dt, rng)
	case entity.EnemyStalfos:
		updateStalfos(e, p, screen, dt, rng)
	case entity.EnemyBoss:
		return updateBoss(e, p, screen, dt, rng)
	}
	return nil
}

func updateOctorok(e *entity.Enemy, p *entity.Player, screen *world.Screen, dt float64, rng *SimpleRNG) *entity.Projectile {
	// Wander randomly
	if e.AITimer <= 0 {
		e.AITimer = 1.0 + float64(rng.Next()%200)/100.0
		switch rng.Next() % 4 {
		case 0:
			e.Dir = entity.DirUp
		case 1:
			e.Dir = entity.DirDown
		case 2:
			e.Dir = entity.DirLeft
		case 3:
			e.Dir = entity.DirRight
		}
		e.Moving = true
	}

	moveEnemy(e, screen, dt)

	// Shoot projectile
	e.ShootTimer -= dt
	if e.ShootTimer <= 0 {
		e.ShootTimer = 2.0 + float64(rng.Next()%100)/100.0
		return fireAtPlayer(e, p)
	}
	return nil
}

func updateMoblin(e *entity.Enemy, p *entity.Player, screen *world.Screen, dt float64, rng *SimpleRNG) {
	dist := distBetween(e.CenterX(), e.CenterY(), p.CenterX(), p.CenterY())

	if dist < 80 {
		// Chase player
		chasePlayer(e, p)
		e.Moving = true
	} else {
		// Wander
		if e.AITimer <= 0 {
			e.AITimer = 1.0 + float64(rng.Next()%200)/100.0
			switch rng.Next() % 5 {
			case 0:
				e.Dir = entity.DirUp
				e.Moving = true
			case 1:
				e.Dir = entity.DirDown
				e.Moving = true
			case 2:
				e.Dir = entity.DirLeft
				e.Moving = true
			case 3:
				e.Dir = entity.DirRight
				e.Moving = true
			case 4:
				e.Moving = false
			}
		}
	}

	moveEnemy(e, screen, dt)
}

func updateStalfos(e *entity.Enemy, p *entity.Player, screen *world.Screen, dt float64, rng *SimpleRNG) {
	dist := distBetween(e.CenterX(), e.CenterY(), p.CenterX(), p.CenterY())

	if dist < 48 {
		// Chase player
		chasePlayer(e, p)
		e.Moving = true
	} else {
		// Fast wander
		if e.AITimer <= 0 {
			e.AITimer = 0.5 + float64(rng.Next()%150)/100.0
			switch rng.Next() % 4 {
			case 0:
				e.Dir = entity.DirUp
			case 1:
				e.Dir = entity.DirDown
			case 2:
				e.Dir = entity.DirLeft
			case 3:
				e.Dir = entity.DirRight
			}
			e.Moving = true
		}
	}

	moveEnemy(e, screen, dt)
}

func chasePlayer(e *entity.Enemy, p *entity.Player) {
	dx := p.CenterX() - e.CenterX()
	dy := p.CenterY() - e.CenterY()

	if math.Abs(dx) > math.Abs(dy) {
		if dx > 0 {
			e.Dir = entity.DirRight
		} else {
			e.Dir = entity.DirLeft
		}
	} else {
		if dy > 0 {
			e.Dir = entity.DirDown
		} else {
			e.Dir = entity.DirUp
		}
	}
}

func moveEnemy(e *entity.Enemy, screen *world.Screen, dt float64) {
	if !e.Moving {
		e.UpdateAnimation(dt)
		return
	}

	dx := e.Dir.DX()
	dy := e.Dir.DY()
	dist := e.Speed * dt

	// Try X
	if dx != 0 {
		newX := e.X + dx*dist
		if newX >= 0 && newX+float64(e.Width) <= float64(config.PlayAreaWidth) &&
			!TileCollision(screen, newX, e.Y, e.Width, e.Height) {
			e.X = newX
		}
	}

	// Try Y
	if dy != 0 {
		newY := e.Y + dy*dist
		if newY >= 0 && newY+float64(e.Height) <= float64(config.PlayAreaHeight) &&
			!TileCollision(screen, e.X, newY, e.Width, e.Height) {
			e.Y = newY
		}
	}

	e.UpdateAnimation(dt)
}

func fireAtPlayer(e *entity.Enemy, p *entity.Player) *entity.Projectile {
	dx := p.CenterX() - e.CenterX()
	dy := p.CenterY() - e.CenterY()
	dist := math.Sqrt(dx*dx + dy*dy)
	if dist < 0.01 {
		return nil
	}
	dx /= dist
	dy /= dist
	return entity.NewEnemyProjectile(e.CenterX(), e.CenterY(), dx, dy)
}

func updateBoss(e *entity.Enemy, p *entity.Player, screen *world.Screen, dt float64, rng *SimpleRNG) *entity.Projectile {
	switch e.AIState {
	case 0: // Wander / idle
		if e.AITimer <= 0 {
			e.AITimer = 3.0 + float64(rng.Next()%100)/100.0

			dist := distBetween(e.CenterX(), e.CenterY(), p.CenterX(), p.CenterY())
			if dist < 80 && rng.Next()%2 == 0 {
				// Start charge
				e.AIState = 1
				e.AITimer = 1.0
				dx := p.CenterX() - e.CenterX()
				dy := p.CenterY() - e.CenterY()
				d := math.Sqrt(dx*dx + dy*dy)
				if d > 0.01 {
					e.ChargeX = dx / d
					e.ChargeY = dy / d
				}
				return nil
			}
			// Start burst fire
			if dist < 120 {
				e.AIState = 2
				e.AITimer = 0.3
				e.BurstCount = 0
				return nil
			}

			// Otherwise wander
			switch rng.Next() % 4 {
			case 0:
				e.Dir = entity.DirUp
			case 1:
				e.Dir = entity.DirDown
			case 2:
				e.Dir = entity.DirLeft
			case 3:
				e.Dir = entity.DirRight
			}
			e.Moving = true
		}
		moveEnemy(e, screen, dt)

	case 1: // Charge
		e.AITimer -= dt
		if e.AITimer <= 0 {
			e.AIState = 0
			e.AITimer = 1.0
			e.Moving = false
			return nil
		}
		chargeSpeed := 120.0
		newX := e.X + e.ChargeX*chargeSpeed*dt
		newY := e.Y + e.ChargeY*chargeSpeed*dt
		if newX >= 0 && newX+float64(e.Width) <= float64(config.PlayAreaWidth) &&
			!TileCollision(screen, newX, e.Y, e.Width, e.Height) {
			e.X = newX
		}
		if newY >= 0 && newY+float64(e.Height) <= float64(config.PlayAreaHeight) &&
			!TileCollision(screen, e.X, newY, e.Width, e.Height) {
			e.Y = newY
		}
		e.Moving = true
		e.UpdateAnimation(dt)

	case 2: // Burst fire
		e.AITimer -= dt
		e.Moving = false
		if e.AITimer <= 0 && e.BurstCount < 4 {
			e.BurstCount++
			e.AITimer = 0.3
			e.UpdateAnimation(dt)
			return fireAtPlayer(e, p)
		}
		if e.BurstCount >= 4 {
			e.AIState = 0
			e.AITimer = 2.0
		}
		e.UpdateAnimation(dt)
	}
	return nil
}

func distBetween(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

// SimpleRNG is a basic pseudo-random number generator (xorshift32).
type SimpleRNG struct {
	state uint32
}

func NewSimpleRNG(seed uint32) *SimpleRNG {
	if seed == 0 {
		seed = 12345
	}
	return &SimpleRNG{state: seed}
}

func (r *SimpleRNG) Next() uint32 {
	r.state ^= r.state << 13
	r.state ^= r.state >> 17
	r.state ^= r.state << 5
	return r.state
}
