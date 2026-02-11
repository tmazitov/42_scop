package rende

import (
	"github.com/tmazitov/42_scop/internal/geom"
)

type Object struct {
	name 	string
	shape 	[]*geom.Vertex
	indices []uint32
	vao 	uint32
	materials []*Material
}

func NewObject(name string) *Object {
	return &Object{
		name: name,
		shape: nil,
		indices: nil,
		materials: nil,
		vao: 0,
	}
}

func (o *Object) SetShape(shape []*geom.Vertex) *Object {
	o.shape = shape
	return o
}
func (o *Object) SetIndices(indices []uint32) *Object {
	o.indices = indices
	return o
}
func (o *Object) SetMaterials(materials []*Material) *Object {
	o.materials = materials
	return o
}

func (o *Object) Name() string {
	return o.name
}

func (o *Object) Shape() []*geom.Vertex{
	return o.shape
}

func (o *Object) VAO(screen ScreenSize) uint32 {

	if o.vao != 0 {
		return o.vao
	}

	o.vao = MakeVao(screen, o.indices, o.shape)

	return o.vao 
}

func (o *Object) NodeCount() int32 {
	return int32(len(o.shape))
}

func (o *Object) IndicesCount() int32 {
	return int32(len(o.indices))
}