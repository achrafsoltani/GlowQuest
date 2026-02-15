package render

import (
	"fmt"

	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
	"github.com/AchrafSoltani/glow"
)

func DrawHUD(sc *ScaledCanvas, p *entity.Player) {
	// HUD background (32px tall)
	sc.DrawRect(0, 0, config.WindowWidth, config.HUDHeight, ColorHUD)

	// --- Left section: B and A item boxes ---
	boxSize := 20
	boxY := (config.HUDHeight - boxSize) / 2

	// B button box (left)
	bBoxX := 4
	sc.DrawRect(bBoxX, boxY, boxSize, boxSize, glow.RGB(30, 30, 60))
	sc.DrawRectOutline(bBoxX, boxY, boxSize, boxSize, glow.RGB(120, 120, 160))
	DrawText(sc, "B", bBoxX+1, boxY-1, glow.RGB(100, 100, 200))
	drawHUDItemIcon(sc, p.Inventory.ButtonB, bBoxX+3, boxY+4)

	// A button box
	aBoxX := bBoxX + boxSize + 4
	sc.DrawRect(aBoxX, boxY, boxSize, boxSize, glow.RGB(60, 30, 30))
	sc.DrawRectOutline(aBoxX, boxY, boxSize, boxSize, glow.RGB(160, 120, 120))
	DrawText(sc, "A", aBoxX+1, boxY-1, glow.RGB(200, 100, 100))
	drawHUDItemIcon(sc, p.Inventory.ButtonA, aBoxX+3, boxY+4)

	// --- Centre section: Hearts ---
	heartSize := 8
	spacing := 1
	totalHearts := p.MaxHP / 2
	heartsW := totalHearts * (heartSize + spacing)
	heartStartX := (config.WindowWidth-heartsW)/2 + 20
	heartStartY := 4

	fullHearts := p.HP / 2
	halfHeart := p.HP%2 == 1

	for i := 0; i < totalHearts; i++ {
		hx := heartStartX + i*(heartSize+spacing)
		hy := heartStartY

		if i < fullHearts {
			drawHeart(sc, hx, hy, heartSize, ColorHeartFull)
		} else if i == fullHearts && halfHeart {
			drawHeart(sc, hx, hy, heartSize, ColorHeartEmpty)
			drawHalfHeart(sc, hx, hy, heartSize, ColorHeartFull)
		} else {
			drawHeart(sc, hx, hy, heartSize, ColorHeartEmpty)
		}
	}

	// --- Bottom row: Rupees and Keys ---
	infoY := 18

	// Rupee icon + count
	rupeeX := heartStartX
	sc.SetPixel(rupeeX+1, infoY+1, ColorRupee)
	sc.DrawRect(rupeeX, infoY+2, 3, 2, ColorRupee)
	sc.SetPixel(rupeeX+1, infoY+4, ColorRupee)
	DrawText(sc, fmt.Sprintf("%03d", p.Inventory.Rupees), rupeeX+5, infoY+1, ColorHUDText)

	// Key icon + count
	keyX := rupeeX + 40
	sc.FillCircle(keyX+2, infoY+2, 2, ColorKey)
	sc.DrawRect(keyX+1, infoY+3, 2, 3, ColorKey)
	sc.SetPixel(keyX+3, infoY+4, ColorKey)
	DrawText(sc, fmt.Sprintf("%d", p.Inventory.Keys), keyX+6, infoY+1, ColorHUDText)

	// Bombs count (if player has bombs)
	if p.Inventory.OwnedItems[entity.EquipBomb] {
		bombX := keyX + 25
		sc.FillCircle(bombX+2, infoY+3, 2, glow.RGB(40, 40, 40))
		sc.SetPixel(bombX+2, infoY, glow.RGB(200, 100, 30))
		DrawText(sc, fmt.Sprintf("%02d", p.Inventory.Bombs), bombX+6, infoY+1, ColorHUDText)
	}

	// Separator line
	sc.DrawLine(0, config.HUDHeight-1, config.WindowWidth-1, config.HUDHeight-1, glow.RGB(60, 60, 60))
}

func drawHUDItemIcon(sc *ScaledCanvas, id entity.EquipItemID, x, y int) {
	switch id {
	case entity.EquipSword:
		sc.DrawRect(x+5, y, 2, 8, ColorSwordBlade)
		sc.DrawRect(x+3, y+8, 6, 1, ColorSword)
		sc.DrawRect(x+5, y+9, 2, 2, ColorKeyDark)
	case entity.EquipShield:
		sc.DrawRect(x+2, y+1, 8, 9, glow.RGB(60, 60, 200))
		sc.DrawRect(x+4, y+3, 4, 5, glow.RGB(200, 50, 50))
	case entity.EquipBow:
		sc.DrawLine(x+2, y+1, x+2, y+9, glow.RGB(140, 80, 20))
		sc.DrawLine(x+2, y+1, x+8, y+5, glow.RGB(140, 80, 20))
		sc.DrawLine(x+2, y+9, x+8, y+5, glow.RGB(140, 80, 20))
	case entity.EquipBomb:
		sc.FillCircle(x+6, y+6, 4, glow.RGB(40, 40, 40))
		sc.DrawRect(x+5, y, 2, 3, glow.RGB(200, 100, 30))
	case entity.EquipRocsFeather:
		sc.DrawRect(x+3, y+2, 2, 8, glow.RGB(200, 200, 200))
		sc.DrawRect(x+5, y+1, 4, 5, glow.RGB(230, 230, 240))
	case entity.EquipNone:
		// empty slot
	default:
		sc.DrawRect(x+2, y+2, 8, 8, glow.RGB(100, 100, 100))
	}
}

func drawHeart(sc *ScaledCanvas, x, y, size int, color glow.Color) {
	half := size / 2
	q := size / 4

	sc.FillCircle(x+q+1, y+q+1, q, color)
	sc.FillCircle(x+half+q-1, y+q+1, q, color)
	sc.DrawRect(x+1, y+q, size-2, half, color)
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

	sc.DrawRect(x-1, y-1, barW+2, barH+2, ColorBG)
	sc.DrawRectOutline(x-1, y-1, barW+2, barH+2, ColorHeartFull)

	fillW := barW * hp / maxHP
	if fillW < 0 {
		fillW = 0
	}
	sc.DrawRect(x, y, fillW, barH, ColorHeartFull)

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
