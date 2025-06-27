package shared

type Particle struct {
	X, Y, Vx, Vy, Size float64
	Lifetime           int
}

func NewParticle(x, y, vx, vy, size float64, lifetime int) *Particle {
	return &Particle{
		X:        x,
		Y:        y,
		Vx:       vx,
		Vy:       vy,
		Size:     size,
		Lifetime: lifetime,
	}
}

func (p *Particle) Update() bool {
	p.Vy += .1
	p.X += p.Vx
	p.Y += p.Vy
	p.Lifetime--
	return p.Lifetime > 0
}
