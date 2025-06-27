package shaderParticles

import (
	_ "embed"
	"math/rand/v2"

	"gamedev.z46.dev/ebiten-particles/shared"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed shader.kage.go
var ParticleShaderSource []byte
var ParticleShader *ebiten.Shader

func init() {
	var err error
	ParticleShader, err = ebiten.NewShader(ParticleShaderSource)
	if err != nil {
		panic(err)
	}
}

const PARTICLE_BATCH_SIZE = 8
const MAX_PARTICLES = 1024

var NumParticles int = 0

var particles [MAX_PARTICLES / PARTICLE_BATCH_SIZE][PARTICLE_BATCH_SIZE]*shared.Particle
var colors [MAX_PARTICLES / PARTICLE_BATCH_SIZE][4]float32
var batchesToDraw []struct {
	Index, Count int
}

func addParticle(particle *shared.Particle) {
	for i := range particles {
		for j := range particles[i] {
			if particles[i][j] == nil {
				particles[i][j] = particle
				NumParticles++

				if colors[i][3] == 0 {
					colors[i] = [4]float32{rand.Float32()*.5 + .25, rand.Float32()*.5 + .25, rand.Float32()*.5 + .25, 1}
				}
				return
			}
		}
	}
}

func drawParticles(dstImg *ebiten.Image, subList [PARTICLE_BATCH_SIZE]*shared.Particle, batchColor [4]float32) {
	bounds := dstImg.Bounds()
	dx, dy := bounds.Dx(), bounds.Dy()

	flatVec3s := make([]float32, PARTICLE_BATCH_SIZE*3)
	num := 0
	for i := range subList {
		if subList[i] == nil {
			continue
		}

		flatVec3s[num*3] = float32(subList[i].X)
		flatVec3s[num*3+1] = float32(subList[i].Y)
		flatVec3s[num*3+2] = float32(subList[i].Size)
		num++
	}

	dstImg.DrawRectShader(dx, dy, ParticleShader, &ebiten.DrawRectShaderOptions{
		Uniforms: map[string]any{
			"ScreenSize":    []float32{float32(dx), float32(dy)},
			"NumParticles":  num,
			"ParticleData":  flatVec3s,
			"ParticleColor": batchColor,
		},
	})
}

func updateParticles() {
	for i := range particles {
		num := 0

		for j := range particles[i] {
			if particles[i][j] == nil {
				continue
			}

			if !particles[i][j].Update() {
				particles[i][j] = nil
				NumParticles--
			} else {
				num++
			}
		}

		if num > 0 {
			// drawParticles(dstImg, particles[i], colors[i])
			batchesToDraw = append(batchesToDraw, struct {
				Index int
				Count int
			}{
				Index: i,
				Count: num,
			})
		}
	}
}

func drawBatches(dstImg *ebiten.Image) {
	for _, batch := range batchesToDraw {
		drawParticles(dstImg, particles[batch.Index], colors[batch.Index])
	}

	batchesToDraw = nil // Clear the batches for the next frame
}
