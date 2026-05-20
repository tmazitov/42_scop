package camera

import (
	"math"
	"github.com/tmazitov/42_scop/internal/geom")

func (c *Camera) updateCameraVectors() {
	// Calculate the new Front vector
	front := geom.Vec3{
		float32(math.Cos(float64(geom.DegToRad(c.Yaw))) * math.Cos(float64(geom.DegToRad(c.Pitch)))),
		float32(math.Sin(float64(geom.DegToRad(c.Pitch)))),
		float32(math.Sin(float64(geom.DegToRad(c.Yaw))) * math.Cos(float64(geom.DegToRad(c.Pitch)))),
	}
	c.Front = front.Normalize()
	c.Right = c.Front.Cross(c.WorldUp).Normalize()
	c.Up = c.Right.Cross(c.Front).Normalize()
}