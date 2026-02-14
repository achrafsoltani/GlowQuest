package render

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
)

func DrawPlayer(sc *ScaledCanvas, p *entity.Player) {
	DrawPlayerAt(sc, p, 0, 0)
}

// DrawPlayerAt draws the player with a pixel offset (for scrolling transitions).
func DrawPlayerAt(sc *ScaledCanvas, p *entity.Player, offsetX, offsetY int) {
	// Player is drawn in the play area (offset by HUD height)
	px := int(p.X) + offsetX
	py := int(p.Y) + config.HUDHeight + offsetY

	// Walk animation leg offset
	legOff := 0
	if p.Moving {
		switch p.WalkFrame {
		case 0:
			legOff = -1
		case 2:
			legOff = 1
		}
	}

	switch p.Dir {
	case entity.DirDown:
		drawPlayerDown(sc, px, py, legOff)
	case entity.DirUp:
		drawPlayerUp(sc, px, py, legOff)
	case entity.DirLeft:
		drawPlayerLeft(sc, px, py, legOff)
	case entity.DirRight:
		drawPlayerRight(sc, px, py, legOff)
	}
}

func drawPlayerDown(sc *ScaledCanvas, px, py, legOff int) {
	// Hat
	sc.DrawRect(px+2, py, 10, 4, ColorHat)
	// Head (face)
	sc.DrawRect(px+3, py+2, 8, 5, ColorSkin)
	// Eyes
	sc.SetPixel(px+5, py+4, ColorBG)
	sc.SetPixel(px+8, py+4, ColorBG)
	// Body
	sc.DrawRect(px+3, py+7, 8, 4, ColorTunic)
	// Legs
	sc.DrawRect(px+4+legOff, py+11, 3, 3, ColorBoot)
	sc.DrawRect(px+8-legOff, py+11, 3, 3, ColorBoot)
}

func drawPlayerUp(sc *ScaledCanvas, px, py, legOff int) {
	// Hat (from behind)
	sc.DrawRect(px+2, py, 10, 5, ColorHat)
	// Head (back of hair)
	sc.DrawRect(px+3, py+3, 8, 4, ColorHat)
	// Body
	sc.DrawRect(px+3, py+7, 8, 4, ColorTunic)
	// Legs
	sc.DrawRect(px+4+legOff, py+11, 3, 3, ColorBoot)
	sc.DrawRect(px+8-legOff, py+11, 3, 3, ColorBoot)
}

func drawPlayerLeft(sc *ScaledCanvas, px, py, legOff int) {
	// Hat (side)
	sc.DrawRect(px+1, py, 8, 4, ColorHat)
	// Head
	sc.DrawRect(px+3, py+2, 7, 5, ColorSkin)
	// Eye
	sc.SetPixel(px+4, py+4, ColorBG)
	// Body
	sc.DrawRect(px+3, py+7, 8, 4, ColorTunic)
	// Legs
	sc.DrawRect(px+4, py+11+legOff, 3, 3, ColorBoot)
	sc.DrawRect(px+7, py+11-legOff, 3, 3, ColorBoot)
}

func drawPlayerRight(sc *ScaledCanvas, px, py, legOff int) {
	// Hat (side)
	sc.DrawRect(px+5, py, 8, 4, ColorHat)
	// Head
	sc.DrawRect(px+4, py+2, 7, 5, ColorSkin)
	// Eye
	sc.SetPixel(px+9, py+4, ColorBG)
	// Body
	sc.DrawRect(px+3, py+7, 8, 4, ColorTunic)
	// Legs
	sc.DrawRect(px+4, py+11-legOff, 3, 3, ColorBoot)
	sc.DrawRect(px+7, py+11+legOff, 3, 3, ColorBoot)
}
