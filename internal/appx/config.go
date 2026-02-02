package appx

import (
	// "github.com/tmazitov/42_scop/internal/appx"
)

type Config struct {
	window *WindowOptions
}

func NewConfig() (*Config, error) {
	return &Config{
		window: &WindowOptions{
			Height: 720,
			Width: 1080,
			Title: "SCOP",
		},
	}, nil
}

func (c *Config) Window() *WindowOptions{
	return c.window
}