package rende

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"fmt"
)
func MakeVao(screen ScreenSize, points []*Point) uint32 {

	var convertedPoints = make([]float32, 0, len(points) * 3)
	for _, point := range points {
		convertedValue := point.Convert(screen)
		for _, value := range convertedValue {
			convertedPoints = append(convertedPoints, value)
		}
		fmt.Println(point.ToString(), convertedValue)
	}
	
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(convertedPoints), gl.Ptr(convertedPoints), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}