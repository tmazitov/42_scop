package rende

import (
	"github.com/go-gl/gl/v2.1/gl"
	// "fmt"
)
func MakeVao(screen ScreenSize, indices []uint32, points []*Vertex) uint32 {
	var convertedVertex = make([]float32, 0, len(points)*3)
	for _, point := range points {
		convertedVertex = append(convertedVertex, point.Vector()...)
	}

	// Create and bind VAO
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// Create and setup VBO
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(convertedVertex), gl.Ptr(convertedVertex), gl.STATIC_DRAW)

	// Create and setup EBO
	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	// Setup vertex attributes (no shader, so use legacy client state)
	gl.EnableClientState(gl.VERTEX_ARRAY)
	gl.VertexPointer(3, gl.FLOAT, 0, nil)

	gl.BindVertexArray(0)
	gl.DisableClientState(gl.VERTEX_ARRAY)

	return vao
}