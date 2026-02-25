package camera

import (
	"github.com/go-gl/mathgl/mgl32"
)

func (c *Camera) SetMovementSpeed(movementSpeed float32) *Camera {
	c.MovementSpeed = movementSpeed
	return c	
}

func (c *Camera) SetMouseSensitivity(sensitivity float32) *Camera {
	c.MouseSensitivity = sensitivity
	return c
}

func (c *Camera) SetVisionAngle(verticalDegree, horizontalDegree float32) *Camera {
	c.Pitch = verticalDegree
	c.Yaw = horizontalDegree

	c.updateCameraVectors()

	return c
}

func (c *Camera) SetPosition(position mgl32.Vec3) *Camera {
	c.Position = position

	return c
}

func (c *Camera) SerZoom(zoom float32) *Camera {
	
	c.Zoom = zoom

	c.updateCameraVectors()

	return c
}