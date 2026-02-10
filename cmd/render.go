package main

import (
	"github.com/tmazitov/42_scop/internal/appx"
	"github.com/tmazitov/42_scop/internal/rende"
)

func render(app *appx.App, config *Config) {
	projection := rende.MakeProjection(app.ScreenSize.Height, app.ScreenSize.Width)

	for !app.Window().Core().ShouldClose() {
		app.Process()
		app.Draw(projection)
	}
}

