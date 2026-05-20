package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/tmazitov/42_scop/internal/appx"
	"github.com/tmazitov/42_scop/internal/appx/window"
)

type Config struct {
	ObjectPath    string
	Window        *window.WindowOptions
	RotationSpeed float32
}

func loadVar(name string, defaultValue string) string {
	envVar := os.Getenv(name)
	if len(envVar) == 0 {
		return defaultValue
	}
	return envVar
}

func SetupConfig() (*Config, error) {

	if len(os.Args) < 2 {
		return nil, fmt.Errorf("usage: %s <path/to/file.obj>", os.Args[0])
	}
	objFile := os.Args[1]
	if filepath.Ext(objFile) != ".obj" {
		return nil, errors.New("config error : file must have .obj extension")
	}

	if err := godotenv.Load(); err != nil {
		log.Printf("config warn : .env not found, using defaults\n")
	}

	var (
		err     error
		rawSize = [2]string{
			loadVar("WINDOW_HEIGHT", "720"),
			loadVar("WINDOW_WIDTH", "1080"),
		}
		size = [2]int{}
	)

	for index, elem := range rawSize {
		size[index], err = strconv.Atoi(elem)
		if err != nil {
			return nil, errors.New("config error : invalid window size")
		}
	}

	rotationSpeed := loadVar("ROTATION_SPEED", "0.1")
	rotationSpeedFloat, err := strconv.ParseFloat(rotationSpeed, 32)
	if err != nil {
		return nil, errors.New("config error : invalid rotation speed")
	}

	objName := filepath.Base(objFile)

	return &Config{
		ObjectPath:    objFile,
		RotationSpeed: float32(rotationSpeedFloat),
		Window: &window.WindowOptions{
			Title:  loadVar("WINDOW_TITLE", "SCOP | "+objName),
			Height: size[0],
			Width:  size[1],
		},
	}, nil
}

func (c *Config) ToAppConfig() *appx.Config {
	return &appx.Config{
		Window: c.Window,
		RotationSpeed: c.RotationSpeed,
	}
}