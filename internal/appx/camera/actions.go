package camera

func (c *Camera) ProcessKeyboard(direction string, deltaTime float32) {
	velocity := c.MovementSpeed * deltaTime

	switch direction {
	case "FORWARD":
		c.Position = c.Position.Add(c.Front.Mul(velocity))
	case "BACKWARD":
		c.Position = c.Position.Sub(c.Front.Mul(velocity))
	case "LEFT":
		c.Position = c.Position.Sub(c.Right.Mul(velocity))
	case "RIGHT":
		c.Position = c.Position.Add(c.Right.Mul(velocity))
	case "UP":
		c.Position = c.Position.Add(c.WorldUp.Mul(velocity))
	case "DOWN":
		c.Position = c.Position.Sub(c.WorldUp.Mul(velocity))
	}
}

func (c *Camera) ProcessMouseMovement(xoffset, yoffset float32) {
	xoffset *= c.MouseSensitivity
	yoffset *= c.MouseSensitivity

	c.Yaw += xoffset
	c.Pitch += yoffset

	// Constrain pitch to avoid gimbal lock
	if c.Pitch > 89.0 {
		c.Pitch = 89.0
	}
	if c.Pitch < -89.0 {
		c.Pitch = -89.0
	}

	c.updateCameraVectors()
}