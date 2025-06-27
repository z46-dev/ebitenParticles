package shared

type Collection struct {
	items      []*Particle
	CachedSize int
}

func NewCollection(max int) *Collection {
	return &Collection{
		items: make([]*Particle, max),
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

func (c *Collection) Update() {
	c.CachedSize = 0
	for i := range c.items {
		if c.items[i] != nil {
			if !c.items[i].Update() {
				c.Remove(c.items[i])
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
