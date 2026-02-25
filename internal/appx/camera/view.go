package camera 

import (
    "github.com/go-gl/mathgl/mgl32"
)

func (c *Camera) GetViewMatrix() mgl32.Mat4 {
	return mgl32.LookAtV(c.Position, c.Position.Add(c.Front), c.Up)
}