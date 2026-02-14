package entity

import "github.com/AchrafSoltani/GlowQuest/config"

type SwordSwing struct {
	Active   bool
	Timer    float64
	Duration float64
	Dir      Direction
}

func (s *SwordSwing) Start(dir Direction) {
	s.Active = true
	s.Timer = 0
	s.Duration = config.SwordDuration
	s.Dir = dir
}

func (s *SwordSwing) Update(dt float64) {
	if !s.Active {
		return
	}
	s.Timer += dt
	if s.Timer >= s.Duration {
		s.Active = false
	}
}

func (s *SwordSwing) Done() bool {
	return !s.Active
}

// Progress returns 0→1 how far through the swing we are.
func (s *SwordSwing) Progress() float64 {
	if s.Duration <= 0 {
		return 1
	}
	p := s.Timer / s.Duration
	if p > 1 {
		return 1
	}
	return p
}

// HitBox returns the sword's AABB extending from the player's bounding box.
func (s *SwordSwing) HitBox(playerX, playerY float64, playerW, playerH int) (x, y, w, h float64) {
	pw := float64(playerW)
	ph := float64(playerH)
	reach := float64(config.SwordReach)
	width := float64(config.SwordWidth)

	switch s.Dir {
	case DirUp:
		return playerX + (pw-width)/2, playerY - reach, width, reach
	case DirDown:
		return playerX + (pw-width)/2, playerY + ph, width, reach
	case DirLeft:
		return playerX - reach, playerY + (ph-width)/2, reach, width
	case DirRight:
		return playerX + pw, playerY + (ph-width)/2, reach, width
	}
	return 0, 0, 0, 0
}
