package rende 

type Material struct {
	name		string
	shininess	float32
}

func NewMaterial() *Material {
	return &Material{
		name: "",
		shininess: 0.0, 
	}
}

func (m *Material) SetName(name string) {
	m.name = name
}

func (m *Material) SetShininess(shininess float32) {
	m.shininess = shininess
}