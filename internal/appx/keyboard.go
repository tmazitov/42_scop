package appx

import (
    "github.com/go-gl/glfw/v3.2/glfw"
	// "fmt"
)


func (c *controller) processInput(window *glfw.Window, camera *Camera) {
    
	currentFrame := float32(glfw.GetTime())
	deltaTime := currentFrame - c.lastFrame
	c.lastFrame = currentFrame

    if window.GetKey(glfw.KeyW) == glfw.Press {
        camera.ProcessKeyboard("FORWARD", deltaTime)
    }
    if window.GetKey(glfw.KeyS) == glfw.Press {
        camera.ProcessKeyboard("BACKWARD", deltaTime)
    }
    if window.GetKey(glfw.KeyA) == glfw.Press {
        camera.ProcessKeyboard("LEFT", deltaTime)
    }
    if window.GetKey(glfw.KeyD) == glfw.Press {
        camera.ProcessKeyboard("RIGHT", deltaTime)
    }
    if window.GetKey(glfw.KeyEscape) == glfw.Press {
        window.SetShouldClose(true)
    }
}

