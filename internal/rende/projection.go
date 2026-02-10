package rende 

import (
	"github.com/go-gl/mathgl/mgl32"
)

func MakeProjection(height, width float32) mgl32.Mat4 {
	return mgl32.Perspective(
		mgl32.DegToRad(45.0),
		width/height,
		0.1,
		100.0,
	)
}