package rende 

import (
	"github.com/tmazitov/42_scop/internal/clr"
	"github.com/go-gl/gl/v2.1/gl"
)

type Material struct {
	name				string
	shininess			float32
	density				float32
	dissolve			float32
	ambientColor		*clr.Color
	diffuseColor		*clr.Color
	specularColor 		*clr.Color
	illuminationModel	int
	textureId			uint32
    startIndex          int
    count               int
}

func NewMaterial() *Material {
	return &Material{
		name: "",
		dissolve: 1.0,
		density: 0.0,
		shininess: 0.0,
		ambientColor: nil,
		diffuseColor: nil,
		specularColor: nil,
		illuminationModel: 0,
		textureId: 0,
        startIndex: 0,
        count: 0,
	}
}

func (m *Material) Name() string {
    return m.name
}

func (m *Material) StartRange(startIndex int) {
    m.startIndex = startIndex
} 

func (m *Material) IncreaseRange(value int) {
    m.count += value
}

func (m *Material) Range() (int, int) {
    return m.startIndex, m.count
}

func (m *Material) Apply() {

    if m.textureId != 0 {
        gl.Enable(gl.TEXTURE_2D)
        gl.BindTexture(gl.TEXTURE_2D, m.textureId)
        
        // Texture combines with material color
        gl.TexEnvi(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE)
    } else {
        gl.Disable(gl.TEXTURE_2D)
    }
    
    // Ambient Color (Ka)
    if m.ambientColor != nil {
        ambient := m.ambientColor.Vector()
        gl.Materialfv(gl.FRONT_AND_BACK, gl.AMBIENT, &ambient[0])
    }

    // Diffuse Color (Kd) — alpha comes from dissolve, not from the color itself
    if m.diffuseColor != nil {
        diffuse := m.diffuseColor.Vector()
        diffuse[3] = m.dissolve
        gl.Materialfv(gl.FRONT_AND_BACK, gl.DIFFUSE, &diffuse[0])
        gl.Color4fv(&diffuse[0])
    }
    
    // Specular Color (Ks)
    if m.specularColor != nil {
        specular := m.specularColor.Vector()
        gl.Materialfv(gl.FRONT_AND_BACK, gl.SPECULAR, &specular[0])
    }
    
    // Shininess (Ns) - clamp to OpenGL range
    shininess := m.shininess
    if shininess > 128.0 {
        shininess = 128.0
    }
    gl.Materialf(gl.FRONT_AND_BACK, gl.SHININESS, shininess)
    
    // Transparency handling
    if m.dissolve < 1.0 {
        gl.Enable(gl.BLEND)
        gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
    } else {
        gl.Disable(gl.BLEND)
    }
    
    // Illumination model
    if m.illuminationModel == 0 {
        gl.Disable(gl.LIGHTING)
    } else {
        gl.Enable(gl.LIGHTING)
        gl.Enable(gl.LIGHT0)
    }
}

func (m *Material) SetName(name string) {
	m.name = name
}

func (m *Material) SetShininess(shininess float32) {
	m.shininess = shininess
}

func (m *Material) SetAmbientColor(color *clr.Color) {
	m.ambientColor = color
}

func (m *Material) SetDiffuseColor(color *clr.Color) {
	m.diffuseColor = color
}

func (m *Material) SetSpecularColor(color *clr.Color) {
	m.specularColor = color
}

func (m *Material) SetDissolve(dissolve float32) {
	m.dissolve = dissolve
}

func (m *Material) SetDensity(density float32) {
	m.density = density
}

func (m *Material) SetIlluminationModel(model int) {
	m.illuminationModel = model
}

func (m *Material) SetTextureId(textureId uint32) {
	m.textureId = textureId
}

func (m *Material) Cleanup() {
    if m.textureId != 0 {
        gl.DeleteTextures(1, &m.textureId)
        m.textureId = 0
    }
}