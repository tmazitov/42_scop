package rende

import (
	"fmt"
)

type Point struct {
	X float32
	Y float32
	Z float32
}

func NewPoint(x, y, z float32) *Point {
	return &Point{
		X: x,
		Y: y,
		Z: z,
	}
}

func convertY(value, center float32) float32{

	if value == center {
		return 0
	}

	if value < center {
		return value / center
	}
	
	return -1 * ((value / center) - 1)
}

func convertX(value, center float32) float32{

	if value == center {
		return 0
	}

	if value < center {
		return -1 * value / center
	}
	
	return ((value / center) - 1)
}


func (p *Point) ToString() string {
	return fmt.Sprintf("Point [%f, %f, %f]", p.X, p.Y, p.Z)
}

func (p *Point) Convert(screen ScreenSize) [3]float32 {
	relativeX := convertX(p.X, screen.Width / 2)
	relativeY := convertY(p.Y, screen.Height / 2)
	return [3]float32{
		relativeX,
		relativeY,
		0,
	}	
}