package controller

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
)

func (c *Controller) mouseClickCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {

	if button == glfw.MouseButtonLeft && action == glfw.Release {
		xpos, ypos := window.GetCursorPos()
	
		pressedButton := c.app.UI().IsPressed(float32(xpos), float32(ypos))
		if pressedButton == nil {
			return ;
		}
		err := pressedButton.HandleClick(float32(xpos), float32(ypos))
		if err != nil {
			log.Println("ui pressed error : ", err)
		}
	}
}

func (c *Controller) mouseMoveCallback(w *glfw.Window, xpos float64, ypos float64) {
	if c.firstMouse {
		c.lastX = xpos
		c.lastY = ypos
		c.firstMouse = false
	}

	xoffset := float32(xpos - c.lastX)
	yoffset := float32(c.lastY - ypos)

	c.lastX = xpos
	c.lastY = ypos

	if w.GetMouseButton(glfw.MouseButtonLeft) == glfw.Press {
		c.app.Camera().ProcessMouseMovement(xoffset, yoffset)
	}

	if w.GetMouseButton(glfw.MouseButtonRight) == glfw.Press {
		scale := c.app.Camera().MovementSpeed * 0.005
		right := c.app.Camera().Right
		up := c.app.Camera().Up
		dx := (right.X()*xoffset + up.X()*yoffset) * scale
		dy := (right.Y()*xoffset + up.Y()*yoffset) * scale
		dz := (right.Z()*xoffset + up.Z()*yoffset) * scale
		c.app.TranslateSelectedObject(dx, dy, dz)
	}
}

func (c *Controller) scrollCallback(w *glfw.Window, xoffset float64, yoffset float64) {
	c.app.Camera().ZoomHandler(yoffset)
}