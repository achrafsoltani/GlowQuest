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

	// Tarin — brown/orange mushroom hunter
	colorTarinRobe     = glow.RGB(160, 100, 40)
	colorTarinRobeDark = glow.RGB(120, 70, 25)
	colorTarinHood     = glow.RGB(140, 85, 30)

	// Marin — red/orange dress
	colorMarinRobe     = glow.RGB(200, 80, 60)
	colorMarinRobeDark = glow.RGB(160, 55, 40)
	colorMarinHood     = glow.RGB(180, 65, 50)
	colorMarinHair     = glow.RGB(200, 120, 50)

	// Madam MeowMeow — purple
	colorMeowRobe     = glow.RGB(140, 60, 160)
	colorMeowRobeDark = glow.RGB(100, 40, 120)
	colorMeowHood     = glow.RGB(120, 50, 140)

	// Shopkeeper — green
	colorShopRobe     = glow.RGB(50, 140, 70)
	colorShopRobeDark = glow.RGB(35, 100, 50)
	colorShopHood     = glow.RGB(40, 120, 60)
)

// DrawNPC renders an NPC sprite at its position.
func DrawNPC(sc *ScaledCanvas, npc *entity.NPC) {
	DrawNPCAt(sc, npc, 0, 0)
}

// DrawNPCAt renders an NPC with a pixel offset.
func DrawNPCAt(sc *ScaledCanvas, npc *entity.NPC, offsetX, offsetY int) {
	px := int(npc.X) + offsetX
	py := int(npc.Y) + config.HUDHeight + offsetY

	hood, robe, robeDark := ColorNPCHood, ColorNPCRobe, ColorNPCRobeDark

	switch npc.ID {
	case "tarin":
		hood, robe, robeDark = colorTarinHood, colorTarinRobe, colorTarinRobeDark
	case "marin":
		hood, robe, robeDark = colorMarinHood, colorMarinRobe, colorMarinRobeDark
	case "meowmeow":
		hood, robe, robeDark = colorMeowHood, colorMeowRobe, colorMeowRobeDark
	case "shopkeeper":
		hood, robe, robeDark = colorShopHood, colorShopRobe, colorShopRobeDark
	}

	// Special rendering for Marin (hair instead of hood)
	if npc.ID == "marin" {
		// Hair
		sc.DrawRect(px+2, py, 10, 3, colorMarinHair)
		// Face
		sc.DrawRect(px+4, py+2, 6, 4, ColorSkin)
		// Eyes
		sc.SetPixel(px+5, py+4, ColorBG)
		sc.SetPixel(px+8, py+4, ColorBG)
		// Dress body
		sc.DrawRect(px+3, py+6, 8, 5, robe)
		sc.DrawRect(px+2, py+7, 10, 3, robeDark)
		// Feet
		sc.DrawRect(px+4, py+11, 3, 3, ColorBoot)
		sc.DrawRect(px+8, py+11, 3, 3, ColorBoot)
		return
	}

	// Standard NPC rendering
	// Hood
	sc.DrawRect(px+3, py, 8, 4, hood)
	// Face
	sc.DrawRect(px+4, py+2, 6, 4, ColorSkin)
	// Eyes
	sc.SetPixel(px+5, py+4, ColorBG)
	sc.SetPixel(px+8, py+4, ColorBG)
	// Robe body
	sc.DrawRect(px+3, py+6, 8, 5, robe)
	sc.DrawRect(px+2, py+7, 10, 3, robeDark)
	// Feet
	sc.DrawRect(px+4, py+11, 3, 3, ColorBoot)
	sc.DrawRect(px+8, py+11, 3, 3, ColorBoot)
}
