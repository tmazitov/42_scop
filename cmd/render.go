package main

import (
	"github.com/tmazitov/42_scop/internal/appx"
)

func render(app *appx.App, config *Config) {
	
	for !app.Window().Core().ShouldClose() {
		app.Process()
		app.Draw()
	}
}

