package game

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
	"github.com/AchrafSoltani/GlowQuest/render"
	"github.com/AchrafSoltani/GlowQuest/system"
	"github.com/AchrafSoltani/GlowQuest/world"
	"github.com/AchrafSoltani/glow"
)

type Game struct {
	State      GameState
	Player     *entity.Player
	Overworld  *world.Overworld
	Input      *system.InputTracker
	Layout     config.Layout
	Transition Transition
}

func NewGame() *Game {
	// Place player on a clear grass tile in the village (row 7, col 7 — sand path)
	startX := float64(7*config.TileSize) + 1
	startY := float64(7*config.TileSize) + 1

	return &Game{
		State:     StatePlaying,
		Player:    entity.NewPlayer(startX, startY),
		Overworld: world.NewOverworld(),
		Input:     system.NewInputTracker(),
		Layout:    config.NewLayout(config.WindowWidth*4, config.WindowHeight*4),
	}
}

func (g *Game) KeyDown(key glow.Key) {
	g.Input.KeyDown(key)
}

func (g *Game) KeyUp(key glow.Key) {
	g.Input.KeyUp(key)
}

func (g *Game) OnResize(w, h int) {
	g.Layout = config.NewLayout(w, h)
}

func (g *Game) Update(dt float64) {
	if g.State != StatePlaying {
		g.Input.Update()
		return
	}

	// Handle active transition
	if g.Transition.Active {
		g.Transition.Timer += dt
		if g.Transition.Done() {
			g.Transition.Active = false
			g.Transition.OldScreen = nil
		}
		g.Input.Update()
		return
	}

	// Read movement input
	var dx, dy float64
	if g.Input.IsHeld(glow.KeyUp) || g.Input.IsHeld(glow.KeyW) {
		dy = -1
		g.Player.Dir = entity.DirUp
	}
	if g.Input.IsHeld(glow.KeyDown) || g.Input.IsHeld(glow.KeyS) {
		dy = 1
		g.Player.Dir = entity.DirDown
	}
	if g.Input.IsHeld(glow.KeyLeft) || g.Input.IsHeld(glow.KeyA) {
		dx = -1
		g.Player.Dir = entity.DirLeft
	}
	if g.Input.IsHeld(glow.KeyRight) || g.Input.IsHeld(glow.KeyD) {
		dx = 1
		g.Player.Dir = entity.DirRight
	}

	g.Player.Moving = dx != 0 || dy != 0

	if g.Player.Moving {
		crossX, crossY := system.MovePlayer(g.Player, g.Overworld.CurrentScreen(), dx, dy, dt)
		g.handleEdgeCrossing(crossX, crossY)
	}

	g.Player.UpdateAnimation(dt)
	g.Input.Update()
}

func (g *Game) handleEdgeCrossing(crossX, crossY int) {
	// Prioritise horizontal crossings (only handle one axis at a time)
	dirX, dirY := 0, 0
	if crossX != 0 {
		dirX = crossX
	} else if crossY != 0 {
		dirY = crossY
	} else {
		return
	}

	if !g.Overworld.CanMove(dirX, dirY) {
		// Clamp player to screen bounds (world boundary)
		g.clampPlayer()
		return
	}

	// Save old screen and move to the new one
	oldScreen := g.Overworld.CurrentScreen()
	g.Overworld.Move(dirX, dirY)

	// Reposition player to the opposite edge
	if dirX == 1 {
		g.Player.X = 1
	} else if dirX == -1 {
		g.Player.X = float64(config.PlayAreaWidth-g.Player.Width) - 1
	}
	if dirY == 1 {
		g.Player.Y = 1
	} else if dirY == -1 {
		g.Player.Y = float64(config.PlayAreaHeight-g.Player.Height) - 1
	}

	// Start transition
	g.Transition.Start(dirX, dirY, oldScreen)
}

func (g *Game) clampPlayer() {
	if g.Player.X < 0 {
		g.Player.X = 0
	}
	maxX := float64(config.PlayAreaWidth - g.Player.Width)
	if g.Player.X > maxX {
		g.Player.X = maxX
	}
	if g.Player.Y < 0 {
		g.Player.Y = 0
	}
	maxY := float64(config.PlayAreaHeight - g.Player.Height)
	if g.Player.Y > maxY {
		g.Player.Y = maxY
	}
}

func (g *Game) Draw(canvas *glow.Canvas) {
	sc := render.NewScaledCanvas(canvas, g.Layout)
	sc.Clear(render.ColorBG)

	if g.Transition.Active {
		progress := easeInOut(g.Transition.Progress())
		render.DrawTransition(
			sc,
			g.Transition.OldScreen,
			g.Overworld.CurrentScreen(),
			g.Player,
			g.Transition.DirX,
			g.Transition.DirY,
			progress,
		)
	} else {
		render.DrawScreen(sc, g.Overworld.CurrentScreen())
		render.DrawPlayer(sc, g.Player)
	}

	render.DrawHUD(sc, g.Player)
}
