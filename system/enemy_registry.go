package system

import "github.com/AchrafSoltani/GlowQuest/entity"

// AIType identifies the AI behaviour pattern for an enemy.
type AIType int

const (
	AIWander     AIType = iota // random movement
	AIChase                    // chase player when close
	AIShooter                  // wander + shoot projectiles
	AIBladeTrap                // stationary, charge when aligned
	AISpark                    // move along walls
	AIBounce                   // bounce off walls
	AIStationary               // doesn't move
)

// EnemyDef is a data-driven enemy definition.
type EnemyDef struct {
	Type      entity.EnemyType
	Name      string
	Width     int
	Height    int
	HP        int
	Speed     float64
	AI        AIType
	ChaseRange float64 // distance to start chasing (0 = always)
	ShootRate  float64 // seconds between shots (0 = no shooting)
	ContactDmg int    // damage on contact (default 1)
}

// EnemyRegistry holds definitions for all enemy types.
var EnemyRegistry = map[entity.EnemyType]*EnemyDef{
	entity.EnemyOctorok: {
		Type: entity.EnemyOctorok, Name: "Octorok",
		Width: 14, Height: 14, HP: 2, Speed: 30,
		AI: AIShooter, ShootRate: 2.0, ContactDmg: 1,
	},
	entity.EnemyMoblin: {
		Type: entity.EnemyMoblin, Name: "Moblin",
		Width: 14, Height: 14, HP: 3, Speed: 35,
		AI: AIChase, ChaseRange: 80, ContactDmg: 1,
	},
	entity.EnemyStalfos: {
		Type: entity.EnemyStalfos, Name: "Stalfos",
		Width: 14, Height: 14, HP: 2, Speed: 45,
		AI: AIChase, ChaseRange: 48, ContactDmg: 1,
	},
	entity.EnemyBoss: {
		Type: entity.EnemyBoss, Name: "Boss",
		Width: 20, Height: 20, HP: 10, Speed: 25,
		AI: AIWander, ContactDmg: 2,
	},
	// Extended enemy types for future phases
	entity.EnemyKeese: {
		Type: entity.EnemyKeese, Name: "Keese",
		Width: 12, Height: 12, HP: 1, Speed: 50,
		AI: AIBounce, ContactDmg: 1,
	},
	entity.EnemyGel: {
		Type: entity.EnemyGel, Name: "Gel",
		Width: 10, Height: 10, HP: 1, Speed: 20,
		AI: AIChase, ChaseRange: 40, ContactDmg: 1,
	},
	entity.EnemyZol: {
		Type: entity.EnemyZol, Name: "Zol",
		Width: 14, Height: 14, HP: 2, Speed: 15,
		AI: AIChase, ChaseRange: 48, ContactDmg: 1,
	},
	entity.EnemyBladeTrap: {
		Type: entity.EnemyBladeTrap, Name: "Blade Trap",
		Width: 16, Height: 16, HP: 99, Speed: 120,
		AI: AIBladeTrap, ContactDmg: 2,
	},
	entity.EnemySpark: {
		Type: entity.EnemySpark, Name: "Spark",
		Width: 12, Height: 12, HP: 99, Speed: 30,
		AI: AISpark, ContactDmg: 1,
	},
}

// GetEnemyDef returns the definition for an enemy type, or nil if unknown.
func GetEnemyDef(t entity.EnemyType) *EnemyDef {
	return EnemyRegistry[t]
}
