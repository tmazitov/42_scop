package rende 

import (
	"github.com/go-gl/mathgl/mgl32"
)

func MakeProjection(screenSize ScreenSize, dimension, angle float32) mgl32.Mat4 {
	return mgl32.Perspective(
		mgl32.DegToRad(angle),
		screenSize.Width/screenSize.Height,
		0.1,
		dimension * 10,
	)
}