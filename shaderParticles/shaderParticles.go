package shaderParticles

import (
	"math/rand/v2"

	"gamedev.z46.dev/ebiten-particles/shared"
	"github.com/hajimehoshi/ebiten/v2"
)

func UpdateFunc(width, height float64) {
	updateParticles()

	for range 8 {
		addParticle(shared.NewParticle(width/2, height/2, rand.Float64()*20-10, rand.Float64()*20-10, 2, 120))
	}
}

func DrawFunc(screen *ebiten.Image) {
	drawBatches(screen)
}
