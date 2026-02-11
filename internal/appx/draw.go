package appx

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func (a *App) Draw(projection mgl32.Mat4) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    
    a.DrawScene(projection)
    
    a.DrawUI()
    
    a.Window().Core().SwapBuffers()
    glfw.PollEvents()
}

func (a *App) DrawScene(projection mgl32.Mat4) {

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
		
		// gl.Color3f(1.0, 1.0, 1.0)
		
		gl.DrawElements(gl.TRIANGLES, int32(obj.IndicesCount()), gl.UNSIGNED_INT, nil)

		gl.PopMatrix()
	}

}

func (a *App) DrawUI() {
	// Disable depth test so UI always appears on top
    gl.Disable(gl.DEPTH_TEST)
    gl.Disable(gl.LIGHTING)
    gl.Disable(gl.TEXTURE_2D)
    gl.Disable(gl.CULL_FACE)
    
	// For alpha
	gl.Enable(gl.BLEND)
    gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

    // Switch to 2D orthographic projection
    gl.MatrixMode(gl.PROJECTION)
    gl.LoadIdentity()
    gl.Ortho(0, float64(a.ScreenSize.Width), float64(a.ScreenSize.Height), 0, -1, 1)
    
    // Reset modelview
    gl.MatrixMode(gl.MODELVIEW)
    gl.LoadIdentity()
    
	a.ui.Draw()
    
    // Re-enable depth test for next frame's 3D rendering
    gl.Enable(gl.DEPTH_TEST)
}