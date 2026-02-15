package render

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
	"github.com/AchrafSoltani/glow"
)

var (
	ColorInvBG       = glow.RGB(20, 20, 40)
	ColorInvBorder   = glow.RGB(180, 180, 180)
	ColorInvCursor   = glow.RGB(255, 255, 100)
	ColorInvSlot     = glow.RGB(40, 40, 60)
	ColorInvSlotSel  = glow.RGB(80, 80, 100)
	ColorInvAssignA  = glow.RGB(200, 50, 50)
	ColorInvAssignB  = glow.RGB(50, 50, 200)
)

// DrawInventoryScreen draws the inventory overlay.
func DrawInventoryScreen(sc *ScaledCanvas, inv *entity.Inventory, cursorX, cursorY int) {
	// Darken background
	for y := 0; y < config.WindowHeight; y++ {
		for x := 0; x < config.WindowWidth; x++ {
			if (x+y)%3 != 0 {
				sc.SetPixel(x, y, ColorInvBG)
			}
		}
	}

	// Title
	title := "INVENTORY"
	tw := TextWidth(title)
	DrawText(sc, title, (config.WindowWidth-tw)/2, 4, ColorHUDText)

	// A/B assignment display at top
	aLabel := "A:"
	bLabel := "B:"
	aName := entity.EquipItemName(inv.ButtonA)
	bName := entity.EquipItemName(inv.ButtonB)
	if aName == "" {
		aName = "---"
	}
	if bName == "" {
		bName = "---"
	}
	DrawText(sc, aLabel, 10, 18, ColorInvAssignA)
	DrawText(sc, aName, 24, 18, ColorHUDText)
	DrawText(sc, bLabel, 130, 18, ColorInvAssignB)
	DrawText(sc, bName, 144, 18, ColorHUDText)

	// Item grid (5 columns x 3 rows)
	items := inv.OwnedItemsList()
	gridX := 28
	gridY := 36
	slotW := 40
	slotH := 28

	for row := 0; row < 3; row++ {
		for col := 0; col < 5; col++ {
			sx := gridX + col*slotW
			sy := gridY + row*slotH
			idx := row*5 + col

			// Slot background
			bgColor := ColorInvSlot
			if col == cursorX && row == cursorY {
				bgColor = ColorInvSlotSel
			}
			sc.DrawRect(sx, sy, slotW-2, slotH-2, bgColor)

			// Cursor highlight
			if col == cursorX && row == cursorY {
				sc.DrawRectOutline(sx-1, sy-1, slotW, slotH, ColorInvCursor)
			}

			// Draw item if present
			if idx < len(items) {
				itemID := items[idx]
				name := entity.EquipItemName(itemID)
				if len(name) > 6 {
					name = name[:6]
				}

				// Show A/B indicator
				if itemID == inv.ButtonA {
					sc.DrawRect(sx, sy, 6, 6, ColorInvAssignA)
					DrawText(sc, "A", sx+1, sy+1, ColorHUDText)
				}
				if itemID == inv.ButtonB {
					bx := sx + slotW - 8
					sc.DrawRect(bx, sy, 6, 6, ColorInvAssignB)
					DrawText(sc, "B", bx+1, sy+1, ColorHUDText)
				}

				// Item icon (small representation)
				drawInventoryItemIcon(sc, itemID, sx+slotW/2-6, sy+4)
				// Item name
				nw := TextWidth(name)
				DrawText(sc, name, sx+(slotW-2-nw)/2, sy+slotH-8, ColorHUDText)
			}
		}
	}

	// Instructions at bottom
	inst := "Z:SET A  X:SET B  TAB:CLOSE"
	iw := TextWidth(inst)
	DrawText(sc, inst, (config.WindowWidth-iw)/2, config.WindowHeight-12, ColorMenuDisabled)
}

func drawInventoryItemIcon(sc *ScaledCanvas, id entity.EquipItemID, x, y int) {
	switch id {
	case entity.EquipSword:
		sc.DrawRect(x+5, y, 2, 8, ColorSwordBlade)
		sc.DrawRect(x+3, y+8, 6, 1, ColorSword)
		sc.DrawRect(x+5, y+9, 2, 3, ColorKeyDark)
	case entity.EquipShield:
		sc.DrawRect(x+2, y, 8, 10, glow.RGB(60, 60, 200))
		sc.DrawRect(x+4, y+2, 4, 6, glow.RGB(200, 50, 50))
	case entity.EquipBow:
		sc.DrawLine(x+2, y+2, x+2, y+10, glow.RGB(140, 80, 20))
		sc.DrawLine(x+2, y+2, x+8, y+6, glow.RGB(140, 80, 20))
		sc.DrawLine(x+2, y+10, x+8, y+6, glow.RGB(140, 80, 20))
	case entity.EquipBomb:
		sc.FillCircle(x+6, y+7, 4, glow.RGB(40, 40, 40))
		sc.DrawRect(x+5, y, 2, 3, glow.RGB(200, 100, 30))
	case entity.EquipRocsFeather:
		sc.DrawRect(x+3, y+2, 2, 8, glow.RGB(200, 200, 200))
		sc.DrawRect(x+5, y+1, 4, 5, glow.RGB(230, 230, 240))
	default:
		// Generic item icon
		sc.DrawRect(x+2, y+2, 8, 8, glow.RGB(150, 150, 150))
		sc.DrawRectOutline(x+2, y+2, 8, 8, ColorHUDText)
	}
}
