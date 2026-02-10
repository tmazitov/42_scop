package appx

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func (a *App) Draw(projection mgl32.Mat4) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadMatrixf(&projection[0])

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	
	view := a.Camera().GetViewMatrix()
	gl.MultMatrixf(&view[0])

	for _, obj := range a.Objects() {
		gl.PushMatrix()
		model := mgl32.Ident4()
		gl.MultMatrixf(&model[0])

		gl.BindVertexArray(obj.VAO(a.ScreenSize))
		
		gl.Color3f(1.0, 1.0, 1.0)
		
		gl.DrawElements(gl.TRIANGLES, int32(obj.IndicesCount()), gl.UNSIGNED_INT, nil)

		gl.PopMatrix()
	}

	// Swap buffers
	a.Window().Core().SwapBuffers()
	glfw.PollEvents()
}