package shared

import "math/rand"

type Collection struct {
	items      []*Particle
	CachedSize int
}

func NewCollection(max int) *Collection {
	return &Collection{
		items: make([]*Particle, max),
	}
}

func (c *Collection) Init(w, h float64) {
	for i := range len(c.items) {
		c.items[i] = NewParticle(rand.Float64()*w, rand.Float64()*h,
			rand.Float64()*20-10, rand.Float64()*20-10,
			8, 120,
		)
	}
}

func (c *Collection) Add(item *Particle) {
	for i := range c.items {
		if c.items[i] == nil {
			c.items[i] = item
			return
		}
	}
}

func (c *Collection) Remove(item *Particle) {
	for i := range c.items {
		if c.items[i] == item {
			c.items[i] = nil
			return
		}
	}
}

func (c *Collection) Reset(item *Particle, w, h float64) {
	item.X = w / 2
	item.Y = h / 2
	item.Vx = rand.Float64()*20 - 10
	item.Vy = rand.Float64()*20 - 10
	item.Size = 8
	item.Lifetime = 60 + rand.Intn(120)
}

func (c *Collection) Update(w, h float64) {
	c.CachedSize = 0
	for i := range c.items {
		if c.items[i] != nil {
			if !c.items[i].Update() {
				c.Reset(c.items[i], w, h)
			} else {
				c.CachedSize++
			}
		}
	}
}

func (c *Collection) ForEach(fn func(*Particle)) {
	for i := range c.items {
		if c.items[i] != nil {
			fn(c.items[i])
		}
	}
}
