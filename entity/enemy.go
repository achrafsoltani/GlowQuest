package entity

import "github.com/AchrafSoltani/GlowQuest/config"

type EnemyType int

const (
	EnemyOctorok  EnemyType = iota // 0
	EnemyMoblin                    // 1
	EnemyStalfos                   // 2
	EnemyBoss                      // 3

	// Extended types for future phases
	EnemyKeese     // 4
	EnemyGel       // 5
	EnemyZol       // 6
	EnemyBladeTrap // 7
	EnemySpark     // 8
	EnemyWizzrobe  // 9
	EnemyIronMask  // 10
	EnemyLikeLike  // 11
	EnemyGoomba    // 12 (side-scrolling)
	EnemyPiranha   // 13 (side-scrolling)
	EnemyZora      // 14
	EnemyArmos     // 15
	EnemyLanmola   // 16
)

type Enemy struct {
	Type            EnemyType
	X, Y            float64
	Width, Height   int
	Dir             Direction
	Speed           float64
	HP, MaxHP       int
	InvTimer        float64
	KnockbackX      float64
	KnockbackY      float64
	KnockbackTimer  float64
	AITimer         float64
	Dead            bool
	Moving          bool
	WalkFrame       int
	WalkTimer       float64
	ShootTimer      float64
	// Boss-specific
	AIState    int
	ChargeX    float64
	ChargeY    float64
	BurstCount int
}

func NewOctorok(x, y float64) *Enemy {
	return &Enemy{
		Type:   EnemyOctorok,
		X:      x,
		Y:      y,
		Width:  14,
		Height: 14,
		Dir:    DirDown,
		Speed:  30,
		HP:     2,
		MaxHP:  2,
		ShootTimer: 2.0,
	}
}

func NewMoblin(x, y float64) *Enemy {
	return &Enemy{
		Type:   EnemyMoblin,
		X:      x,
		Y:      y,
		Width:  14,
		Height: 14,
		Dir:    DirDown,
		Speed:  35,
		HP:     3,
		MaxHP:  3,
	}
}

func NewStalfos(x, y float64) *Enemy {
	return &Enemy{
		Type:   EnemyStalfos,
		X:      x,
		Y:      y,
		Width:  14,
		Height: 14,
		Dir:    DirDown,
		Speed:  45,
		HP:     2,
		MaxHP:  2,
	}
}

func NewBoss(x, y float64) *Enemy {
	return &Enemy{
		Type:       EnemyBoss,
		X:          x,
		Y:          y,
		Width:      config.BossSize,
		Height:     config.BossSize,
		Dir:        DirDown,
		Speed:      config.BossSpeed,
		HP:         config.BossHP,
		MaxHP:      config.BossHP,
		ShootTimer: 3.0,
	}
}

func (e *Enemy) CenterX() float64 { return e.X + float64(e.Width)/2 }
func (e *Enemy) CenterY() float64 { return e.Y + float64(e.Height)/2 }

func (e *Enemy) UpdateAnimation(dt float64) {
	if e.Moving {
		e.WalkTimer += dt
		if e.WalkTimer >= 0.15 {
			e.WalkTimer -= 0.15
			e.WalkFrame = (e.WalkFrame + 1) % 4
		}
	} else {
		e.WalkFrame = 0
		e.WalkTimer = 0
	}
}
