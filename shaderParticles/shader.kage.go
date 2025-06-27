//go:build ignore

//kage:unit pixels

package main

const MAX_PARTICLES = 8

var ScreenSize vec2
var NumParticles int
var ParticleData [MAX_PARTICLES]vec3 // x, y, size
var ParticleColor vec4

func Fragment(dstPos vec4, srcPos vec2, _ vec4) vec4 {
	for i := 0; i < MAX_PARTICLES; i++ {
		if i >= NumParticles {
			break
		}

        if abs(dstPos.x-ParticleData[i].x) <= ParticleData[i].z &&
           abs(dstPos.y-ParticleData[i].y) <= ParticleData[i].z {
            return ParticleColor
        }
	}

	return vec4(0.0, 0.0, 0.0, 0.0)
}
