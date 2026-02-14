package render

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
	"github.com/AchrafSoltani/glow"
)

var (
	ColorRupee          = glow.RGB(50, 200, 50)
	ColorRupeeDark      = glow.RGB(30, 140, 30)
	ColorKey            = glow.RGB(230, 200, 50)
	ColorKeyDark        = glow.RGB(180, 150, 30)
	ColorSword          = glow.RGB(200, 200, 220)
	ColorSwordDark      = glow.RGB(140, 140, 160)
	ColorHeartContainer = glow.RGB(255, 50, 50)
	ColorHeartContGold  = glow.RGB(230, 190, 50)
)

// DrawItem renders an item sprite at its position with bobbing animation.
func DrawItem(sc *ScaledCanvas, item *entity.Item) {
	DrawItemAt(sc, item, 0, 0)
}

// DrawItemAt renders an item with a pixel offset.
func DrawItemAt(sc *ScaledCanvas, item *entity.Item, offsetX, offsetY int) {
	if item.Collected {
		return
	}

	bob := int(item.BobOffset())
	px := int(item.X) + offsetX
	py := int(item.Y) + config.HUDHeight + offsetY + bob

	switch item.Type {
	case entity.ItemHeart:
		drawItemHeart(sc, px, py)
	case entity.ItemRupee:
		drawItemRupee(sc, px, py)
	case entity.ItemKey:
		drawItemKey(sc, px, py)
	case entity.ItemSword:
		drawItemSword(sc, px, py)
	case entity.ItemHeartContainer:
		drawItemHeartContainer(sc, px, py)
	}
}

func drawItemHeart(sc *ScaledCanvas, px, py int) {
	// Small heart
	sc.FillCircle(px+3, py+3, 2, ColorHeartFull)
	sc.FillCircle(px+7, py+3, 2, ColorHeartFull)
	sc.DrawRect(px+2, py+4, 8, 4, ColorHeartFull)
	sc.DrawRect(px+3, py+8, 6, 2, ColorHeartFull)
	sc.DrawRect(px+4, py+10, 4, 1, ColorHeartFull)
}

func drawItemRupee(sc *ScaledCanvas, px, py int) {
	// Green diamond
	sc.SetPixel(px+5, py+1, ColorRupee)
	sc.DrawRect(px+4, py+2, 3, 1, ColorRupee)
	sc.DrawRect(px+3, py+3, 5, 1, ColorRupee)
	sc.DrawRect(px+2, py+4, 7, 3, ColorRupee)
	sc.DrawRect(px+3, py+7, 5, 1, ColorRupee)
	sc.DrawRect(px+4, py+8, 3, 1, ColorRupee)
	sc.SetPixel(px+5, py+9, ColorRupee)
	// Highlight
	sc.SetPixel(px+4, py+4, ColorRupeeDark)
	sc.SetPixel(px+4, py+5, ColorRupeeDark)
}

func drawItemKey(sc *ScaledCanvas, px, py int) {
	// Key head (circle)
	sc.FillCircle(px+6, py+3, 3, ColorKey)
	sc.FillCircle(px+6, py+3, 1, ColorKeyDark)
	// Shaft
	sc.DrawRect(px+5, py+5, 2, 5, ColorKey)
	// Teeth
	sc.DrawRect(px+7, py+8, 2, 1, ColorKey)
	sc.DrawRect(px+7, py+10, 2, 1, ColorKey)
}

func drawItemSword(sc *ScaledCanvas, px, py int) {
	// Blade
	sc.DrawRect(px+5, py+1, 2, 7, ColorSword)
	// Guard
	sc.DrawRect(px+3, py+7, 6, 1, ColorSwordDark)
	// Handle
	sc.DrawRect(px+5, py+8, 2, 3, ColorKeyDark)
}

func drawItemHeartContainer(sc *ScaledCanvas, px, py int) {
	// Large heart with gold border
	sc.FillCircle(px+3, py+3, 3, ColorHeartContGold)
	sc.FillCircle(px+8, py+3, 3, ColorHeartContGold)
	sc.DrawRect(px+1, py+4, 10, 4, ColorHeartContGold)
	sc.DrawRect(px+2, py+8, 8, 2, ColorHeartContGold)
	// Inner heart
	sc.FillCircle(px+3, py+3, 2, ColorHeartContainer)
	sc.FillCircle(px+8, py+3, 2, ColorHeartContainer)
	sc.DrawRect(px+2, py+4, 8, 3, ColorHeartContainer)
	sc.DrawRect(px+3, py+7, 6, 2, ColorHeartContainer)
}
