package app

import "github.com/go-gl/glfw/v3.2/glfw"

type Window struct {
	height int
	width  int
	core   *glfw.Window
}

func NewWindow(width, height int) (*Window, error) {
	core, err := glfw.CreateWindow(width, height, "Conway's Game of Life", nil, nil)
	if err != nil {
		return nil, err
	}
	core.MakeContextCurrent()
	return &Window{
		height: height,
		width:  width,
		core:   core,
	}, nil
}

func (w *Window) Core() *glfw.Window {
	return w.core
}

func (w *Window) Run() {

}
