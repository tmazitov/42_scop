package appx

// In your main/init function:
// window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled) // Capture mouse

import (
    "github.com/go-gl/glfw/v3.2/glfw"
)

func (c *controller) BindMouseControl() {
	c.app.Window().Core().SetCursorPosCallback(c.mouseCallback)
}

// Mouse callback (set this up in your initialization)
func (c *controller) mouseCallback(w *glfw.Window, xpos float64, ypos float64) {
    if c.firstMouse {
        c.lastX = xpos
        c.lastY = ypos
        c.firstMouse = false
    }
    
    xoffset := float32(xpos - c.lastX)
    yoffset := float32(c.lastY - ypos) // reversed: y ranges bottom to top
    
    c.lastX = xpos
    c.lastY = ypos
    
    c.app.Camera().ProcessMouseMovement(xoffset, yoffset)
}
