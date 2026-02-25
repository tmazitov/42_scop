package rende 

import (
	"github.com/go-gl/mathgl/mgl32"
)

func MakeProjection(screenSize ScreenSize, maxDimension float32) mgl32.Mat4 {
	return mgl32.Perspective(
		mgl32.DegToRad(45.0),
		screenSize.Width/screenSize.Height,
		0.1,
		maxDimension * 10,
	)
}