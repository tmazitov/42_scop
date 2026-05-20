package rende 

import (
	"github.com/tmazitov/42_scop/internal/geom"
)

func MakeProjection(screenSize ScreenSize, dimension, angle float32) geom.Mat4 {
	return geom.Perspective(
		angle,
		screenSize.Width/screenSize.Height,
		0.1,
		dimension * 10,
	)
}