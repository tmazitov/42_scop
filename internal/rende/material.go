package rende 

import (
	"github.com/tmazitov/42_scop/internal/clr"
)

type Material struct {
	name		string
	shininess	float32
	density		float32
	dissolve	float32
	ambientColor *clr.Color
	diffuseColor *clr.Color
	specularColor *clr.Color
	illuminationModel int
}

func NewMaterial() *Material {
	return &Material{
		name: "",
		dissolve: 0.0,
		density: 0.0,
		shininess: 0.0,
		ambientColor: nil,
		diffuseColor: nil,
		specularColor: nil,
		illuminationModel: 0,
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