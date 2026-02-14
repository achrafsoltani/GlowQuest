package game

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/world"
)

type TransitionType int

const (
	TransitionScroll TransitionType = iota
	TransitionFade
)

type Transition struct {
	Active    bool
	Type      TransitionType
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
	t.Type = TransitionScroll
	t.Timer = 0
	t.Duration = config.TransitionDuration
	t.DirX = dirX
	t.DirY = dirY
	t.OldScreen = oldScreen
}

func (t *Transition) StartFade() {
	t.Active = true
	t.Type = TransitionFade
	t.Timer = 0
	t.Duration = config.FadeDuration
	t.DirX = 0
	t.DirY = 0
	t.OldScreen = nil
}

// FadeProgress returns 0→1→0 for the fade (dim to black, then brighten).
func (t *Transition) FadeProgress() float64 {
	p := t.Progress()
	if p < 0.5 {
		return p * 2 // 0→1 (dim to black)
	}
	return (1 - p) * 2 // 1→0 (brighten)
}

// easeInOut gives smooth acceleration/deceleration for screen transitions.
func easeInOut(t float64) float64 {
	if t < 0.5 {
		return 2 * t * t
	}
	return 1 - (-2*t+2)*(-2*t+2)/2
}
