package entity

type Particle struct {
	X, Y      float64
	VX, VY    float64
	Life      float64
	MaxLife   float64
	Size      int
	ColorR    uint8
	ColorG    uint8
	ColorB    uint8
}

type ParticlePool struct {
	Particles []*Particle
}

func NewParticlePool() *ParticlePool {
	return &ParticlePool{}
}

func (pp *ParticlePool) SpawnExplosion(x, y float64, count int, r, g, b uint8, velocities []float64) {
	for i := 0; i < count && i*2+1 < len(velocities); i++ {
		p := &Particle{
			X:      x,
			Y:      y,
			VX:     velocities[i*2],
			VY:     velocities[i*2+1],
			Life:   0.5,
			MaxLife: 0.5,
			Size:   2,
			ColorR: r,
			ColorG: g,
			ColorB: b,
		}
		pp.Particles = append(pp.Particles, p)
	}
}

func (pp *ParticlePool) Update(dt float64) {
	alive := pp.Particles[:0]
	for _, p := range pp.Particles {
		p.Life -= dt
		if p.Life <= 0 {
			continue
		}
		p.X += p.VX * dt
		p.Y += p.VY * dt
		alive = append(alive, p)
	}
	pp.Particles = alive
}
