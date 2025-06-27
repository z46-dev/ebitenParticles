package config

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type Configuration struct {
	ImageParticles struct {
		MaxParticles     int `env:"IMAGE_PARTICLES_MAX_PARTICLES,default=16384"`
		ParticlesPerTick int `env:"IMAGE_PARTICLES_PARTICLES_PER_TICK,default=192"`
	}

	ShaderParticles struct {
		MaxParticles     int `env:"SHADER_PARTICLES_MAX_PARTICLES,default=128"`
		ParticlesPerTick int `env:"SHADER_PARTICLES_PARTICLES_PER_TICK,default=4"`
		BatchSize        int `env:"SHADER_PARTICLES_BATCH_SIZE,default=8"`
	}
}

var Config Configuration

// Try to initialize the environment variables from a .env in the directory the program is run from.
// If the .env file is not present, we will create a sample .env file based on the Configuration struct.
// You can then use config.Config globally
func InitEnv(path string) error {
	if _, err := os.Stat(path); err != nil {
		if e := GenerateSampleEnvFile(path); e != nil {
			return e
		}

		return fmt.Errorf("no .env file found, created a sample .env file. Please fill in the required values and try again")
	}

	if err := godotenv.Load(path); err != nil {
		return err
	}

	_, err := env.UnmarshalFromEnviron(&Config)
	if err != nil {
		return err
	}

	// This is annoying, but we need to ensure the batch size lines up within the kage shader too!
	const SHADER_PATH = "shaderParticles/shader.kage.go"
	if file, err := os.Open(SHADER_PATH); err == nil {
		defer file.Close()

		if data, err := os.ReadFile(SHADER_PATH); err == nil {
			lines := strings.Split(string(data), "\n")
			for _, line := range lines {
				if strings.Contains(line, "const MAX_PARTICLES =") {
					parts := strings.Split(line, "=")
					if len(parts) != 2 {
						return fmt.Errorf("failed to parse MAX_PARTICLES line in shader: %s", line)
					}
					currentValue := strings.TrimSpace(parts[1])
					if currentValue != fmt.Sprintf("%d", Config.ShaderParticles.BatchSize) {
						return fmt.Errorf("batch size in shader (%s) does not match config (%d)", currentValue, Config.ShaderParticles.BatchSize)
					}
				}
			}
		}
	} else {
		return fmt.Errorf("failed to open shader file: %w", err)
	}

	return nil
}
