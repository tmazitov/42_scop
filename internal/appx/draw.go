package appx

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/tmazitov/42_scop/internal/rende"
)

func (a *App) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    
	projection := rende.MakeProjection(a.screenSize, 1093.55, a.Camera().Zoom)

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

    lightPos := []float32{10.0, 10.0, 10.0, 1.0}
    gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPos[0])

	for _, obj := range a.Objects() {
		gl.PushMatrix()
		model := mgl32.Ident4()
		gl.MultMatrixf(&model[0])

		obj.Draw(a.screenSize)

		gl.PopMatrix()
	}

}

// func logModelBounds(vertices []*geom.Vertex) {
//     if len(vertices) == 0 {
//         return
//     }

//     minX, maxX := vertices[0].Pos.X, vertices[0].Pos.X
//     minY, maxY := vertices[0].Pos.Y, vertices[0].Pos.Y
//     minZ, maxZ := vertices[0].Pos.Z, vertices[0].Pos.Z

//     for _, v := range vertices {
//         if v.Pos.X < minX { minX = v.Pos.X }
//         if v.Pos.X > maxX { maxX = v.Pos.X }
//         if v.Pos.Y < minY { minY = v.Pos.Y }
//         if v.Pos.Y > maxY { maxY = v.Pos.Y }
//         if v.Pos.Z < minZ { minZ = v.Pos.Z }
//         if v.Pos.Z > maxZ { maxZ = v.Pos.Z }
//     }

//     log.Printf("Model bounds:")
//     log.Printf("  X: %.2f to %.2f (size: %.2f)", minX, maxX, maxX-minX)
//     log.Printf("  Y: %.2f to %.2f (size: %.2f)", minY, maxY, maxY-minY)
//     log.Printf("  Z: %.2f to %.2f (size: %.2f)", minZ, maxZ, maxZ-minZ)
    
//     centerX := (minX + maxX) / 2
//     centerY := (minY + maxY) / 2
//     centerZ := (minZ + maxZ) / 2
    
//     maxSize := maxX - minX
//     if maxY-minY > maxSize { maxSize = maxY - minY }
//     if maxZ-minZ > maxSize { maxSize = maxZ - minZ }
    
//     log.Printf("  Center: %.2f, %.2f, %.2f", centerX, centerY, centerZ)
//     log.Printf("  Max dimension: %.2f", maxSize)
// }

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
    gl.Ortho(0, float64(a.screenSize.Width), float64(a.screenSize.Height), 0, -1, 1)
    
    // Reset modelview
    gl.MatrixMode(gl.MODELVIEW)
    gl.LoadIdentity()
    
	a.ui.Draw()
    
    // Re-enable depth test for next frame's 3D rendering
    gl.Enable(gl.DEPTH_TEST)
}