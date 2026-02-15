package config

const (
	TileSize    = 16
	ScreenGridW = 16
	ScreenGridH = 12
	HUDHeight   = 32

	PlayAreaWidth  = ScreenGridW * TileSize // 256
	PlayAreaHeight = ScreenGridH * TileSize // 192
	WindowWidth    = PlayAreaWidth          // 256
	WindowHeight   = PlayAreaHeight + HUDHeight // 224

	PlayerSize  = 14
	PlayerSpeed = 80.0
	MaxHP       = 6 // 3 full hearts

	WalkFrameTime = 0.12 // seconds per walk animation frame
	WalkFrames    = 4

	TransitionDuration = 0.5 // seconds for screen scroll transition

	// Combat
	SwordDuration   = 0.2
	SwordReach      = 12
	SwordWidth      = 10
	KnockbackDist   = 32.0
	KnockbackTime   = 0.15
	PlayerInvTime   = 1.0
	EnemyInvTime    = 0.5
	ProjectileSpeed = 100.0

	// Items
	ItemBobSpeed  = 4.0
	ItemBobAmount = 1

	// Dialogue
	DialogueBoxH  = 48
	InteractRadius = 20.0

	// Interior transitions
	FadeDuration = 0.6

	// Boss
	BossHP    = 10
	BossSize  = 20
	BossSpeed = 25.0

	// Polish
	ShakeDuration  = 0.2
	ShakeIntensity = 3
	FlashDuration  = 0.08

	// Menu
	GameOverDelay = 2.0
	VictoryDelay  = 3.0
)

// Overworld dimensions — set dynamically by the map loader.
var (
	OverworldW = 16
	OverworldH = 16
)
