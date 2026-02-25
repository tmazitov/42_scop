package controller

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tmazitov/42_scop/internal/appx/camera"
)

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