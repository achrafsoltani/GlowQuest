package world

import "github.com/AchrafSoltani/GlowQuest/config"

type Overworld struct {
	Screens  [config.OverworldH][config.OverworldW]*Screen
	CurrentX int // 0–2
	CurrentY int // 0–2
}

func NewOverworld() *Overworld {
	ow := &Overworld{
		CurrentX: 1,
		CurrentY: 1,
	}
	// Row 0 (top)
	ow.Screens[0][0] = ForestScreen()
	ow.Screens[0][1] = ForestPathScreen()
	ow.Screens[0][2] = MountainScreen()
	// Row 1 (middle)
	ow.Screens[1][0] = LakeShoreScreen()
	ow.Screens[1][1] = VillageScreen()
	ow.Screens[1][2] = EastFieldScreen()
	// Row 2 (bottom)
	ow.Screens[2][0] = SwampScreen()
	ow.Screens[2][1] = SouthFieldScreen()
	ow.Screens[2][2] = RuinsScreen()
	return ow
}

func (ow *Overworld) CurrentScreen() *Screen {
	return ow.Screens[ow.CurrentY][ow.CurrentX]
}

func (ow *Overworld) ScreenAt(x, y int) *Screen {
	if x < 0 || x >= config.OverworldW || y < 0 || y >= config.OverworldH {
		return nil
	}
	return ow.Screens[y][x]
}

func (ow *Overworld) CanMove(dx, dy int) bool {
	nx := ow.CurrentX + dx
	ny := ow.CurrentY + dy
	return nx >= 0 && nx < config.OverworldW && ny >= 0 && ny < config.OverworldH
}

func (ow *Overworld) Move(dx, dy int) {
	ow.CurrentX += dx
	ow.CurrentY += dy
}
