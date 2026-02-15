package game

type GameState int

const (
	StateMenu     GameState = iota
	StatePlaying
	StatePaused
	StateGameOver
	StateDialogue
	StateVictory
	StateInventory
)
