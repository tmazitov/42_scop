package rende

import (
	"github.com/tmazitov/42_scop/internal/geom"
	"github.com/go-gl/gl/v2.1/gl"
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

func (o *Object) Draw(screenSize ScreenSize) {
	gl.BindVertexArray(o.VAO(screenSize))
	
	// Apply material before drawing
	if materials := o.materials; len(materials) != 0 {
		for _, material := range materials {
			material.Apply()
		}
	} else {
		// Default material if none specified
		gl.Color3f(1.0, 1.0, 1.0)
	}
	
	gl.DrawElements(gl.TRIANGLES, int32(o.IndicesCount()), gl.UNSIGNED_INT, nil)
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

func (o *Object) Materials() []*Material {
	return o.materials
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