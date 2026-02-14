package system

import "github.com/AchrafSoltani/glow"

type InputTracker struct {
	held        map[glow.Key]bool
	justPressed map[glow.Key]bool
}

func NewInputTracker() *InputTracker {
	return &InputTracker{
		held:        make(map[glow.Key]bool),
		justPressed: make(map[glow.Key]bool),
	}
}

func (it *InputTracker) KeyDown(key glow.Key) {
	if !it.held[key] {
		it.justPressed[key] = true
	}
	it.held[key] = true
}

func (it *InputTracker) KeyUp(key glow.Key) {
	it.held[key] = false
	it.justPressed[key] = false
}

func (it *InputTracker) IsHeld(key glow.Key) bool {
	return it.held[key]
}

func (it *InputTracker) JustPressed(key glow.Key) bool {
	return it.justPressed[key]
}

func (it *InputTracker) Update() {
	for k := range it.justPressed {
		delete(it.justPressed, k)
	}
}
