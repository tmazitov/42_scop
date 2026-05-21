package rende

import (
	"github.com/tmazitov/42_scop/internal/geom"
	"github.com/go-gl/gl/v2.1/gl"
	"fmt"
	"unsafe"
)

type Object struct {
	name 	  string
	shape 	  []*geom.Vertex
	indices   []uint32
	vao 	  uint32
	materials []*Material
	pivot     geom.Pos
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

func bottomCenter(shape []*geom.Vertex) geom.Pos {
	minX, maxX := shape[0].Pos.X, shape[0].Pos.X
	minY := shape[0].Pos.Y
	minZ, maxZ := shape[0].Pos.Z, shape[0].Pos.Z
	for _, v := range shape[1:] {
		if v.Pos.X < minX { minX = v.Pos.X }
		if v.Pos.X > maxX { maxX = v.Pos.X }
		if v.Pos.Y < minY { minY = v.Pos.Y }
		if v.Pos.Z < minZ { minZ = v.Pos.Z }
		if v.Pos.Z > maxZ { maxZ = v.Pos.Z }
	}
	return geom.Pos{X: (minX + maxX) / 2, Y: minY, Z: (minZ + maxZ) / 2}
}

func (o *Object) Info() []string{
	return []string{
		fmt.Sprintf("Name: %s", o.name),
		fmt.Sprintf("Vertices: %d", len(o.shape)),
		fmt.Sprintf("Faces: %d", len(o.indices) / 3),
	}
}

func (o *Object) Rotate(angleX, angleY, angleZ float32) {
	for _, vertex := range o.shape {
		vertex.Pos.X -= o.pivot.X
		vertex.Pos.Y -= o.pivot.Y
		vertex.Pos.Z -= o.pivot.Z
		vertex.Rotate(angleX, angleY, angleZ)
		vertex.Pos.X += o.pivot.X
		vertex.Pos.Y += o.pivot.Y
		vertex.Pos.Z += o.pivot.Z
	}
	o.vao = 0
}
  
func (o *Object) Draw(screenSize ScreenSize, textureBlend float32) {

	gl.BindVertexArray(o.VAO(screenSize))

	if len(o.materials) == 0 {
		gl.Color3f(1.0, 1.0, 1.0)
		gl.DrawElements(gl.TRIANGLES, int32(o.IndicesCount()), gl.UNSIGNED_INT, nil)
		return
	}

	if textureBlend < 1.0 {
		// Pass 1: solid material colors
		for _, material := range o.materials {
			start, count := material.Range()
			material.Apply(0)
			gl.DrawElements(gl.TRIANGLES, int32(count), gl.UNSIGNED_INT, unsafe.Pointer(uintptr(start*4)))
		}

		// Pass 2: texture fading in — LEQUAL allows same-depth overdraw, no depth write
		if textureBlend > 0 {
			gl.DepthFunc(gl.LEQUAL)
			gl.DepthMask(false)
			for _, material := range o.materials {
				if material.TextureId() == 0 {
					continue
				}
				start, count := material.Range()
				material.Apply(textureBlend)
				gl.DrawElements(gl.TRIANGLES, int32(count), gl.UNSIGNED_INT, unsafe.Pointer(uintptr(start*4)))
			}
			gl.DepthMask(true)
			gl.DepthFunc(gl.LESS)
		}
	} else {
		// blend == 1: single pass, full texture
		for _, material := range o.materials {
			start, count := material.Range()
			material.Apply(1.0)
			gl.DrawElements(gl.TRIANGLES, int32(count), gl.UNSIGNED_INT, unsafe.Pointer(uintptr(start*4)))
		}
	}
}


func (o *Object) SetShape(shape []*geom.Vertex) *Object {
	o.shape = shape
	if len(shape) > 0 {
		o.pivot = bottomCenter(shape)
	}
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

func (o *Object) Cleanup() {
	for _, material := range o.materials {
		material.Cleanup()
	}
}