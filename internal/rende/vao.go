package rende

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/tmazitov/42_scop/internal/geom"
	// "fmt"
)
func MakeVao(screen ScreenSize, indices []uint32, points []*geom.Vertex) uint32 {
    // Interleave position and normal data
    var vertexData []float32
    for _, point := range points {
        vertexData = append(vertexData, point.Pos.X, point.Pos.Y, point.Pos.Z)     // Position
        vertexData = append(vertexData, point.Norm.X, point.Norm.Y, point.Norm.Z)  // Normal
    }

    var vao uint32
    gl.GenVertexArrays(1, &vao)
    gl.BindVertexArray(vao)

    var vbo uint32
    gl.GenBuffers(1, &vbo)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
    gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertexData), gl.Ptr(vertexData), gl.STATIC_DRAW)

    var ebo uint32
    gl.GenBuffers(1, &ebo)
    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
    gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

    stride := int32(6 * 4) // 6 floats per vertex (3 pos + 3 normal)

    // Enable and set vertex position
    gl.EnableClientState(gl.VERTEX_ARRAY)
    gl.VertexPointer(3, gl.FLOAT, stride, gl.PtrOffset(0))

    // Enable and set vertex normal
    gl.EnableClientState(gl.NORMAL_ARRAY)
    gl.NormalPointer(gl.FLOAT, stride, gl.PtrOffset(3*4))

    gl.BindVertexArray(0)
    gl.DisableClientState(gl.VERTEX_ARRAY)
    gl.DisableClientState(gl.NORMAL_ARRAY)

    return vao
}