package render

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
	"github.com/AchrafSoltani/glow"
)

func DrawHUD(sc *ScaledCanvas, p *entity.Player) {
	// HUD background
	sc.DrawRect(0, 0, config.WindowWidth, config.HUDHeight, ColorHUD)

	// Draw hearts
	heartSize := 8
	spacing := 2
	startX := 4
	startY := (config.HUDHeight - heartSize) / 2

	fullHearts := p.HP / 2
	halfHeart := p.HP%2 == 1
	totalHearts := p.MaxHP / 2

	for i := 0; i < totalHearts; i++ {
		hx := startX + i*(heartSize+spacing)
		hy := startY

		if i < fullHearts {
			drawHeart(sc, hx, hy, heartSize, ColorHeartFull)
		} else if i == fullHearts && halfHeart {
			drawHeart(sc, hx, hy, heartSize, ColorHeartEmpty)
			drawHalfHeart(sc, hx, hy, heartSize, ColorHeartFull)
		} else {
			drawHeart(sc, hx, hy, heartSize, ColorHeartEmpty)
		}
	}
}

func drawHeart(sc *ScaledCanvas, x, y, size int, color glow.Color) {
	// Heart shape: two bumps on top, pointed bottom
	half := size / 2
	q := size / 4

	// Top left bump
	sc.FillCircle(x+q+1, y+q+1, q, color)
	// Top right bump
	sc.FillCircle(x+half+q-1, y+q+1, q, color)
	// Bottom triangle (rectangle approximation)
	sc.DrawRect(x+1, y+q, size-2, half, color)
	// Point
	for i := 0; i < q; i++ {
		w := size - 2 - i*2
		if w > 0 {
			sc.DrawRect(x+1+i, y+q+half+i-1, w, 1, color)
		}
	}
}

func drawHalfHeart(sc *ScaledCanvas, x, y, size int, color glow.Color) {
	half := size / 2
	q := size / 4

	sc.FillCircle(x+q+1, y+q+1, q, color)
	sc.DrawRect(x+1, y+q, half-1, half, color)
	for i := 0; i < q; i++ {
		w := half - 1 - i
		if w > 0 {
			sc.DrawRect(x+1+i, y+q+half+i-1, w, 1, color)
		}
	}
}
