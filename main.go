package main

import (
	"fmt"
	"image/color"

	"gamedev.z46.dev/ebiten-particles/imageParticles"
	"gamedev.z46.dev/ebiten-particles/shaderParticles"
	"gamedev.z46.dev/ebiten-particles/triangleParticles"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/z46-dev/go-logger"
)

type particleMode int

const (
	ModeImageParticles particleMode = iota
	ModeShaderParticles
	ModeTrianglesParticles
)

var ModeNames = map[particleMode]string{
	ModeImageParticles:     "Image Particles",
	ModeShaderParticles:    "Shader Particles",
	ModeTrianglesParticles: "Triangles Particles",
}

var switchString string = ""

func init() {
	for mode, name := range ModeNames {
		switchString += fmt.Sprintf("%s: [%s]\n", name, ebiten.Key(mode).String())
	}
}

type Game struct {
	Mode particleMode
}

func (g *Game) Update() error {
	intWidth, intHeight := ebiten.Monitor().Size()
	floatWidth, floatHeight := float64(intWidth), float64(intHeight)

	switch g.Mode {
	case ModeImageParticles:
		imageParticles.UpdateFunc(floatWidth, floatHeight)
	case ModeShaderParticles:
		shaderParticles.UpdateFunc(floatWidth, floatHeight)
	case ModeTrianglesParticles:
		triangleParticles.UpdateFunc(floatWidth, floatHeight)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	for mode := range ModeNames {
		if inpututil.IsKeyJustPressed(ebiten.Key(mode)) {
			g.Mode = mode
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{1, 1, 1, 1})
	var num int = 0

	switch g.Mode {
	case ModeImageParticles:
		imageParticles.DrawFunc(screen)
		num = imageParticles.Particles.CachedSize
	case ModeShaderParticles:
		shaderParticles.DrawFunc(screen)
		num = shaderParticles.NumParticles
	case ModeTrianglesParticles:
		triangleParticles.DrawFunc(screen)
		num = triangleParticles.Particles.CachedSize
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f | TPS: %.2f\nMode: %s | Particles: %d\n\nSwitch Mode:\n%s", ebiten.ActualFPS(), ebiten.ActualTPS(), ModeNames[g.Mode], num, switchString))
}

func (g *Game) Layout(intWidth, intHeight int) (int, int) {
	return ebiten.Monitor().Size()
}

func main() {
	ebiten.SetWindowSize(ebiten.Monitor().Size())
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowTitle("Particles Benchmark")
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
	ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	ebiten.SetVsyncEnabled(false)
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowDecorated(false)

	var log *logger.Logger = logger.NewLogger().SetPrefix("[MAIN]", logger.BoldCyan).IncludeTimestamp()

	log.Basicf("Startup Debug: Graphics Engine %s\n", func() string {
		var i ebiten.DebugInfo
		ebiten.ReadDebugInfo(&i)

		return i.GraphicsLibrary.String()
	}())

	game := &Game{
		Mode: ModeImageParticles,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Errorf("Failed to run game: %v", err)
		return
	}
}
