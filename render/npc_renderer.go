package render

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
	"github.com/AchrafSoltani/glow"
)

var (
	ColorNPCRobe     = glow.RGB(60, 80, 180)
	ColorNPCRobeDark = glow.RGB(40, 55, 130)
	ColorNPCHood     = glow.RGB(50, 70, 160)
)

// DrawNPC renders an NPC sprite at its position.
func DrawNPC(sc *ScaledCanvas, npc *entity.NPC) {
	DrawNPCAt(sc, npc, 0, 0)
}

// DrawNPCAt renders an NPC with a pixel offset.
func DrawNPCAt(sc *ScaledCanvas, npc *entity.NPC, offsetX, offsetY int) {
	px := int(npc.X) + offsetX
	py := int(npc.Y) + config.HUDHeight + offsetY

	// Hood
	sc.DrawRect(px+3, py, 8, 4, ColorNPCHood)
	// Face
	sc.DrawRect(px+4, py+2, 6, 4, ColorSkin)
	// Eyes
	sc.SetPixel(px+5, py+4, ColorBG)
	sc.SetPixel(px+8, py+4, ColorBG)
	// Robe body
	sc.DrawRect(px+3, py+6, 8, 5, ColorNPCRobe)
	sc.DrawRect(px+2, py+7, 10, 3, ColorNPCRobeDark)
	// Feet
	sc.DrawRect(px+4, py+11, 3, 3, ColorBoot)
	sc.DrawRect(px+8, py+11, 3, 3, ColorBoot)
}
