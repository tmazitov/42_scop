package geom

import "math"

func DegToRad(deg float32) float32 {
	return deg * (math.Pi / 180)
}