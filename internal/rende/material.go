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
	emissiveColor		*clr.Color
	transmissionFilter	*clr.Color
	illuminationModel	    int
	textureId			    uint32
	specularTextureId	    uint32
	shininessTextureId	    uint32
	dissolveTextureId	    uint32
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

// Apply sets up material properties. blend controls texture visibility:
// 0 = colors only, 1 = full texture, between = texture fading in at that alpha.
func (m *Material) Apply(blend float32) {

    // Ambient Color (Ka)
    if m.ambientColor != nil {
        ambient := m.ambientColor.Vector()
        gl.Materialfv(gl.FRONT_AND_BACK, gl.AMBIENT, &ambient[0])
    }

    // Diffuse Color (Kd)
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

    // Emissive Color (Ke)
    if m.emissiveColor != nil {
        emissive := m.emissiveColor.Vector()
        emissive[3] = 1.0
        gl.Materialfv(gl.FRONT_AND_BACK, gl.EMISSION, &emissive[0])
    } else {
        black := []float32{0, 0, 0, 1}
        gl.Materialfv(gl.FRONT_AND_BACK, gl.EMISSION, &black[0])
    }

    // Shininess (Ns)
    shininess := m.shininess
    if shininess > 128.0 {
        shininess = 128.0
    }
    gl.Materialf(gl.FRONT_AND_BACK, gl.SHININESS, shininess)

    // Illumination model
    if m.illuminationModel == 0 {
        gl.Disable(gl.LIGHTING)
    } else {
        gl.Enable(gl.LIGHTING)
        gl.Enable(gl.LIGHT0)
    }

    // Texture
    if m.textureId != 0 && blend > 0 {
        gl.Enable(gl.TEXTURE_2D)
        gl.BindTexture(gl.TEXTURE_2D, m.textureId)
        gl.TexEnvi(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE)
        gl.Color4f(1, 1, 1, blend*m.dissolve)
        gl.Enable(gl.BLEND)
        gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
    } else {
        gl.Disable(gl.TEXTURE_2D)
        if m.dissolve < 1.0 {
            gl.Enable(gl.BLEND)
            gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
        } else {
            gl.Disable(gl.BLEND)
        }
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

func (m *Material) SetEmissiveColor(color *clr.Color) {
	m.emissiveColor = color
}

func (m *Material) SetTransmissionFilter(color *clr.Color) {
	m.transmissionFilter = color
}

func (m *Material) TextureId() uint32 {
	return m.textureId
}

func (m *Material) SetTextureId(textureId uint32) {
	m.textureId = textureId
}

func (m *Material) SpecularTextureId() uint32 {
	return m.specularTextureId
}

func (m *Material) SetSpecularTextureId(id uint32) {
	m.specularTextureId = id
}

func (m *Material) ShininessTextureId() uint32 {
	return m.shininessTextureId
}

func (m *Material) SetShininessTextureId(id uint32) {
	m.shininessTextureId = id
}

func (m *Material) DissolveTextureId() uint32 {
	return m.dissolveTextureId
}

func (m *Material) SetDissolveTextureId(id uint32) {
	m.dissolveTextureId = id
}

func (m *Material) Cleanup() {
    if m.textureId != 0 {
        gl.DeleteTextures(1, &m.textureId)
        m.textureId = 0
    }
    if m.specularTextureId != 0 {
        gl.DeleteTextures(1, &m.specularTextureId)
        m.specularTextureId = 0
    }
    if m.shininessTextureId != 0 {
        gl.DeleteTextures(1, &m.shininessTextureId)
        m.shininessTextureId = 0
    }
    if m.dissolveTextureId != 0 {
        gl.DeleteTextures(1, &m.dissolveTextureId)
        m.dissolveTextureId = 0
    }
}