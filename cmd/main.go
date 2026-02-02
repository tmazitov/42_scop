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

	obj := rende.NewObject("test", rende.NewPoint(0, 0, 0), []*rende.Point{
		rende.NewPoint(200, 180, 0), // A
		rende.NewPoint(810, 180, 0), // B
		rende.NewPoint(270, 560, 0), // D
		rende.NewPoint(810, 560, 0), // C
	})

	app.AddObject(obj)

	for !app.Window().Core().ShouldClose() {
		draw(app)
	}
}

func draw(app *appx.App) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(app.Core())

	for _, obj := range app.Objects() {
		gl.BindVertexArray(obj.VAO(app.ScreenSize))
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
		gl.DrawElements(gl.TRIANGLES, obj.IndicesCount(), gl.UNSIGNED_INT, gl.PtrOffset(0))
		gl.BindVertexArray(0)
	}

	glfw.PollEvents()
	app.Window().Core().SwapBuffers()
}



// initGlfw initializes glfw and returns a Window to use.
