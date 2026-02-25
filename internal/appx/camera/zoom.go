package camera

func (c *Camera) ZoomIn() {

	if c.Zoom == 0.1 {
		return
	}

	c.Zoom -= 1.0 
}

func (c *Camera) ZoomOut() {
	if c.Zoom == 45.0 {
		return
	}

	c.Zoom += 1.0
}