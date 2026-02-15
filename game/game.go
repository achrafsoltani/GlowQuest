package game

import (
	"fmt"

	"github.com/AchrafSoltani/GlowQuest/audio"
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
	"github.com/AchrafSoltani/GlowQuest/render"
	"github.com/AchrafSoltani/GlowQuest/save"
	"github.com/AchrafSoltani/GlowQuest/system"
	"github.com/AchrafSoltani/GlowQuest/world"
	"github.com/AchrafSoltani/glow"
)

// LocationType identifies what kind of area the player is in.
type LocationType int

const (
	LocationOverworld LocationType = iota
	LocationInterior
	LocationDungeon
)

type Game struct {
	State      GameState
	Player     *entity.Player
	Overworld  *world.Overworld
	Input      *system.InputTracker
	Layout     config.Layout
	Transition Transition
	Audio      *audio.Engine

	// Combat & entities
	Enemies     []*entity.Enemy
	Projectiles []*entity.Projectile
	Items       []*entity.Item
	NPCs        []*entity.NPC
	RNG         *system.SimpleRNG

	// Dialogue
	Dialogue DialogueState

	// Persistent item tracking — items collected are remembered by "sx,sy,idx"
	CollectedItems map[string]bool
	UnlockedDoors  map[string]bool

	// Interiors
	Interiors       map[string]*world.InteriorDef
	DoorLinks       []world.DoorLink
	InInterior      bool
	CurrentInterior *world.InteriorDef
	ReturnLink      *world.DoorLink

	// Location tracking
	Location LocationType

	// Dungeon support (Phase 8+)
	CurrentDungeon *world.Dungeon

	// Quest state
	Quest *QuestState

	// Fade transition state for interiors
	PendingInterior *world.InteriorDef
	PendingDoorLink *world.DoorLink
	PendingExitLink *world.DoorLink

	// Menu
	Menu           MenuState
	GameOverTimer  float64
	VictoryTimer   float64
	ShouldQuit     bool

	// Polish
	Particles  *entity.ParticlePool
	ShakeTimer float64
	FlashTimer float64

	// Boss
	BossDefeated bool

	// Inventory screen cursor
	InventoryCursorX int
	InventoryCursorY int
}

func NewGame() *Game {
	g := &Game{
		State:          StateMenu,
		Input:          system.NewInputTracker(),
		Layout:         config.NewLayout(config.WindowWidth*4, config.WindowHeight*4),
		RNG:            system.NewSimpleRNG(42),
		CollectedItems: make(map[string]bool),
		UnlockedDoors:  make(map[string]bool),
		Audio:          audio.NewEngine(),
		Menu:           NewMenuState(save.Exists()),
		Particles:      entity.NewParticlePool(),
		Quest:          NewQuestState(),
	}
	return g
}

func (g *Game) initWorld() {
	g.Overworld = world.NewOverworld()
	g.Interiors, _ = world.LoadInteriors()
	g.DoorLinks = world.BuildDoorLinksFromScreens(g.Overworld.Screens)
}

func (g *Game) StartNewGame() {
	g.initWorld()

	meta := world.LoadOverworldMeta()
	startX := meta.StartPos[0]
	startY := meta.StartPos[1]

	g.Player = entity.NewPlayer(startX, startY)
	g.CollectedItems = make(map[string]bool)
	g.UnlockedDoors = make(map[string]bool)
	g.InInterior = false
	g.CurrentInterior = nil
	g.ReturnLink = nil
	g.Location = LocationOverworld
	g.CurrentDungeon = nil
	g.BossDefeated = false
	g.Quest = NewQuestState()
	g.Particles = entity.NewParticlePool()
	g.ShakeTimer = 0
	g.FlashTimer = 0
	g.State = StatePlaying

	g.Quest.SetFlag("game_started")

	// Start in interior if configured
	if meta.StartInterior != "" {
		interior, ok := g.Interiors[meta.StartInterior]
		if ok {
			g.InInterior = true
			g.Location = LocationInterior
			g.CurrentInterior = interior
			// Place player in the centre of the interior
			g.Player.X = float64(config.PlayAreaWidth/2 - g.Player.Width/2)
			g.Player.Y = float64(config.PlayAreaHeight/2 - g.Player.Height/2)
			// Find a door link back to overworld for this interior
			for i := range g.DoorLinks {
				if g.DoorLinks[i].InteriorID == meta.StartInterior {
					g.ReturnLink = &g.DoorLinks[i]
					break
				}
			}
			// Also check the interior's own door links for overworld exit
			if g.ReturnLink == nil {
				for i := range interior.DoorLinks {
					if interior.DoorLinks[i].InteriorID == "overworld" {
						// Create a synthetic return link — exit one tile below the door
						exitX := interior.DoorLinks[i].ExitX
						exitY := interior.DoorLinks[i].ExitY
						if exitX == 0 && exitY == 0 {
							exitX = meta.StartPos[0]
							exitY = meta.StartPos[1]
						}
						g.ReturnLink = &world.DoorLink{
							ScreenX:    meta.StartScreen[0],
							ScreenY:    meta.StartScreen[1],
							DoorTileX:  interior.DoorLinks[i].DoorTileX,
							DoorTileY:  interior.DoorLinks[i].DoorTileY,
							InteriorID: meta.StartInterior,
							SpawnX:     g.Player.X,
							SpawnY:     g.Player.Y,
							ExitX:      exitX,
							ExitY:      exitY,
						}
						break
					}
				}
			}
		}
	}

	g.spawnScreenEntities()
}

func (g *Game) StartContinue() {
	data := save.Load()
	if data == nil {
		g.StartNewGame()
		return
	}

	g.initWorld()

	// Handle V2 save data
	g.Player = entity.NewPlayer(data.PlayerX, data.PlayerY)
	g.Player.HasSword = data.HasSword
	g.Player.MaxHP = data.MaxHP
	g.Player.HP = data.HP
	g.Player.Inventory.Rupees = data.Rupees
	g.Player.Inventory.Keys = data.Keys

	// Load V2 inventory fields if present
	if data.Version >= 2 {
		g.Player.Inventory.Bombs = data.Bombs
		g.Player.Inventory.Arrows = data.Arrows
		g.Player.Inventory.SwordLevel = data.SwordLevel
		g.Player.Inventory.ShieldLevel = data.ShieldLevel
		g.Player.Inventory.BraceletLevel = data.BraceletLevel
		g.Player.Inventory.ButtonA = entity.EquipItemID(data.ButtonA)
		g.Player.Inventory.ButtonB = entity.EquipItemID(data.ButtonB)
		for _, id := range data.OwnedItems {
			g.Player.Inventory.OwnedItems[entity.EquipItemID(id)] = true
		}
		// Restore quest state
		if data.Quest != nil {
			g.Quest.Flags = data.Quest.Flags
			if g.Quest.Flags == nil {
				g.Quest.Flags = make(map[string]bool)
			}
			g.Quest.DungeonsCompleted = data.Quest.DungeonsCompleted
			g.Quest.TradingItem = data.Quest.TradingItem
		}
	}

	// Backwards compat: if player has sword but no equip, auto-assign
	if g.Player.HasSword && !g.Player.Inventory.OwnedItems[entity.EquipSword] {
		g.Player.Inventory.OwnedItems[entity.EquipSword] = true
		if g.Player.Inventory.SwordLevel == 0 {
			g.Player.Inventory.SwordLevel = 1
		}
		if g.Player.Inventory.ButtonA == entity.EquipNone {
			g.Player.Inventory.ButtonA = entity.EquipSword
		}
	}

	g.CollectedItems = data.CollectedItems
	if g.CollectedItems == nil {
		g.CollectedItems = make(map[string]bool)
	}
	g.UnlockedDoors = data.UnlockedDoors
	if g.UnlockedDoors == nil {
		g.UnlockedDoors = make(map[string]bool)
	}
	g.BossDefeated = data.BossDefeated
	g.Particles = entity.NewParticlePool()
	g.ShakeTimer = 0
	g.FlashTimer = 0
	g.Location = LocationOverworld

	g.Overworld.CurrentX = data.ScreenX
	g.Overworld.CurrentY = data.ScreenY

	if data.InInterior && data.InteriorID != "" {
		interior, ok := g.Interiors[data.InteriorID]
		if ok {
			g.InInterior = true
			g.CurrentInterior = interior
			g.Location = LocationInterior
			// Find the door link for return
			for i := range g.DoorLinks {
				if g.DoorLinks[i].InteriorID == data.InteriorID {
					g.ReturnLink = &g.DoorLinks[i]
					break
				}
			}
		}
	}

	g.State = StatePlaying
	g.spawnScreenEntities()
}

func (g *Game) SaveGame() {
	data := &save.SaveData{
		Version:        2,
		HasSword:       g.Player.HasSword,
		MaxHP:          g.Player.MaxHP,
		HP:             g.Player.HP,
		Rupees:         g.Player.Inventory.Rupees,
		Keys:           g.Player.Inventory.Keys,
		Bombs:          g.Player.Inventory.Bombs,
		Arrows:         g.Player.Inventory.Arrows,
		SwordLevel:     g.Player.Inventory.SwordLevel,
		ShieldLevel:    g.Player.Inventory.ShieldLevel,
		BraceletLevel:  g.Player.Inventory.BraceletLevel,
		ButtonA:        int(g.Player.Inventory.ButtonA),
		ButtonB:        int(g.Player.Inventory.ButtonB),
		CollectedItems: g.CollectedItems,
		UnlockedDoors:  g.UnlockedDoors,
		ScreenX:        g.Overworld.CurrentX,
		ScreenY:        g.Overworld.CurrentY,
		PlayerX:        g.Player.X,
		PlayerY:        g.Player.Y,
		InInterior:     g.InInterior,
		BossDefeated:   g.BossDefeated,
	}

	// Save owned items as int slice
	for id := range g.Player.Inventory.OwnedItems {
		data.OwnedItems = append(data.OwnedItems, int(id))
	}

	// Save quest state
	data.Quest = &save.QuestSaveData{
		Flags:              g.Quest.Flags,
		DungeonsCompleted:  g.Quest.DungeonsCompleted,
		TradingItem:        g.Quest.TradingItem,
	}

	if g.InInterior && g.CurrentInterior != nil {
		data.InteriorID = g.CurrentInterior.ID
	}
	save.Save(data)
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

func (g *Game) ToggleMute()  { g.Audio.ToggleMute() }
func (g *Game) VolumeUp()    { g.Audio.VolumeUp() }
func (g *Game) VolumeDown()  { g.Audio.VolumeDown() }

func (g *Game) Update(dt float64) {
	switch g.State {
	case StateMenu:
		g.updateMenu()
	case StatePaused:
		g.updatePaused()
	case StateGameOver:
		g.updateGameOver(dt)
	case StateVictory:
		g.updateVictory(dt)
	case StateDialogue:
		g.updateDialogue()
	case StateInventory:
		g.updateInventory()
	case StatePlaying:
		g.updatePlaying(dt)
	}
	g.Input.Update()
}

func (g *Game) updateMenu() {
	if g.Input.JustPressed(glow.KeyUp) || g.Input.JustPressed(glow.KeyW) {
		g.Menu.MoveUp()
		g.Audio.PlayMenuSelect()
	}
	if g.Input.JustPressed(glow.KeyDown) || g.Input.JustPressed(glow.KeyS) {
		g.Menu.MoveDown()
		g.Audio.PlayMenuSelect()
	}
	if g.Input.JustPressed(glow.KeyEnter) || g.Input.JustPressed(glow.KeySpace) {
		opt := g.Menu.Options[g.Menu.SelectedIndex]
		if opt.Disabled {
			return
		}
		g.Audio.PlayMenuSelect()
		switch g.Menu.SelectedIndex {
		case 0:
			g.StartNewGame()
		case 1:
			g.StartContinue()
		}
	}
}

func (g *Game) updatePaused() {
	if g.Input.JustPressed(glow.KeyEscape) || g.Input.JustPressed(glow.KeyEnter) {
		g.State = StatePlaying
	}
}

func (g *Game) updateGameOver(dt float64) {
	g.GameOverTimer += dt
	if g.GameOverTimer >= config.GameOverDelay {
		if g.Input.JustPressed(glow.KeyEnter) || g.Input.JustPressed(glow.KeySpace) {
			g.State = StateMenu
			g.Menu = NewMenuState(save.Exists())
			g.GameOverTimer = 0
		}
	}
}

func (g *Game) updateVictory(dt float64) {
	g.VictoryTimer += dt
	if g.VictoryTimer >= config.VictoryDelay {
		if g.Input.JustPressed(glow.KeyEnter) || g.Input.JustPressed(glow.KeySpace) {
			g.State = StateMenu
			g.Menu = NewMenuState(save.Exists())
			g.VictoryTimer = 0
		}
	}
}

func (g *Game) updateDialogue() {
	if g.Input.JustPressed(glow.KeySpace) || g.Input.JustPressed(glow.KeyJ) {
		done := g.Dialogue.Advance()
		if done {
			g.State = StatePlaying
		}
	}
}

func (g *Game) updateInventory() {
	if g.Input.JustPressed(glow.KeyTab) || g.Input.JustPressed(glow.KeyEscape) {
		g.State = StatePlaying
	}

	// Navigate inventory grid
	if g.Input.JustPressed(glow.KeyUp) || g.Input.JustPressed(glow.KeyW) {
		g.InventoryCursorY--
		if g.InventoryCursorY < 0 {
			g.InventoryCursorY = 0
		}
	}
	if g.Input.JustPressed(glow.KeyDown) || g.Input.JustPressed(glow.KeyS) {
		g.InventoryCursorY++
		if g.InventoryCursorY > 2 {
			g.InventoryCursorY = 2
		}
	}
	if g.Input.JustPressed(glow.KeyLeft) || g.Input.JustPressed(glow.KeyA) {
		g.InventoryCursorX--
		if g.InventoryCursorX < 0 {
			g.InventoryCursorX = 0
		}
	}
	if g.Input.JustPressed(glow.KeyRight) || g.Input.JustPressed(glow.KeyD) {
		g.InventoryCursorX++
		if g.InventoryCursorX > 4 {
			g.InventoryCursorX = 4
		}
	}

	// Assign to A (Z key) or B (X key)
	idx := g.InventoryCursorY*5 + g.InventoryCursorX
	items := g.Player.Inventory.OwnedItemsList()
	if idx < len(items) {
		item := items[idx]
		if g.Input.JustPressed(glow.KeyJ) {
			g.Player.Inventory.ButtonA = item
		}
		if g.Input.JustPressed(glow.KeyK) {
			g.Player.Inventory.ButtonB = item
		}
	}
}

func (g *Game) updatePlaying(dt float64) {
	// Check pause (Enter)
	if g.Input.JustPressed(glow.KeyEnter) {
		g.State = StatePaused
		return
	}

	// Check inventory (Tab)
	if g.Input.JustPressed(glow.KeyTab) {
		g.State = StateInventory
		return
	}

	// Handle active transition
	if g.Transition.Active {
		g.Transition.Timer += dt
		if g.Transition.Done() {
			g.Transition.Active = false
			if g.Transition.Type == TransitionFade {
				g.completeFadeTransition()
			}
			g.Transition.OldScreen = nil
		}
		return
	}

	// Update polish timers
	if g.ShakeTimer > 0 {
		g.ShakeTimer -= dt
		if g.ShakeTimer < 0 {
			g.ShakeTimer = 0
		}
	}
	if g.FlashTimer > 0 {
		g.FlashTimer -= dt
		if g.FlashTimer < 0 {
			g.FlashTimer = 0
		}
	}

	// Update particles
	g.Particles.Update(dt)

	// Update sword swing
	g.Player.Sword.Update(dt)

	// Update invincibility timer
	if g.Player.InvTimer > 0 {
		g.Player.InvTimer -= dt
		if g.Player.InvTimer < 0 {
			g.Player.InvTimer = 0
		}
	}

	// Space = interact (talk, read signs, open chests, use doors)
	if g.Input.JustPressed(glow.KeySpace) {
		if g.tryInteractNPC() {
			return
		}
		if g.tryInteractDoor() {
			return
		}
	}

	// Z = A button (use equipped item)
	if g.Input.JustPressed(glow.KeyJ) && !g.Player.Sword.Active {
		g.useEquippedItem(g.Player.Inventory.ButtonA)
	}

	// X = B button (use equipped item)
	if g.Input.JustPressed(glow.KeyK) && !g.Player.Sword.Active {
		g.useEquippedItem(g.Player.Inventory.ButtonB)
	}

	// Combat: check sword hits
	if g.Player.Sword.Active {
		hitEnemies := system.CheckSwordHits(g.Player, g.Enemies)
		for _, e := range hitEnemies {
			e.HP--
			e.InvTimer = config.EnemyInvTime
			system.ApplyKnockback(e, g.Player.CenterX(), g.Player.CenterY())
			if e.HP <= 0 {
				e.Dead = true
				g.Audio.PlayEnemyDie()
				g.tryDropItem(e)
				g.spawnDeathParticles(e)
				// Check boss death
				if e.Type == entity.EnemyBoss {
					g.BossDefeated = true
					g.State = StateVictory
					g.VictoryTimer = 0
					g.SaveGame()
					return
				}
			} else {
				g.Audio.PlayEnemyHit()
			}
		}
	}

	// Destroy projectiles deflected by sword
	for _, proj := range g.Projectiles {
		if !proj.Dead && proj.FromEnemy && system.CheckProjectileSwordCollision(g.Player, proj) {
			proj.Dead = true
		}
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
		screen := g.currentScreen()
		crossX, crossY := system.MovePlayer(g.Player, screen, dx, dy, dt)
		if !g.InInterior {
			g.handleEdgeCrossing(crossX, crossY)
		} else {
			g.clampPlayer()
		}
	}

	// Check door entry (standing on door/stairs tile)
	g.checkDoorEntry()

	// Update enemies
	g.updateEnemies(dt)

	// Update projectiles
	g.updateProjectiles(dt)

	// Update items
	g.updateItems(dt)

	// Check item pickup
	g.checkItemPickup()

	// Enemy→player collision
	g.checkEnemyCollisions()

	// Projectile→player collision
	g.checkProjectileCollisions()

	g.Player.UpdateAnimation(dt)
}

// useEquippedItem activates the item assigned to a button.
func (g *Game) useEquippedItem(item entity.EquipItemID) {
	switch item {
	case entity.EquipSword:
		if g.Player.HasSword || g.Player.Inventory.SwordLevel > 0 {
			g.Player.Sword.Start(g.Player.Dir)
			g.Audio.PlaySwordSwing()
		}
	// Other items will be implemented in later phases
	case entity.EquipNone:
		// nothing
	}
}

func (g *Game) handleEdgeCrossing(crossX, crossY int) {
	dirX, dirY := 0, 0
	if crossX != 0 {
		dirX = crossX
	} else if crossY != 0 {
		dirY = crossY
	} else {
		return
	}

	if !g.Overworld.CanMove(dirX, dirY) {
		g.clampPlayer()
		return
	}

	oldScreen := g.Overworld.CurrentScreen()
	g.Overworld.Move(dirX, dirY)

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

	g.Transition.Start(dirX, dirY, oldScreen)
	g.spawnScreenEntities()
	g.SaveGame()
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

func (g *Game) currentScreen() *world.Screen {
	if g.InInterior && g.CurrentInterior != nil {
		return g.CurrentInterior.Screen
	}
	return g.Overworld.CurrentScreen()
}

func (g *Game) spawnScreenEntities() {
	g.Enemies = nil
	g.Projectiles = nil
	g.Items = nil
	g.NPCs = nil

	var screen *world.Screen
	var screenKey string

	if g.InInterior && g.CurrentInterior != nil {
		screen = g.CurrentInterior.Screen
		screenKey = g.CurrentInterior.ID

		for _, es := range g.CurrentInterior.EnemySpawns {
			e := spawnEnemy(es)
			g.Enemies = append(g.Enemies, e)
		}
		for i, is := range g.CurrentInterior.ItemSpawns {
			key := fmt.Sprintf("int_%s_%d", screenKey, i)
			if g.CollectedItems[key] {
				continue
			}
			item := entity.NewItem(entity.ItemType(is.Type),
				float64(is.TileX*config.TileSize)+2,
				float64(is.TileY*config.TileSize)+2)
			g.Items = append(g.Items, item)
		}
		for _, ns := range g.CurrentInterior.NPCSpawns {
			npc := entity.NewNPC(
				ns.ID,
				float64(ns.TileX*config.TileSize)+1,
				float64(ns.TileY*config.TileSize)+1,
				entity.Direction(ns.Dir),
				ns.Name,
				ns.Dialogue,
				ns.ConditionalDialogues,
			)
			g.NPCs = append(g.NPCs, npc)
		}
		_ = screen
		return
	}

	screen = g.Overworld.CurrentScreen()
	screenKey = fmt.Sprintf("%d,%d", g.Overworld.CurrentX, g.Overworld.CurrentY)

	// Apply unlocked doors
	for doorKey := range g.UnlockedDoors {
		var sx, sy, tx, ty int
		if _, err := fmt.Sscanf(doorKey, "%d,%d_%d,%d", &sx, &sy, &tx, &ty); err == nil {
			if sx == g.Overworld.CurrentX && sy == g.Overworld.CurrentY {
				if tx >= 0 && tx < config.ScreenGridW && ty >= 0 && ty < config.ScreenGridH {
					screen.Tiles[ty][tx] = world.TileDoorOpen
				}
			}
		}
	}

	for _, es := range screen.EnemySpawns {
		e := spawnEnemy(es)
		g.Enemies = append(g.Enemies, e)
	}

	for i, is := range screen.ItemSpawns {
		key := fmt.Sprintf("%s_%d", screenKey, i)
		if g.CollectedItems[key] {
			continue
		}
		item := entity.NewItem(entity.ItemType(is.Type),
			float64(is.TileX*config.TileSize)+2,
			float64(is.TileY*config.TileSize)+2)
		g.Items = append(g.Items, item)
	}

	for _, ns := range screen.NPCSpawns {
		npc := entity.NewNPC(
			ns.ID,
			float64(ns.TileX*config.TileSize)+1,
			float64(ns.TileY*config.TileSize)+1,
			entity.Direction(ns.Dir),
			ns.Name,
			ns.Dialogue,
			ns.ConditionalDialogues,
		)
		g.NPCs = append(g.NPCs, npc)
	}
}

func spawnEnemy(es world.EnemySpawn) *entity.Enemy {
	x := float64(es.TileX*config.TileSize) + 1
	y := float64(es.TileY*config.TileSize) + 1
	switch entity.EnemyType(es.Type) {
	case entity.EnemyMoblin:
		return entity.NewMoblin(x, y)
	case entity.EnemyStalfos:
		return entity.NewStalfos(x, y)
	case entity.EnemyBoss:
		return entity.NewBoss(x, y)
	default:
		return entity.NewOctorok(x, y)
	}
}

func (g *Game) updateEnemies(dt float64) {
	screen := g.currentScreen()
	for _, e := range g.Enemies {
		if e.Dead {
			continue
		}
		proj := system.UpdateEnemyAI(e, g.Player, screen, dt, g.RNG)
		if proj != nil {
			g.Projectiles = append(g.Projectiles, proj)
		}
	}
}

func (g *Game) updateProjectiles(dt float64) {
	alive := g.Projectiles[:0]
	for _, p := range g.Projectiles {
		p.Update(dt)
		screen := g.currentScreen()
		if system.TileCollision(screen, p.X, p.Y, p.Width, p.Height) {
			p.Dead = true
		}
		if !p.Dead {
			alive = append(alive, p)
		}
	}
	g.Projectiles = alive
}

func (g *Game) updateItems(dt float64) {
	for _, item := range g.Items {
		if !item.Collected {
			item.Update(dt)
		}
	}
}

func (g *Game) checkItemPickup() {
	px, py, pw, ph := g.Player.BBox()

	var screenKey string
	if g.InInterior && g.CurrentInterior != nil {
		screenKey = "int_" + g.CurrentInterior.ID
	} else {
		screenKey = fmt.Sprintf("%d,%d", g.Overworld.CurrentX, g.Overworld.CurrentY)
	}

	for i, item := range g.Items {
		if item.Collected {
			continue
		}
		if system.AABBOverlap(px, py, pw, ph, item.X, item.Y, float64(item.Width), float64(item.Height)) {
			item.Collected = true
			key := fmt.Sprintf("%s_%d", screenKey, i)
			g.CollectedItems[key] = true
			g.applyItemEffect(item)
			g.Audio.PlayItemPickup()
			g.FlashTimer = config.FlashDuration
			g.SaveGame()
		}
	}
}

func (g *Game) applyItemEffect(item *entity.Item) {
	switch item.Type {
	case entity.ItemHeart:
		g.Player.HP += 2
		if g.Player.HP > g.Player.MaxHP {
			g.Player.HP = g.Player.MaxHP
		}
	case entity.ItemRupee:
		g.Player.Inventory.Rupees++
	case entity.ItemKey:
		g.Player.Inventory.Keys++
	case entity.ItemSword:
		g.Player.HasSword = true
		g.Player.Inventory.OwnedItems[entity.EquipSword] = true
		if g.Player.Inventory.SwordLevel == 0 {
			g.Player.Inventory.SwordLevel = 1
		}
		// Auto-assign sword to A if nothing assigned
		if g.Player.Inventory.ButtonA == entity.EquipNone {
			g.Player.Inventory.ButtonA = entity.EquipSword
		}
		g.Quest.SetFlag("got_sword")
	case entity.ItemHeartContainer:
		g.Player.MaxHP += 2
		g.Player.HP = g.Player.MaxHP
	}
}

func (g *Game) tryDropItem(e *entity.Enemy) {
	roll := g.RNG.Next() % 100
	if roll >= 50 {
		return
	}
	var typ entity.ItemType
	if roll < 25 {
		typ = entity.ItemHeart
	} else {
		typ = entity.ItemRupee
	}
	item := entity.NewItem(typ, e.X, e.Y)
	g.Items = append(g.Items, item)
}

func (g *Game) damagePlayer(amount int) {
	g.Player.HP -= amount
	g.Audio.PlayPlayerHit()
	g.ShakeTimer = config.ShakeDuration
	if g.Player.HP <= 0 {
		g.Player.HP = 0
		g.State = StateGameOver
		g.GameOverTimer = 0
		g.Audio.PlayGameOver()
		return
	}
	g.Player.InvTimer = config.PlayerInvTime
}

func (g *Game) checkEnemyCollisions() {
	if g.Player.InvTimer > 0 {
		return
	}
	for _, e := range g.Enemies {
		if e.Dead {
			continue
		}
		if system.CheckEnemyPlayerCollision(g.Player, e) {
			g.damagePlayer(1)
			return
		}
	}
}

func (g *Game) checkProjectileCollisions() {
	if g.Player.InvTimer > 0 {
		return
	}
	for _, proj := range g.Projectiles {
		if proj.Dead || !proj.FromEnemy {
			continue
		}
		if system.CheckProjectilePlayerCollision(g.Player, proj) {
			proj.Dead = true
			g.damagePlayer(proj.Damage)
			return
		}
	}
}

// getActiveDialogue returns the appropriate dialogue lines for an NPC,
// checking conditional dialogues first, then falling back to default.
func (g *Game) getActiveDialogue(npc *entity.NPC) []string {
	for _, d := range npc.Dialogues {
		if CheckCondition(d.Condition, g.Quest, &g.Player.Inventory) {
			return d.Lines
		}
	}
	return npc.Dialogue
}

func (g *Game) tryInteractNPC() bool {
	for _, npc := range g.NPCs {
		if system.ProximityCheck(g.Player.CenterX(), g.Player.CenterY(),
			npc.CenterX(), npc.CenterY(), config.InteractRadius) {
			// Use conditional dialogue
			activeLines := g.getActiveDialogue(npc)
			g.Dialogue.StartWithLines(npc, activeLines)
			g.State = StateDialogue

			// Set quest flags based on NPC ID
			switch npc.ID {
			case "tarin":
				g.Quest.SetFlag("met_tarin")
			case "marin":
				g.Quest.SetFlag("talked_marin")
			case "meowmeow":
				g.Quest.SetFlag("met_meowmeow")
			case "librarian":
				g.Quest.SetFlag("visited_library")
			}

			return true
		}
	}
	return false
}

func (g *Game) tryInteractDoor() bool {
	if g.InInterior {
		return false
	}
	px := int(g.Player.CenterX()) / config.TileSize
	py := int(g.Player.CenterY()) / config.TileSize

	screen := g.Overworld.CurrentScreen()

	tx := px + int(g.Player.Dir.DX())
	ty := py + int(g.Player.Dir.DY())

	if tx >= 0 && tx < config.ScreenGridW && ty >= 0 && ty < config.ScreenGridH {
		if screen.TileAt(tx, ty) == world.TileDoorLocked && g.Player.Inventory.Keys > 0 {
			g.Player.Inventory.Keys--
			screen.Tiles[ty][tx] = world.TileDoorOpen
			doorKey := fmt.Sprintf("%d,%d_%d,%d", g.Overworld.CurrentX, g.Overworld.CurrentY, tx, ty)
			g.UnlockedDoors[doorKey] = true
			g.Audio.PlayDoorOpen()
			g.SaveGame()
			return true
		}
	}
	return false
}

func (g *Game) checkDoorEntry() {
	if g.Transition.Active {
		return
	}

	px := int(g.Player.CenterX()) / config.TileSize
	py := int(g.Player.CenterY()) / config.TileSize

	if g.InInterior {
		screen := g.CurrentInterior.Screen
		tile := screen.TileAt(px, py)

		// Check interior-to-interior door links first
		if tile == world.TileStairs || tile == world.TileDoorOpen {
			for i := range g.CurrentInterior.DoorLinks {
				dl := &g.CurrentInterior.DoorLinks[i]
				if dl.DoorTileX == px && dl.DoorTileY == py {
					target := dl.InteriorID
					// Handle "interior:X" format
					if len(target) > 9 && target[:9] == "interior:" {
						target = target[9:]
					}
					if target == "overworld" {
						// Exit to overworld
						g.PendingExitLink = g.ReturnLink
						g.Transition.StartFade()
						g.Audio.PlayDoorOpen()
						return
					}
					interior, ok := g.Interiors[target]
					if !ok {
						continue
					}
					g.PendingInterior = interior
					g.PendingDoorLink = dl
					g.Transition.StartFade()
					g.Audio.PlayDoorOpen()
					return
				}
			}
		}

		if tile == world.TileDoorOpen || tile == world.TileStairs {
			if g.ReturnLink != nil {
				g.PendingExitLink = g.ReturnLink
				g.Transition.StartFade()
				g.Audio.PlayDoorOpen()
			}
		}
		return
	}

	screen := g.Overworld.CurrentScreen()
	tile := screen.TileAt(px, py)
	if tile != world.TileDoorOpen && tile != world.TileStairs {
		return
	}

	for i := range g.DoorLinks {
		dl := &g.DoorLinks[i]
		if dl.ScreenX == g.Overworld.CurrentX && dl.ScreenY == g.Overworld.CurrentY &&
			dl.DoorTileX == px && dl.DoorTileY == py {
			interior, ok := g.Interiors[dl.InteriorID]
			if !ok {
				continue
			}
			g.PendingInterior = interior
			g.PendingDoorLink = dl
			g.Transition.StartFade()
			g.Audio.PlayDoorOpen()
			return
		}
	}
}

func (g *Game) completeFadeTransition() {
	if g.PendingInterior != nil {
		g.InInterior = true
		g.Location = LocationInterior
		g.CurrentInterior = g.PendingInterior
		g.ReturnLink = g.PendingDoorLink
		g.Player.X = g.PendingDoorLink.SpawnX
		g.Player.Y = g.PendingDoorLink.SpawnY
		g.PendingInterior = nil
		g.PendingDoorLink = nil
		g.spawnScreenEntities()
		g.SaveGame()
	} else if g.PendingExitLink != nil {
		g.InInterior = false
		g.Location = LocationOverworld
		g.CurrentInterior = nil
		g.Player.X = g.PendingExitLink.ExitX
		g.Player.Y = g.PendingExitLink.ExitY
		g.PendingExitLink = nil
		g.ReturnLink = nil
		g.spawnScreenEntities()
		g.SaveGame()
	}
}

func (g *Game) spawnDeathParticles(e *entity.Enemy) {
	var r, g2, b uint8
	switch e.Type {
	case entity.EnemyOctorok:
		r, g2, b = 200, 50, 50
	case entity.EnemyMoblin:
		r, g2, b = 160, 100, 50
	case entity.EnemyStalfos:
		r, g2, b = 180, 180, 180
	case entity.EnemyBoss:
		r, g2, b = 100, 40, 120
	}
	// Generate 6 random velocities via RNG
	velocities := make([]float64, 12)
	for i := 0; i < 12; i++ {
		v := float64(int32(g.RNG.Next()%200)-100) / 2.0
		velocities[i] = v
	}
	g.Particles.SpawnExplosion(e.CenterX(), e.CenterY(), 6, r, g2, b, velocities)
}

func (g *Game) Draw(canvas *glow.Canvas) {
	sc := render.NewScaledCanvas(canvas, g.Layout)
	sc.Clear(render.ColorBG)

	switch g.State {
	case StateMenu:
		opts := make([]render.MenuOption, len(g.Menu.Options))
		for i, o := range g.Menu.Options {
			opts[i] = render.MenuOption{Label: o.Label, Disabled: o.Disabled}
		}
		render.DrawTitleScreen(sc, opts, g.Menu.SelectedIndex)
		return
	case StateVictory:
		render.DrawVictoryScreen(sc, g.VictoryTimer)
		return
	case StateGameOver:
		render.DrawGameOverScreen(sc, g.GameOverTimer)
		return
	}

	// Compute shake offset
	shakeX, shakeY := 0, 0
	if g.ShakeTimer > 0 {
		shakeX = int(int32(g.RNG.Next()%uint32(config.ShakeIntensity*2+1))) - config.ShakeIntensity
		shakeY = int(int32(g.RNG.Next()%uint32(config.ShakeIntensity*2+1))) - config.ShakeIntensity
	}

	if g.Transition.Active {
		switch g.Transition.Type {
		case TransitionScroll:
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
		case TransitionFade:
			screen := g.currentScreen()
			render.DrawScreenAt(sc, screen, shakeX, shakeY)
			g.drawEntities(sc, shakeX, shakeY)
			render.DrawPlayerAt(sc, g.Player, shakeX, shakeY)
			render.DrawParticles(sc, g.Particles.Particles, shakeX, shakeY)
			render.DrawFade(sc, g.Transition.FadeProgress())
		}
	} else {
		screen := g.currentScreen()
		render.DrawScreenAt(sc, screen, shakeX, shakeY)
		g.drawEntities(sc, shakeX, shakeY)
		render.DrawPlayerAt(sc, g.Player, shakeX, shakeY)
		render.DrawParticles(sc, g.Particles.Particles, shakeX, shakeY)
	}

	// Boss health bar
	g.drawBossHealthBar(sc)

	// Flash overlay
	if g.FlashTimer > 0 {
		intensity := g.FlashTimer / config.FlashDuration
		render.DrawFlash(sc, intensity)
	}

	render.DrawHUD(sc, g.Player)

	// Dialogue box on top
	if g.State == StateDialogue && g.Dialogue.Active {
		render.DrawDialogueBox(sc, g.Dialogue.NPC.Name, g.Dialogue.Lines,
			g.Dialogue.CurrentLine, g.Dialogue.HasMore())
	}

	// Pause overlay
	if g.State == StatePaused {
		render.DrawPauseOverlay(sc)
	}

	// Inventory overlay
	if g.State == StateInventory {
		render.DrawInventoryScreen(sc, &g.Player.Inventory, g.InventoryCursorX, g.InventoryCursorY)
	}
}

func (g *Game) drawBossHealthBar(sc *render.ScaledCanvas) {
	for _, e := range g.Enemies {
		if e.Type == entity.EnemyBoss && !e.Dead {
			render.DrawBossHealthBar(sc, e.HP, e.MaxHP)
			return
		}
	}
}

func (g *Game) drawEntities(sc *render.ScaledCanvas, offsetX, offsetY int) {
	for _, item := range g.Items {
		render.DrawItemAt(sc, item, offsetX, offsetY)
	}
	for _, e := range g.Enemies {
		render.DrawEnemyAt(sc, e, offsetX, offsetY)
	}
	for _, p := range g.Projectiles {
		render.DrawProjectileAt(sc, p, offsetX, offsetY)
	}
	for _, npc := range g.NPCs {
		render.DrawNPCAt(sc, npc, offsetX, offsetY)
	}
}
