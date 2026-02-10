package appx

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type controller struct {
	app         *App
	lastX       float64
	lastY       float64
	firstMouse  bool
	deltaTime   float32
	lastFrame   float32
}

func newController(app *App) *controller {
	return &controller{
		app:        app,
		lastX:      float64(app.ScreenSize.Width) / 2,
		lastY:      float64(app.ScreenSize.Height) / 2,
		firstMouse: true,
		deltaTime:  0.0,
		lastFrame:  0.0,
	}
}

// Process keyboard input for camera movement
func (c *controller) processInput(window *glfw.Window, camera *Camera) {
	// Update delta time
	currentFrame := float32(glfw.GetTime())
	c.deltaTime = currentFrame - c.lastFrame
	c.lastFrame = currentFrame

	// Movement speed
	if window.GetKey(glfw.KeyW) == glfw.Press {
		camera.ProcessKeyboard("FORWARD", c.deltaTime)
	}
	if window.GetKey(glfw.KeyS) == glfw.Press {
		camera.ProcessKeyboard("BACKWARD", c.deltaTime)
	}
	if window.GetKey(glfw.KeyA) == glfw.Press {
		camera.ProcessKeyboard("LEFT", c.deltaTime)
	}
	if window.GetKey(glfw.KeyD) == glfw.Press {
		camera.ProcessKeyboard("RIGHT", c.deltaTime)
	}
	if window.GetKey(glfw.KeySpace) == glfw.Press {
		camera.ProcessKeyboard("UP", c.deltaTime)
	}
	if window.GetKey(glfw.KeyLeftShift) == glfw.Press {
		camera.ProcessKeyboard("DOWN", c.deltaTime)
	}

	// ESC to close
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}

// Mouse callback for camera look
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

	c.app.camera.ProcessMouseMovement(xoffset, yoffset)
}

// Bind mouse control to window
func (c *controller) BindMouseControl() {
	window := c.app.window.Core()
	
	// Capture and hide cursor
	// window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	
	// Set cursor position callback
	window.SetCursorPosCallback(c.mouseCallback)
}