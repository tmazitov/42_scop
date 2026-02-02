package main

import (
	"log"
	"runtime"

	// OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	appx "github.com/tmazitov/42_scop/internal/app"
)

const (
	width  = 500
	height = 500
)

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

var (
	triangle = []float32{
		0.5, 0.5, 0, // left
		-0.5, -0.5, 0, // left
		0.5, -0.5, 0, // right
		-0.5, 0.5, 0, // right
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

	for !app.Window().Core().ShouldClose() {
		vao := makeVao(triangle)
		draw(vao, app)
	}
}

func draw(vao uint32, app *appx.App) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(app.Core())

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

	glfw.PollEvents()
	app.Window().Core().SwapBuffers()
}



// initGlfw initializes glfw and returns a Window to use.
