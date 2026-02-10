package appx

import (
	// "github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)


type Window struct {
	opts	*WindowOptions
	core	*glfw.Window
}

type WindowOptions struct {
	Height int
	Width int
	Title string
}

func initGlfw(config *WindowOptions) (*glfw.Window, error) {
	if err := glfw.Init(); err != nil {
		return nil, err
	}

	// Set OpenGL version to 2.1
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	
	// DO NOT SET PROFILE for OpenGL 2.1
	// Remove any lines like:
	// glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)
	
	glfw.WindowHint(glfw.Resizable, glfw.False)

	window, err := glfw.CreateWindow(config.Width, config.Height, config.Title, nil, nil)
	if err != nil {
		return nil, err
	}

	window.MakeContextCurrent()
	
	return window, nil
}
func NewWindow(opts *WindowOptions) (*Window, error) {
	core, err := initGlfw(opts)
	if err != nil {
		return nil, err
	}
	core.MakeContextCurrent()
	return &Window{
		opts:	opts,
		core:   core,
	}, nil
}

func (w *Window) Core() *glfw.Window {
	return w.core
}

func (w *Window) Close() {
	glfw.Terminate()
} 