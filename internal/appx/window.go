package appx

import (
	// "github.com/go-gl/gl/v3.3-core/gl"
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

func initGlfw(opts *WindowOptions) (*glfw.Window, error) {
	if err := glfw.Init(); err != nil {
		return nil, err
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(opts.Width, opts.Height, opts.Title, nil, nil)
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