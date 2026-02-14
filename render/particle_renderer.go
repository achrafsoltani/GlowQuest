package render

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
	"github.com/AchrafSoltani/glow"
)

func DrawParticles(sc *ScaledCanvas, particles []*entity.Particle, offsetX, offsetY int) {
	for _, p := range particles {
		alpha := p.Life / p.MaxLife
		r := uint8(float64(p.ColorR) * alpha)
		g := uint8(float64(p.ColorG) * alpha)
		b := uint8(float64(p.ColorB) * alpha)

		px := int(p.X) + offsetX
		py := int(p.Y) + config.HUDHeight + offsetY
		sc.DrawRect(px, py, p.Size, p.Size, glow.RGB(r, g, b))
	}
}
