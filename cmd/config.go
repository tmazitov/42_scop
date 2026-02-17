package main

import (
	"strconv"
	"os"
	"errors"
	"path/filepath"
	"github.com/joho/godotenv"
	"github.com/tmazitov/42_scop/internal/appx"
)

type Config struct {
	ObjectPath 		string
	Window *appx.WindowOptions
}

func loadVar(name string, defaultValue string) string {
	envVar := os.Getenv(name)
	if len(envVar) == 0 {
		return defaultValue
	}
	return envVar
}

func SetupConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	var (
		rawSize = [2]string{
			loadVar("WINDOW_HEIGHT", "500"),
			loadVar("WINDOW_WIDTH", "500"),
		}
		size = [2]int {
			0,
			0,
		}
	)

	for index, elem := range rawSize {
		size[index], err = strconv.Atoi(elem)
		if err != nil {
			return nil, errors.New("config error : invalid window size")
		}
	}

	objFile := loadVar("OBJ_FILE_PATH", "")
	if len(objFile) == 0 {
		return nil, errors.New("config error : path to .obj file is not defined")
	}

	objName := filepath.Base(objFile)

	return &Config{
		ObjectPath: objFile,
		Window:  &appx.WindowOptions{
			Title: loadVar("WINDOW_TITLE", "SCOP | " + objName),
			Height: size[0],
			Width: size[1],
		},
	}, nil
}

func (c *Config) ToAppConfig() *appx.Config {
	return &appx.Config{
		Window: c.Window,
	}
}