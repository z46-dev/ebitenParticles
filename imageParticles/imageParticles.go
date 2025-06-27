package imageParticles

import (
	"gamedev.z46.dev/ebiten-particles/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

var img *ebiten.Image
var Particles *shared.Collection = shared.NewCollection(32768)

func init() {
	w, h := ebiten.Monitor().Size()
	Particles.Init(float64(w), float64(h))
	img = ebiten.NewImage(32, 32)
	img.Fill(colornames.Red)
}

func UpdateFunc(width, height float64) {
	Particles.Update(width, height)
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
