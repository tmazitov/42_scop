package camera 

import (
	"github.com/tmazitov/42_scop/internal/geom"
)

func (c *Camera) GetViewMatrix() geom.Mat4 {
	return geom.LookAtV(c.Position, c.Position.Add(c.Front), c.Up)
}