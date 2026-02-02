package main

import (
	"log"
	"runtime"

	// OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tmazitov/42_scop/internal/appx"
	"github.com/tmazitov/42_scop/internal/rende"
)

// makeVao initializes and returns a vertex array from the points provided.


var (
	triangle = []*rende.Point{
		rende.NewPoint(270, 180, 0),
		rende.NewPoint(810, 180, 0),
		rende.NewPoint(810, 560, 0),
	}
)

func main() {
	runtime.LockOSThread()

	config, err := appx.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	app, err := appx.NewApp(config)
	if err != nil {
		log.Fatal(err)
	}
	defer app.Close()

	screenSize := rende.ScreenSize{
		Height: float32(config.Window().Height),
		Width: float32(config.Window().Width),
	}
	vao := rende.MakeVao(screenSize, triangle)
	log.Println("vaoted")
	for !app.Window().Core().ShouldClose() {
		draw(vao, app)
	}
}

func draw(vao uint32, app *appx.App) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(app.Core())

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)))

	glfw.PollEvents()
	app.Window().Core().SwapBuffers()
}



// initGlfw initializes glfw and returns a Window to use.
