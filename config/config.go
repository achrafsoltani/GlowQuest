package config

const (
	TileSize    = 16
	ScreenGridW = 16
	ScreenGridH = 12
	HUDHeight   = 16

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
	OverworldW         = 3
	OverworldH         = 3
)
