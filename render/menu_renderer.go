package render

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/glow"
)

// MenuOption mirrors game.MenuOption to avoid circular imports.
type MenuOption struct {
	Label    string
	Disabled bool
}

func DrawTitleScreen(sc *ScaledCanvas, options []MenuOption, selectedIdx int) {
	sc.DrawRect(0, 0, config.WindowWidth, config.WindowHeight, ColorBG)

	// Title "GLOWQUEST" drawn double-sized at top
	title := "GLOWQUEST"
	titleW := len(title) * 10
	titleX := (config.WindowWidth - titleW) / 2
	titleY := 40
	drawDoubleText(sc, title, titleX, titleY, ColorMenuTitle)

	// Sword graphic under title
	swordX := config.WindowWidth/2 - 1
	swordY := 75
	sc.DrawRect(swordX, swordY, 2, 14, ColorSwordBlade)
	sc.DrawRect(swordX-3, swordY+14, 8, 2, ColorSword)
	sc.DrawRect(swordX, swordY+16, 2, 4, ColorKeyDark)

	// Menu options
	optY := 120
	for i, opt := range options {
		color := ColorHUDText
		if opt.Disabled {
			color = ColorMenuDisabled
		}
		label := "  " + opt.Label
		if i == selectedIdx {
			label = "> " + opt.Label
		}
		labelW := TextWidth(label)
		lx := (config.WindowWidth - labelW) / 2
		DrawText(sc, label, lx, optY, color)
		optY += 16
	}
}

func DrawGameOverScreen(sc *ScaledCanvas, timer float64) {
	sc.DrawRect(0, 0, config.WindowWidth, config.WindowHeight, ColorBG)

	text := "GAME OVER"
	tw := len(text) * 10
	tx := (config.WindowWidth - tw) / 2
	ty := 80
	drawDoubleText(sc, text, tx, ty, ColorHeartFull)

	if timer >= config.GameOverDelay {
		blink := int(timer * 3)
		if blink%2 == 0 {
			prompt := "PRESS ENTER"
			pw := TextWidth(prompt)
			px := (config.WindowWidth - pw) / 2
			DrawText(sc, prompt, px, 140, ColorHUDText)
		}
	}
}

func DrawPauseOverlay(sc *ScaledCanvas) {
	for y := 0; y < config.WindowHeight; y++ {
		for x := 0; x < config.WindowWidth; x++ {
			if (x+y)%2 == 0 {
				sc.SetPixel(x, y, ColorBG)
			}
		}
	}

	text := "PAUSED"
	tw := TextWidth(text)
	tx := (config.WindowWidth - tw) / 2
	DrawText(sc, text, tx, 100, ColorHUDText)

	sub := "PRESS ESC TO RESUME"
	sw := TextWidth(sub)
	sx := (config.WindowWidth - sw) / 2
	DrawText(sc, sub, sx, 120, ColorMenuDisabled)
}

func DrawVictoryScreen(sc *ScaledCanvas, timer float64) {
	sc.DrawRect(0, 0, config.WindowWidth, config.WindowHeight, ColorBG)

	text := "VICTORY!"
	tw := len(text) * 10
	tx := (config.WindowWidth - tw) / 2
	ty := 60
	drawDoubleText(sc, text, tx, ty, ColorVictoryGold)

	line1 := "The ancient evil"
	line2 := "has been vanquished!"
	l1w := TextWidth(line1)
	l2w := TextWidth(line2)
	DrawText(sc, line1, (config.WindowWidth-l1w)/2, 100, ColorHUDText)
	DrawText(sc, line2, (config.WindowWidth-l2w)/2, 112, ColorHUDText)

	if timer >= config.VictoryDelay {
		blink := int(timer * 3)
		if blink%2 == 0 {
			prompt := "PRESS ENTER"
			pw := TextWidth(prompt)
			px := (config.WindowWidth - pw) / 2
			DrawText(sc, prompt, px, 150, ColorHUDText)
		}
	}
}

func drawDoubleText(sc *ScaledCanvas, text string, x, y int, color glow.Color) {
	cx := x
	for i := 0; i < len(text); i++ {
		ch := text[i]
		glyph, ok := glyphData[ch]
		if !ok {
			cx += 10
			continue
		}
		for row := 0; row < glyphH; row++ {
			bits := glyph[row]
			for col := 0; col < glyphW; col++ {
				if bits&(0x8>>col) != 0 {
					sc.DrawRect(cx+col*2, y+row*2, 2, 2, color)
				}
			}
		}
		cx += 10
	}
}
