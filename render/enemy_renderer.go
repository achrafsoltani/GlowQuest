package render

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
	"github.com/AchrafSoltani/glow"
)

var (
	ColorOctorok     = glow.RGB(200, 50, 50)
	ColorOctorokDark = glow.RGB(140, 30, 30)
	ColorMoblin      = glow.RGB(160, 100, 50)
	ColorMoblinDark  = glow.RGB(110, 70, 30)
	ColorStalfos     = glow.RGB(180, 180, 180)
	ColorStalfosDark = glow.RGB(120, 120, 120)
)

// DrawEnemy renders an enemy sprite at its position.
func DrawEnemy(sc *ScaledCanvas, e *entity.Enemy) {
	DrawEnemyAt(sc, e, 0, 0)
}

// DrawEnemyAt renders an enemy with a pixel offset.
func DrawEnemyAt(sc *ScaledCanvas, e *entity.Enemy, offsetX, offsetY int) {
	if e.Dead {
		return
	}

	// Flash during invincibility (skip draw on alternating frames)
	if e.InvTimer > 0 {
		frame := int(e.InvTimer * 20)
		if frame%2 == 0 {
			return
		}
	}

	px := int(e.X) + offsetX
	py := int(e.Y) + config.HUDHeight + offsetY

	switch e.Type {
	case entity.EnemyOctorok:
		drawOctorok(sc, px, py, e)
	case entity.EnemyMoblin:
		drawMoblin(sc, px, py, e)
	case entity.EnemyStalfos:
		drawStalfos(sc, px, py, e)
	case entity.EnemyBoss:
		drawBoss(sc, px, py, e)
	}
}

func drawOctorok(sc *ScaledCanvas, px, py int, e *entity.Enemy) {
	// Round red body
	sc.FillCircle(px+7, py+7, 6, ColorOctorok)
	sc.FillCircle(px+7, py+7, 4, ColorOctorokDark)

	// Eyes based on direction
	switch e.Dir {
	case entity.DirDown:
		sc.SetPixel(px+5, py+6, ColorBG)
		sc.SetPixel(px+9, py+6, ColorBG)
	case entity.DirUp:
		sc.SetPixel(px+5, py+5, ColorBG)
		sc.SetPixel(px+9, py+5, ColorBG)
	case entity.DirLeft:
		sc.SetPixel(px+4, py+6, ColorBG)
		sc.SetPixel(px+4, py+8, ColorBG)
	case entity.DirRight:
		sc.SetPixel(px+10, py+6, ColorBG)
		sc.SetPixel(px+10, py+8, ColorBG)
	}

	// Little legs (walk animation)
	legOff := 0
	if e.WalkFrame%2 == 1 {
		legOff = 1
	}
	sc.DrawRect(px+3+legOff, py+12, 2, 2, ColorOctorokDark)
	sc.DrawRect(px+9-legOff, py+12, 2, 2, ColorOctorokDark)
}

func drawMoblin(sc *ScaledCanvas, px, py int, e *entity.Enemy) {
	// Brown pig body
	sc.DrawRect(px+3, py+2, 8, 5, ColorMoblin)
	// Snout
	sc.DrawRect(px+5, py+4, 4, 3, ColorMoblinDark)
	// Eyes
	sc.SetPixel(px+5, py+3, ColorBG)
	sc.SetPixel(px+8, py+3, ColorBG)
	// Body
	sc.DrawRect(px+3, py+7, 8, 4, ColorMoblinDark)
	// Legs
	legOff := 0
	if e.WalkFrame%2 == 1 {
		legOff = 1
	}
	sc.DrawRect(px+4+legOff, py+11, 3, 3, ColorMoblin)
	sc.DrawRect(px+8-legOff, py+11, 3, 3, ColorMoblin)
}

func drawStalfos(sc *ScaledCanvas, px, py int, e *entity.Enemy) {
	// Grey skeleton
	// Skull
	sc.DrawRect(px+4, py+1, 6, 5, ColorStalfos)
	// Eye sockets
	sc.SetPixel(px+5, py+3, ColorBG)
	sc.SetPixel(px+8, py+3, ColorBG)
	// Jaw
	sc.DrawRect(px+5, py+5, 4, 1, ColorStalfosDark)
	// Ribcage
	sc.DrawRect(px+4, py+6, 6, 3, ColorStalfosDark)
	sc.SetPixel(px+5, py+7, ColorStalfos)
	sc.SetPixel(px+8, py+7, ColorStalfos)
	// Spine
	sc.DrawRect(px+6, py+9, 2, 2, ColorStalfosDark)
	// Legs
	legOff := 0
	if e.WalkFrame%2 == 1 {
		legOff = 1
	}
	sc.DrawRect(px+4+legOff, py+11, 2, 3, ColorStalfos)
	sc.DrawRect(px+8-legOff, py+11, 2, 3, ColorStalfos)
}

func drawBoss(sc *ScaledCanvas, px, py int, e *entity.Enemy) {
	// Large dark purple skeleton boss (20×20)
	// Oversized skull
	sc.DrawRect(px+4, py+1, 12, 8, ColorBoss)
	sc.DrawRect(px+5, py+2, 10, 6, ColorBossDark)
	// Red eyes
	sc.DrawRect(px+6, py+4, 2, 2, ColorBossEye)
	sc.DrawRect(px+12, py+4, 2, 2, ColorBossEye)
	// Jaw
	sc.DrawRect(px+7, py+8, 6, 2, ColorBoss)
	// Armoured shoulders
	sc.DrawRect(px+1, py+9, 5, 3, ColorBoss)
	sc.DrawRect(px+14, py+9, 5, 3, ColorBoss)
	// Body / ribcage
	sc.DrawRect(px+6, py+9, 8, 5, ColorBossDark)
	sc.SetPixel(px+8, py+11, ColorBoss)
	sc.SetPixel(px+11, py+11, ColorBoss)
	// Legs
	legOff := 0
	if e.WalkFrame%2 == 1 {
		legOff = 1
	}
	sc.DrawRect(px+6+legOff, py+14, 3, 5, ColorBoss)
	sc.DrawRect(px+11-legOff, py+14, 3, 5, ColorBoss)
	// Feet
	sc.DrawRect(px+5+legOff, py+18, 4, 2, ColorBossDark)
	sc.DrawRect(px+11-legOff, py+18, 4, 2, ColorBossDark)
}
