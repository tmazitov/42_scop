package rende

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	// "fmt"
)
func MakeVao(screen ScreenSize, indices []uint32, points []*Vertex) uint32 {
    var convertedVertex = make([]float32, 0, len(points) * 3)
    for _, point := range points {
        convertedVertex = append(convertedVertex, point.Vector()...)
    }
    
    // Create and bind VAO FIRST
    var vao uint32
    gl.GenVertexArrays(1, &vao)
    gl.BindVertexArray(vao)  // Bind VAO before setting up buffers
    
    // Create and setup VBO
    var vbo uint32
    gl.GenBuffers(1, &vbo)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
    gl.BufferData(gl.ARRAY_BUFFER, 4*len(convertedVertex), gl.Ptr(convertedVertex), gl.STATIC_DRAW)
    
    // Create and setup EBO (AFTER VAO is bound)
    var ebo uint32
    gl.GenBuffers(1, &ebo)
    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
    gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)
    
    // Setup vertex attributes
    gl.EnableVertexAttribArray(0)
    gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
    
    // Unbind (optional, but good practice)
    gl.BindVertexArray(0)
    
    return vao
}