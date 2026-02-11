package rende

import (
	"github.com/tmazitov/42_scop/internal/geom"
)

type Object struct {
	name 	string
	shape 	[]*geom.Vertex
	indices []uint32
	vao 	uint32
}

func NewObject(name string, shape []*geom.Vertex, indices []uint32) *Object {
	return &Object{
		name: name,
		shape: shape,
		indices: indices,
		vao: 0,
	}
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