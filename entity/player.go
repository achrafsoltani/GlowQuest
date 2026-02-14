package entity

import "github.com/AchrafSoltani/GlowQuest/config"

type Player struct {
	X, Y      float64
	Width     int
	Height    int
	Dir       Direction
	Speed     float64
	HP, MaxHP int
	Moving    bool
	WalkFrame int
	WalkTimer float64
	Sword     SwordSwing
	InvTimer  float64
	HasSword  bool
	Inventory Inventory
}

func NewPlayer(x, y float64) *Player {
	return &Player{
		X:      x,
		Y:      y,
		Width:  config.PlayerSize,
		Height: config.PlayerSize,
		Dir:    DirDown,
		Speed:  config.PlayerSpeed,
		HP:     config.MaxHP,
		MaxHP:  config.MaxHP,
	}
}

func (p *Player) CenterX() float64 { return p.X + float64(p.Width)/2 }
func (p *Player) CenterY() float64 { return p.Y + float64(p.Height)/2 }

func (p *Player) BBox() (float64, float64, float64, float64) {
	return p.X, p.Y, float64(p.Width), float64(p.Height)
}

func (p *Player) UpdateAnimation(dt float64) {
	if p.Moving {
		p.WalkTimer += dt
		if p.WalkTimer >= config.WalkFrameTime {
			p.WalkTimer -= config.WalkFrameTime
			p.WalkFrame = (p.WalkFrame + 1) % config.WalkFrames
		}
	} else {
		p.WalkFrame = 0
		p.WalkTimer = 0
	}
}
