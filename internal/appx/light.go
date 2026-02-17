package appx

import (
	"github.com/tmazitov/42_scop/internal/geom"
	"github.com/go-gl/gl/v2.1/gl"
)

type Lighter struct {
	pos *geom.Pos
}

func initLight() {
    gl.Enable(gl.LIGHTING)
    gl.Enable(gl.LIGHT0)
    
    gl.Enable(gl.COLOR_MATERIAL)
    gl.ColorMaterial(gl.FRONT_AND_BACK, gl.AMBIENT_AND_DIFFUSE)
    
    // REDUCE ambient light (was probably 0.3, make it darker)
    ambient := []float32{0.3, 0.3, 0.3, 1.0}  // Lower = darker shadows
    gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])
    
    // Adjust diffuse light
    diffuse := []float32{0.7, 0.7, 0.7, 1.0}  // Main light intensity
    gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[0])
    
    // Specular for highlights
    specular := []float32{0.3, 0.3, 0.3, 1.0}  // Subtle highlights
    gl.Lightfv(gl.LIGHT0, gl.SPECULAR, &specular[0])
    
    gl.ShadeModel(gl.SMOOTH)
}
