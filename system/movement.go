package system

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
	"github.com/AchrafSoltani/GlowQuest/world"
)

// MovePlayer moves the player and returns an edge-crossing signal.
// Returns (0,0) normally. Returns (-1,0)/(1,0)/(0,-1)/(0,1) if the
// player walked off the left/right/top/bottom edge respectively.
func MovePlayer(p *entity.Player, screen *world.Screen, dx, dy float64, dt float64) (crossX, crossY int) {
	dist := p.Speed * dt

	// Try X axis
	if dx != 0 {
		newX := p.X + dx*dist
		if !TileCollision(screen, newX, p.Y, p.Width, p.Height) {
			p.X = newX
		}
	}

	// Try Y axis
	if dy != 0 {
		newY := p.Y + dy*dist
		if !TileCollision(screen, p.X, newY, p.Width, p.Height) {
			p.Y = newY
		}
	}

	// Check edge crossings
	if p.X < 0 {
		crossX = -1
	} else if p.X+float64(p.Width) > float64(config.PlayAreaWidth) {
		crossX = 1
	}
	if p.Y < 0 {
		crossY = -1
	} else if p.Y+float64(p.Height) > float64(config.PlayAreaHeight) {
		crossY = 1
	}

	return crossX, crossY
}
