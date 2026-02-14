package render

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
	"github.com/AchrafSoltani/GlowQuest/world"
	"github.com/AchrafSoltani/glow"
)

func DrawScreen(sc *ScaledCanvas, screen *world.Screen) {
	DrawScreenAt(sc, screen, 0, 0)
}

// DrawScreenAt draws a screen with a pixel offset (for scrolling transitions).
func DrawScreenAt(sc *ScaledCanvas, screen *world.Screen, offsetX, offsetY int) {
	for gy := 0; gy < config.ScreenGridH; gy++ {
		for gx := 0; gx < config.ScreenGridW; gx++ {
			drawTileAt(sc, screen.Tiles[gy][gx], gx, gy, offsetX, offsetY)
		}
	}
}

// DrawTransition renders both old and new screens sliding during a transition.
// The easeProgress should be the eased value (0→1) of the transition progress.
func DrawTransition(sc *ScaledCanvas, oldScreen, newScreen *world.Screen, player *entity.Player, dirX, dirY int, easeProgress float64) {
	// Calculate pixel offsets based on direction and progress
	totalX := dirX * config.PlayAreaWidth
	totalY := dirY * config.PlayAreaHeight

	// Old screen slides out in the transition direction
	oldOffX := -int(float64(totalX) * easeProgress)
	oldOffY := -int(float64(totalY) * easeProgress)

	// New screen slides in from the opposite side
	newOffX := totalX + oldOffX
	newOffY := totalY + oldOffY

	// Draw old screen sliding out
	DrawScreenAt(sc, oldScreen, oldOffX, oldOffY)
	// Draw new screen sliding in
	DrawScreenAt(sc, newScreen, newOffX, newOffY)

	// Draw player on the new screen, offset by the new screen's slide amount
	DrawPlayerAt(sc, player, newOffX, newOffY)
}

// 4×4 Bayer ordered dithering matrix (values 0–15)
var bayerMatrix = [4][4]int{
	{0, 8, 2, 10},
	{12, 4, 14, 6},
	{3, 11, 1, 9},
	{15, 7, 13, 5},
}

// DrawFade draws a black overlay using 4×4 Bayer ordered dithering for smooth fade.
func DrawFade(sc *ScaledCanvas, progress float64) {
	if progress <= 0 {
		return
	}
	if progress >= 1.0 {
		sc.DrawRect(0, config.HUDHeight, config.PlayAreaWidth, config.PlayAreaHeight, ColorBG)
		return
	}
	threshold := int(progress * 16)
	for y := 0; y < config.PlayAreaHeight; y++ {
		for x := 0; x < config.PlayAreaWidth; x++ {
			if bayerMatrix[y%4][x%4] < threshold {
				sc.SetPixel(x, config.HUDHeight+y, ColorBG)
			}
		}
	}
}

// DrawFlash draws a white dithered overlay for item pickup flash.
func DrawFlash(sc *ScaledCanvas, intensity float64) {
	if intensity <= 0 {
		return
	}
	white := glow.RGB(255, 255, 255)
	threshold := int(intensity * 16)
	for y := 0; y < config.PlayAreaHeight; y++ {
		for x := 0; x < config.PlayAreaWidth; x++ {
			if bayerMatrix[y%4][x%4] < threshold {
				sc.SetPixel(x, config.HUDHeight+y, white)
			}
		}
	}
}

func drawTile(sc *ScaledCanvas, tile world.TileType, gx, gy int) {
	drawTileAt(sc, tile, gx, gy, 0, 0)
}

func drawTileAt(sc *ScaledCanvas, tile world.TileType, gx, gy int, offsetX, offsetY int) {
	ts := config.TileSize
	px := gx*ts + offsetX
	py := gy*ts + config.HUDHeight + offsetY

	switch tile {
	case world.TileGrass:
		sc.DrawRect(px, py, ts, ts, ColorGrass)
		// Grass blade details
		sc.SetPixel(px+3, py+4, ColorGrassDark)
		sc.SetPixel(px+7, py+9, ColorGrassDark)
		sc.SetPixel(px+12, py+3, ColorGrassDark)
		sc.SetPixel(px+5, py+13, ColorGrassDark)
		sc.SetPixel(px+11, py+11, ColorGrassDark)

	case world.TileWall:
		sc.DrawRect(px, py, ts, ts, ColorWall)
		// Brick pattern
		sc.DrawLine(px, py+4, px+ts-1, py+4, ColorWallDark)
		sc.DrawLine(px, py+8, px+ts-1, py+8, ColorWallDark)
		sc.DrawLine(px, py+12, px+ts-1, py+12, ColorWallDark)
		sc.DrawLine(px+8, py, px+8, py+4, ColorWallDark)
		sc.DrawLine(px+4, py+4, px+4, py+8, ColorWallDark)
		sc.DrawLine(px+12, py+4, px+12, py+8, ColorWallDark)
		sc.DrawLine(px+8, py+8, px+8, py+12, ColorWallDark)
		sc.DrawLine(px+4, py+12, px+4, py+ts-1, ColorWallDark)
		sc.DrawLine(px+12, py+12, px+12, py+ts-1, ColorWallDark)

	case world.TileWater:
		sc.DrawRect(px, py, ts, ts, ColorWater)
		// Wave lines
		sc.DrawLine(px+1, py+4, px+5, py+3, ColorWaterLt)
		sc.DrawLine(px+5, py+3, px+9, py+5, ColorWaterLt)
		sc.DrawLine(px+2, py+10, px+6, py+9, ColorWaterLt)
		sc.DrawLine(px+6, py+9, px+10, py+11, ColorWaterLt)

	case world.TileTree:
		// Grass background
		sc.DrawRect(px, py, ts, ts, ColorGrass)
		// Trunk
		sc.DrawRect(px+6, py+10, 4, 6, ColorBoot)
		// Canopy (layered circles approximation)
		sc.FillCircle(px+8, py+6, 6, ColorTree)
		sc.FillCircle(px+5, py+8, 4, ColorTreeDark)
		sc.FillCircle(px+11, py+8, 4, ColorTreeDark)

	case world.TileSand:
		sc.DrawRect(px, py, ts, ts, ColorSand)
		sc.SetPixel(px+4, py+6, ColorWall)
		sc.SetPixel(px+10, py+3, ColorWall)
		sc.SetPixel(px+7, py+12, ColorWall)

	case world.TileFloor:
		sc.DrawRect(px, py, ts, ts, ColorFloor)
		sc.DrawRectOutline(px, py, ts, ts, ColorFloorDark)

	case world.TileStairs:
		sc.DrawRect(px, py, ts, ts, ColorFloor)
		// Stair steps
		for i := 0; i < 4; i++ {
			y := py + 2 + i*3
			sc.DrawRect(px+2+i, y, ts-4-i*2, 2, ColorFloorDark)
		}

	case world.TileDoorLocked:
		sc.DrawRect(px, py, ts, ts, ColorWall)
		sc.DrawRect(px+4, py+4, 8, 12, ColorBoot)
		sc.SetPixel(px+10, py+9, ColorSand) // Keyhole

	case world.TileDoorOpen:
		sc.DrawRect(px, py, ts, ts, ColorFloor)
		sc.DrawRect(px+4, py+0, 8, 2, ColorWall)
	}
}
