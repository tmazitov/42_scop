package controller

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	// "fmt"
	"github.com/tmazitov/42_scop/internal/appx/camera"
)

type Controller struct {
	app			App		
	lastX       float64
	lastY       float64
	firstMouse  bool
	deltaTime   float32
	lastFrame   float32
}

func NewController(app App) *Controller {
	return &Controller{
		app:     	app,
		lastX:      float64(app.ScreenSize().Width) / 2,
		lastY:      float64(app.ScreenSize().Height) / 2,
		firstMouse: true,
		deltaTime:  0.0,
		lastFrame:  0.0,
	}
}

// Process keyboard input for camera movement
func (c *Controller) ProcessInput(window *glfw.Window, camera *camera.Camera) {
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




// Bind mouse control to window
func (c *Controller) BindMouseControl() {
	window := c.app.Window().Core()
	
	// Capture and hide cursor
	// window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)

	// Set cursor position callback
	window.SetCursorPosCallback(c.mouseMoveCallback)
	window.SetMouseButtonCallback(c.mouseClickCallback)
}