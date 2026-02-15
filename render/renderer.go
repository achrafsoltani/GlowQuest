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
func DrawTransition(sc *ScaledCanvas, oldScreen, newScreen *world.Screen, player *entity.Player, dirX, dirY int, easeProgress float64) {
	totalX := dirX * config.PlayAreaWidth
	totalY := dirY * config.PlayAreaHeight

	oldOffX := -int(float64(totalX) * easeProgress)
	oldOffY := -int(float64(totalY) * easeProgress)

	newOffX := totalX + oldOffX
	newOffY := totalY + oldOffY

	DrawScreenAt(sc, oldScreen, oldOffX, oldOffY)
	DrawScreenAt(sc, newScreen, newOffX, newOffY)
	DrawPlayerAt(sc, player, newOffX, newOffY)
}

// 4×4 Bayer ordered dithering matrix (values 0–15)
var bayerMatrix = [4][4]int{
	{0, 8, 2, 10},
	{12, 4, 14, 6},
	{3, 11, 1, 9},
	{15, 7, 13, 5},
}

// DrawFade draws a black overlay using 4×4 Bayer ordered dithering.
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
		sc.SetPixel(px+3, py+4, ColorGrassDark)
		sc.SetPixel(px+7, py+9, ColorGrassDark)
		sc.SetPixel(px+12, py+3, ColorGrassDark)
		sc.SetPixel(px+5, py+13, ColorGrassDark)
		sc.SetPixel(px+11, py+11, ColorGrassDark)

	case world.TileWall:
		sc.DrawRect(px, py, ts, ts, ColorWall)
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
		sc.DrawLine(px+1, py+4, px+5, py+3, ColorWaterLt)
		sc.DrawLine(px+5, py+3, px+9, py+5, ColorWaterLt)
		sc.DrawLine(px+2, py+10, px+6, py+9, ColorWaterLt)
		sc.DrawLine(px+6, py+9, px+10, py+11, ColorWaterLt)

	case world.TileTree:
		sc.DrawRect(px, py, ts, ts, ColorGrass)
		sc.DrawRect(px+6, py+10, 4, 6, ColorBoot)
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
		for i := 0; i < 4; i++ {
			y := py + 2 + i*3
			sc.DrawRect(px+2+i, y, ts-4-i*2, 2, ColorFloorDark)
		}

	case world.TileDoorLocked:
		sc.DrawRect(px, py, ts, ts, ColorWall)
		sc.DrawRect(px+4, py+4, 8, 12, ColorBoot)
		sc.SetPixel(px+10, py+9, ColorSand)

	case world.TileDoorOpen:
		sc.DrawRect(px, py, ts, ts, ColorFloor)
		sc.DrawRect(px+4, py+0, 8, 2, ColorWall)

	case world.TileShallowWater:
		sc.DrawRect(px, py, ts, ts, ColorShallowWater)
		sc.SetPixel(px+3, py+5, ColorWaterLt)
		sc.SetPixel(px+10, py+8, ColorWaterLt)

	case world.TileCliffN, world.TileCliffS, world.TileCliffE, world.TileCliffW:
		sc.DrawRect(px, py, ts, ts, ColorGrass)
		sc.DrawRect(px, py, ts, 3, ColorCliff)
		sc.DrawLine(px, py+2, px+ts-1, py+2, ColorCliffEdge)

	case world.TileBridge:
		sc.DrawRect(px, py, ts, ts, ColorBridge)
		sc.DrawLine(px+2, py, px+2, py+ts-1, ColorBridgeDark)
		sc.DrawLine(px+ts-3, py, px+ts-3, py+ts-1, ColorBridgeDark)
		sc.DrawLine(px, py+4, px+ts-1, py+4, ColorBridgeDark)
		sc.DrawLine(px, py+11, px+ts-1, py+11, ColorBridgeDark)

	case world.TilePit:
		sc.DrawRect(px, py, ts, ts, ColorPit)
		sc.DrawRectOutline(px, py, ts, ts, ColorFloorDark)

	case world.TileBush:
		sc.DrawRect(px, py, ts, ts, ColorGrass)
		sc.FillCircle(px+8, py+8, 6, ColorBush)
		sc.FillCircle(px+5, py+6, 3, ColorBushDark)
		sc.FillCircle(px+11, py+10, 3, ColorBushDark)

	case world.TileRock:
		sc.DrawRect(px, py, ts, ts, ColorGrass)
		sc.DrawRect(px+2, py+4, 12, 10, ColorRock)
		sc.DrawRect(px+3, py+5, 10, 8, ColorRockDark)
		sc.DrawLine(px+4, py+6, px+8, py+6, ColorRock)

	case world.TileHeavyRock:
		sc.DrawRect(px, py, ts, ts, ColorGrass)
		sc.DrawRect(px+1, py+3, 14, 11, ColorRockDark)
		sc.DrawRect(px+2, py+4, 12, 9, ColorRock)
		sc.SetPixel(px+4, py+6, ColorRockDark)
		sc.SetPixel(px+10, py+8, ColorRockDark)

	case world.TilePot:
		sc.DrawRect(px, py, ts, ts, ColorFloor)
		sc.DrawRect(px+3, py+4, 10, 10, ColorPot)
		sc.DrawRect(px+4, py+2, 8, 3, ColorPotDark)
		sc.DrawRect(px+5, py+1, 6, 2, ColorPot)

	case world.TileSignpost:
		sc.DrawRect(px, py, ts, ts, ColorGrass)
		sc.DrawRect(px+6, py+8, 4, 6, ColorBoot)
		sc.DrawRect(px+2, py+2, 12, 7, ColorSignpost)
		sc.DrawRectOutline(px+2, py+2, 12, 7, ColorBoot)

	case world.TileChest:
		sc.DrawRect(px, py, ts, ts, ColorFloor)
		sc.DrawRect(px+2, py+4, 12, 10, ColorChest)
		sc.DrawRect(px+2, py+4, 12, 4, ColorChestDark)
		sc.DrawRect(px+6, py+7, 4, 3, ColorKeyBlock)

	case world.TileChestOpen:
		sc.DrawRect(px, py, ts, ts, ColorFloor)
		sc.DrawRect(px+2, py+6, 12, 8, ColorChest)
		sc.DrawRect(px+2, py+2, 12, 4, ColorChestDark)

	case world.TileOwlStatue:
		sc.DrawRect(px, py, ts, ts, ColorGrass)
		sc.DrawRect(px+4, py+2, 8, 12, ColorOwlStatue)
		sc.SetPixel(px+6, py+4, ColorBG)
		sc.SetPixel(px+9, py+4, ColorBG)
		sc.DrawRect(px+6, py+7, 4, 2, ColorWallDark)

	case world.TileKeyBlock:
		sc.DrawRect(px, py, ts, ts, ColorFloor)
		sc.DrawRect(px+2, py+2, 12, 12, ColorKeyBlock)
		sc.DrawRectOutline(px+2, py+2, 12, 12, ColorWallDark)
		sc.FillCircle(px+8, py+7, 2, ColorWallDark)
		sc.DrawRect(px+7, py+8, 2, 4, ColorWallDark)

	case world.TileCrackedWall:
		sc.DrawRect(px, py, ts, ts, ColorWall)
		sc.DrawLine(px+3, py+3, px+8, py+6, ColorCrackedWall)
		sc.DrawLine(px+8, py+6, px+12, py+4, ColorCrackedWall)
		sc.DrawLine(px+5, py+8, px+10, py+11, ColorCrackedWall)

	case world.TileConveyorN, world.TileConveyorS, world.TileConveyorE, world.TileConveyorW:
		sc.DrawRect(px, py, ts, ts, ColorConveyor)
		sc.DrawRect(px+2, py+2, ts-4, ts-4, ColorFloor)
		// Arrow marks
		switch tile {
		case world.TileConveyorN:
			sc.SetPixel(px+7, py+3, ColorConveyorMark)
			sc.DrawRect(px+6, py+5, 4, 1, ColorConveyorMark)
		case world.TileConveyorS:
			sc.SetPixel(px+7, py+11, ColorConveyorMark)
			sc.DrawRect(px+6, py+9, 4, 1, ColorConveyorMark)
		case world.TileConveyorE:
			sc.SetPixel(px+11, py+7, ColorConveyorMark)
			sc.DrawRect(px+9, py+6, 1, 4, ColorConveyorMark)
		case world.TileConveyorW:
			sc.SetPixel(px+3, py+7, ColorConveyorMark)
			sc.DrawRect(px+5, py+6, 1, 4, ColorConveyorMark)
		}

	case world.TileSpikes:
		sc.DrawRect(px, py, ts, ts, ColorFloor)
		// Draw spike pattern
		for sy := 0; sy < 4; sy++ {
			for sx := 0; sx < 4; sx++ {
				tipX := px + 2 + sx*4
				tipY := py + 2 + sy*4
				sc.SetPixel(tipX, tipY, ColorSpikesTip)
				sc.SetPixel(tipX-1, tipY+1, ColorSpikes)
				sc.SetPixel(tipX+1, tipY+1, ColorSpikes)
			}
		}

	case world.TileLava:
		sc.DrawRect(px, py, ts, ts, ColorLava)
		sc.SetPixel(px+3, py+4, ColorLavaDark)
		sc.SetPixel(px+10, py+8, ColorLavaDark)
		sc.SetPixel(px+6, py+11, ColorFlame)

	case world.TileIce:
		sc.DrawRect(px, py, ts, ts, ColorIce)
		sc.SetPixel(px+4, py+3, ColorIceDark)
		sc.SetPixel(px+10, py+9, ColorIceDark)
		sc.DrawLine(px+2, py+6, px+6, py+4, glow.RGB(220, 240, 255))

	case world.TileSwitchOff:
		sc.DrawRect(px, py, ts, ts, ColorFloor)
		sc.DrawRect(px+4, py+4, 8, 8, ColorSwitchOff)
		sc.DrawRectOutline(px+4, py+4, 8, 8, ColorWallDark)

	case world.TileSwitchOn:
		sc.DrawRect(px, py, ts, ts, ColorFloor)
		sc.DrawRect(px+4, py+4, 8, 8, ColorSwitch)
		sc.DrawRectOutline(px+4, py+4, 8, 8, ColorWallDark)

	case world.TileWarpTile:
		sc.DrawRect(px, py, ts, ts, ColorFloor)
		sc.FillCircle(px+8, py+8, 5, ColorWarp)
		sc.FillCircle(px+8, py+8, 3, ColorFloor)
		sc.FillCircle(px+8, py+8, 1, ColorWarp)

	case world.TileBossLocked:
		sc.DrawRect(px, py, ts, ts, ColorBossLocked)
		sc.DrawRect(px+4, py+2, 8, 12, ColorWallDark)
		sc.FillCircle(px+8, py+7, 2, ColorBossEye)
		sc.DrawRect(px+7, py+9, 2, 3, ColorBossEye)

	case world.TileBombable:
		sc.DrawRect(px, py, ts, ts, ColorWall)
		sc.SetPixel(px+4, py+4, ColorCrackedWall)
		sc.SetPixel(px+8, py+8, ColorCrackedWall)
		sc.SetPixel(px+11, py+5, ColorCrackedWall)

	case world.TileTorch:
		sc.DrawRect(px, py, ts, ts, ColorFloor)
		sc.DrawRect(px+5, py+4, 6, 10, ColorTorch)
		sc.DrawRect(px+6, py+2, 4, 3, ColorWallDark)

	case world.TileTorchLit:
		sc.DrawRect(px, py, ts, ts, ColorFloor)
		sc.DrawRect(px+5, py+6, 6, 8, ColorTorch)
		sc.FillCircle(px+8, py+4, 3, ColorFlame)
		sc.FillCircle(px+8, py+3, 2, glow.RGB(255, 220, 100))

	case world.TileGrassFlower:
		sc.DrawRect(px, py, ts, ts, ColorGrass)
		sc.SetPixel(px+4, py+5, ColorFlower)
		sc.SetPixel(px+3, py+6, ColorGrassDark)
		sc.SetPixel(px+5, py+6, ColorGrassDark)
		sc.SetPixel(px+11, py+9, ColorFlower)
		sc.SetPixel(px+10, py+10, ColorGrassDark)
		sc.SetPixel(px+12, py+10, ColorGrassDark)

	case world.TilePathH:
		sc.DrawRect(px, py, ts, ts, ColorPath)
		sc.DrawLine(px, py+2, px+ts-1, py+2, ColorSand)
		sc.DrawLine(px, py+ts-3, px+ts-1, py+ts-3, ColorSand)

	case world.TilePathV:
		sc.DrawRect(px, py, ts, ts, ColorPath)
		sc.DrawLine(px+2, py, px+2, py+ts-1, ColorSand)
		sc.DrawLine(px+ts-3, py, px+ts-3, py+ts-1, ColorSand)

	case world.TileHouseFront:
		sc.DrawRect(px, py, ts, ts, ColorHouseFront)
		sc.DrawLine(px, py+ts-1, px+ts-1, py+ts-1, ColorWallDark)

	case world.TileRoof:
		sc.DrawRect(px, py, ts, ts, ColorRoof)
		sc.DrawLine(px, py+ts-1, px+ts-1, py+ts-1, glow.RGB(120, 40, 30))

	case world.TileWindow:
		sc.DrawRect(px, py, ts, ts, ColorHouseFront)
		sc.DrawRect(px+3, py+3, 10, 10, ColorWindow)
		sc.DrawLine(px+8, py+3, px+8, py+12, ColorHouseFront)
		sc.DrawLine(px+3, py+8, px+12, py+8, ColorHouseFront)

	case world.TileFenceH:
		sc.DrawRect(px, py, ts, ts, ColorGrass)
		sc.DrawRect(px, py+5, ts, 2, ColorFence)
		sc.DrawRect(px, py+10, ts, 2, ColorFence)
		sc.DrawRect(px+2, py+5, 2, 7, ColorFence)
		sc.DrawRect(px+12, py+5, 2, 7, ColorFence)

	case world.TileFenceV:
		sc.DrawRect(px, py, ts, ts, ColorGrass)
		sc.DrawRect(px+5, py, 2, ts, ColorFence)
		sc.DrawRect(px+10, py, 2, ts, ColorFence)
		sc.DrawRect(px+5, py+2, 7, 2, ColorFence)
		sc.DrawRect(px+5, py+12, 7, 2, ColorFence)

	default:
		// Unknown tile — render as grass
		sc.DrawRect(px, py, ts, ts, ColorGrass)
	}
}
