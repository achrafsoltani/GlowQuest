package entity

import "github.com/AchrafSoltani/GlowQuest/config"

type Projectile struct {
	X, Y          float64
	DirX, DirY    float64
	Speed         float64
	Damage        int
	FromEnemy     bool
	Width, Height int
	Dead          bool
}

func NewEnemyProjectile(x, y, dirX, dirY float64) *Projectile {
	return &Projectile{
		X:         x,
		Y:         y,
		DirX:      dirX,
		DirY:      dirY,
		Speed:     config.ProjectileSpeed,
		Damage:    1,
		FromEnemy: true,
		Width:     4,
		Height:    4,
	}
}

func (p *Projectile) Update(dt float64) {
	p.X += p.DirX * p.Speed * dt
	p.Y += p.DirY * p.Speed * dt

	// Kill if out of bounds
	if p.X < -10 || p.X > float64(config.PlayAreaWidth)+10 ||
		p.Y < -10 || p.Y > float64(config.PlayAreaHeight)+10 {
		p.Dead = true
	}
}
