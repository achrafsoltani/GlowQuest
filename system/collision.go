package system

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/world"
)

func AABBOverlap(ax, ay, aw, ah, bx, by, bw, bh float64) bool {
	return ax < bx+bw && ax+aw > bx && ay < by+bh && ay+ah > by
}

// ProximityCheck returns true if the point (px, py) is within radius of (nx, ny).
func ProximityCheck(px, py, nx, ny, radius float64) bool {
	dx := px - nx
	dy := py - ny
	return dx*dx+dy*dy <= radius*radius
}

func TileCollision(screen *world.Screen, x, y float64, w, h int) bool {
	fw := float64(w)
	fh := float64(h)
	ts := float64(config.TileSize)

	startX := int(x / ts)
	startY := int(y / ts)
	endX := int((x + fw - 0.01) / ts)
	endY := int((y + fh - 0.01) / ts)

	for gy := startY; gy <= endY; gy++ {
		for gx := startX; gx <= endX; gx++ {
			tile := screen.TileAt(gx, gy)
			if !tile.IsPassable() {
				tileX := float64(gx) * ts
				tileY := float64(gy) * ts
				if AABBOverlap(x, y, fw, fh, tileX, tileY, ts, ts) {
					return true
				}
			}
		}
	}
	return false
}
