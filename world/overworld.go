package world

import "github.com/AchrafSoltani/GlowQuest/config"

type Overworld struct {
	Screens  map[[2]int]*Screen
	Width    int // grid width (16)
	Height   int // grid height (16)
	CurrentX int
	CurrentY int
}

func NewOverworld() *Overworld {
	meta := LoadOverworldMeta()
	screens := LoadOverworldScreens()

	ow := &Overworld{
		Screens:  screens,
		Width:    meta.Width,
		Height:   meta.Height,
		CurrentX: meta.StartScreen[0],
		CurrentY: meta.StartScreen[1],
	}

	// Update config overworld dimensions
	config.OverworldW = meta.Width
	config.OverworldH = meta.Height

	return ow
}

func (ow *Overworld) CurrentScreen() *Screen {
	s := ow.Screens[[2]int{ow.CurrentX, ow.CurrentY}]
	if s == nil {
		// Return an empty screen (all grass) for unmapped positions
		return &Screen{}
	}
	return s
}

func (ow *Overworld) ScreenAt(x, y int) *Screen {
	if x < 0 || x >= ow.Width || y < 0 || y >= ow.Height {
		return nil
	}
	return ow.Screens[[2]int{x, y}]
}

func (ow *Overworld) CanMove(dx, dy int) bool {
	nx := ow.CurrentX + dx
	ny := ow.CurrentY + dy
	if nx < 0 || nx >= ow.Width || ny < 0 || ny >= ow.Height {
		return false
	}
	// Allow movement if the target screen exists
	_, exists := ow.Screens[[2]int{nx, ny}]
	return exists
}

func (ow *Overworld) Move(dx, dy int) {
	ow.CurrentX += dx
	ow.CurrentY += dy
}
