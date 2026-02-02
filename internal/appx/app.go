package appx

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	// "github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"github.com/tmazitov/42_scop/internal/rende"
)

type App struct {
	config	*Config
	window	*Window
	core	uint32
	objects []*rende.Object
	ScreenSize rende.ScreenSize
}


// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() (uint32, error) {

	if err := gl.Init(); err != nil {
		return 0, err
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog, nil
}

func NewApp(config *Config) (*App, error) {

	window, err := NewWindow(config.Window())
	if err != nil {
		return nil, err
	}

	core, err := initOpenGL() 
	if err != nil{
		return nil, err
	}

	return &App{
		config: config,
		window: window,
		core: core,
		objects: nil,
		ScreenSize: rende.ScreenSize{
			Height: float32(config.Window().Height),
			Width: float32(config.Window().Width),
		},
	}, nil
}

func (a *App) AddObject(obj *rende.Object) {
	a.objects = append(a.objects, obj)
}

func (a *App) Objects() []*rende.Object{
	return a.objects
}

func (a *App) Core() uint32 {
	return a.core
}

func (a *App) Window() *Window {
	return a.window
}

func (a *App) Close() {
	a.window.Close()
}
