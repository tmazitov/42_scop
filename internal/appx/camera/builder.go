package camera


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

func (c *Camera) SetZoom(zoom float32) *Camera {
	
	c.Zoom = zoom

	c.updateCameraVectors()

	return c
}