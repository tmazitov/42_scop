package camera

func (c *Camera) ZoomHandler(yoffset float64) {

    // Adjust FOV for zoom effect
    c.Zoom -= float32(yoffset) * 2.0
    
    // Clamp FOV
    if c.Zoom < 10.0 {
        c.Zoom = 10.0
    }
    if c.Zoom > 120.0 {
        c.Zoom = 120.0
    }
}
