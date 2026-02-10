package main

import (
	"log"
	"runtime"

	// OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tmazitov/42_scop/internal/appx"
	"github.com/tmazitov/42_scop/internal/parsing"
	"github.com/go-gl/mathgl/mgl32"
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

	obj, err := parsing.ParseObj("./resources/teapot.obj")
	if err != nil {
		log.Fatal(err)
	}
	app.AddObject(obj)
	modelLoc, viewLoc, projectionLoc := appx.GetUniformLocations(app.Core())

	projection := mgl32.Perspective(
		mgl32.DegToRad(45.0),           // field of view
		float32(config.Window().Width)/float32(config.Window().Height), // aspect ratio
		0.1,                             // near plane
		100.0,                           // far plane
	)

	for !app.Window().Core().ShouldClose() {
		draw(app, modelLoc, viewLoc, projectionLoc, projection)
	}
}

func draw(app *appx.App, modelLoc, viewLoc, projectionLoc int32, projection mgl32.Mat4) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(app.Core())

	app.Process()

	// Set up matrices
    model := mgl32.Ident4() // Identity matrix (no transformation)
    view := app.Camera().GetViewMatrix()

    // Send matrices to shaders
    gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])
    gl.UniformMatrix4fv(viewLoc, 1, false, &view[0])
    gl.UniformMatrix4fv(projectionLoc, 1, false, &projection[0])

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
