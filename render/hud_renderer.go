package render

import (
	"fmt"

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

	// Rupee count (right side of HUD)
	rupeeX := config.WindowWidth - 80
	rupeeY := 2
	// Small green diamond icon
	sc.SetPixel(rupeeX+1, rupeeY+1, ColorRupee)
	sc.DrawRect(rupeeX, rupeeY+2, 3, 2, ColorRupee)
	sc.SetPixel(rupeeX+1, rupeeY+4, ColorRupee)
	// Count text
	DrawText(sc, fmt.Sprintf("%d", p.Inventory.Rupees), rupeeX+5, rupeeY+1, ColorHUDText)

	// Key count
	keyX := config.WindowWidth - 45
	keyY := 2
	// Small key icon
	sc.FillCircle(keyX+2, keyY+2, 2, ColorKey)
	sc.DrawRect(keyX+1, keyY+3, 2, 4, ColorKey)
	sc.SetPixel(keyX+3, keyY+5, ColorKey)
	// Count text
	DrawText(sc, fmt.Sprintf("%d", p.Inventory.Keys), keyX+6, keyY+1, ColorHUDText)

	// Sword indicator
	if p.HasSword {
		swordX := config.WindowWidth - 15
		swordY := 3
		sc.DrawRect(swordX, swordY, 2, 6, ColorSword)
		sc.DrawRect(swordX-1, swordY+6, 4, 1, ColorSwordDark)
		sc.DrawRect(swordX, swordY+7, 2, 2, ColorKeyDark)
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

func DrawBossHealthBar(sc *ScaledCanvas, hp, maxHP int) {
	barW := 80
	barH := 4
	x := (config.WindowWidth - barW) / 2
	y := config.HUDHeight + 2

	// Background
	sc.DrawRect(x-1, y-1, barW+2, barH+2, ColorBG)
	sc.DrawRectOutline(x-1, y-1, barW+2, barH+2, ColorHeartFull)

	// Fill
	fillW := barW * hp / maxHP
	if fillW < 0 {
		fillW = 0
	}
	sc.DrawRect(x, y, fillW, barH, ColorHeartFull)

	// "BOSS" label
	label := "BOSS"
	lw := TextWidth(label)
	DrawText(sc, label, (config.WindowWidth-lw)/2, y+barH+2, ColorHeartFull)
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
