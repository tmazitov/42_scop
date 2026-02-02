package rende

import (
	"fmt"
)

type Point struct {
	x float32
	y float32
	z float32
}

func NewPoint(x, y, z float32) *Point {
	return &Point{
		x: x,
		y: y,
		z: z,
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
	return fmt.Sprintf("Point [%f, %f, %f]", p.x, p.y, p.z)
}

func (p *Point) Convert(screen ScreenSize) [3]float32 {
	relativeX := convertX(p.x, screen.Width / 2)
	relativeY := convertY(p.y, screen.Height / 2)
	return [3]float32{
		relativeX,
		relativeY,
		0,
	}
}