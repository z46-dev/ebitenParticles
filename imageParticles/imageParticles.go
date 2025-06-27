package imageParticles

import (
	"math/rand/v2"

	"gamedev.z46.dev/ebiten-particles/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

var img *ebiten.Image

func init() {
	img = ebiten.NewImage(32, 32)
	img.Fill(colornames.Red)
}

var Particles *shared.Collection = shared.NewCollection(16384)

func UpdateFunc(width, height float64) {
	Particles.Update()

	for range 192 {
		Particles.Add(shared.NewParticle(width/2, height/2, rand.Float64()*20-10, rand.Float64()*20-10, 8, 120))
	}
}

func DrawFunc(screen *ebiten.Image) {
	var bounds = img.Bounds()
	var dx, dy float64 = float64(bounds.Dx()), float64(bounds.Dy())
	Particles.ForEach(func(p *shared.Particle) {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(p.Size/dx, p.Size/dy)
		op.GeoM.Translate(p.X, p.Y)
		screen.DrawImage(img, op)
	})
}
