package triangleParticles

import (
	"image/color"

	"gamedev.z46.dev/ebiten-particles/shared"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	maxParticles = 128 * 1024
	batchSize    = 4096
)

var (
	Particles          = shared.NewCollection(maxParticles)
	whiteImage         *ebiten.Image
	numBatches         = (maxParticles + batchSize - 1) / batchSize
	verticesBatch      [][]ebiten.Vertex
	indicesBatch       [][]uint16
	screenWf, screenHf float32
)

func init() {
	w, h := ebiten.Monitor().Size()
	Particles.Init(float64(w), float64(h))
	whiteImage = ebiten.NewImage(1, 1)
	whiteImage.Fill(color.White)

	verticesBatch = make([][]ebiten.Vertex, numBatches)
	indicesBatch = make([][]uint16, numBatches)

	for b := 0; b < numBatches; b++ {
		vs := make([]ebiten.Vertex, batchSize*4)
		is := make([]uint16, batchSize*6)

		for i := 0; i < batchSize; i++ {
			vi := 4 * i
			vs[vi+0].SrcX, vs[vi+0].SrcY = 0, 0
			vs[vi+1].SrcX, vs[vi+1].SrcY = 1, 0
			vs[vi+2].SrcX, vs[vi+2].SrcY = 1, 1
			vs[vi+3].SrcX, vs[vi+3].SrcY = 0, 1
			for j := 0; j < 4; j++ {
				vs[vi+j].ColorR = 1
				vs[vi+j].ColorG = 1
				vs[vi+j].ColorB = 1
				vs[vi+j].ColorA = 1
			}

			ii := 6 * i
			base := uint16(4 * i)
			is[ii+0], is[ii+1], is[ii+2] = base+0, base+1, base+2
			is[ii+3], is[ii+4], is[ii+5] = base+0, base+2, base+3
		}

		verticesBatch[b] = vs
		indicesBatch[b] = is
	}
}

func UpdateFunc(sw, sh float64) {
	Particles.Update(sw, sh)
	// for i := 0; i < 1024; i++ {
	// 	Particles.Add(shared.NewParticle(
	// 		sw/2, sh/2, rand.Float64()*20-10, rand.Float64()*20-10,
	// 		8, 120,
	// 	))
	// }
}

func DrawFunc(screen *ebiten.Image) {
	idx := 0
	var op ebiten.DrawTrianglesOptions

	flush := func(b, used int) {
		if used == 0 {
			return
		}

		screen.DrawTriangles(
			verticesBatch[b][:used*4],
			indicesBatch[b][:used*6],
			whiteImage,
			&op,
		)
	}

	batch := 0
	used := 0
	Particles.ForEach(func(p *shared.Particle) {
		if used == batchSize {
			flush(batch, used)
			batch++
			used = 0
		}

		x, y, s := float32(p.X), float32(p.Y), float32(p.Size)
		vi := used * 4
		vs := verticesBatch[batch]

		vs[vi+0].DstX, vs[vi+0].DstY = x, y
		vs[vi+1].DstX, vs[vi+1].DstY = x+s, y
		vs[vi+2].DstX, vs[vi+2].DstY = x+s, y+s
		vs[vi+3].DstX, vs[vi+3].DstY = x, y+s

		used++
		idx++
	})

	flush(batch, used)
}
