package rende 

import (
	"github.com/tmazitov/42_scop/internal/clr"
)

type Material struct {
	name		string
	shininess	float32
	ambientColor *clr.Color
	diffuseColor *clr.Color
}

func NewMaterial() *Material {
	return &Material{
		name: "",
		shininess: 0.0,
		ambientColor: nil,
		diffuseColor: nil, 
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