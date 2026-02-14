package game

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/world"
)

type Transition struct {
	Active    bool
	Timer     float64
	Duration  float64
	DirX      int // -1, 0, +1
	DirY      int // -1, 0, +1
	OldScreen *world.Screen
}

func (t *Transition) Progress() float64 {
	if t.Duration <= 0 {
		return 1
	}
	p := t.Timer / t.Duration
	if p > 1 {
		return 1
	}
	return p
}

func (t *Transition) Done() bool {
	return t.Timer >= t.Duration
}

func (t *Transition) Start(dirX, dirY int, oldScreen *world.Screen) {
	t.Active = true
	t.Timer = 0
	t.Duration = config.TransitionDuration
	t.DirX = dirX
	t.DirY = dirY
	t.OldScreen = oldScreen
}

// easeInOut gives smooth acceleration/deceleration for screen transitions.
func easeInOut(t float64) float64 {
	if t < 0.5 {
		return 2 * t * t
	}
	return 1 - (-2*t+2)*(-2*t+2)/2
}
