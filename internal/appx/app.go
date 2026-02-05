package appx

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	// "github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/tmazitov/42_scop/internal/rende"
)

type App struct {
	controller *controller
	config	*Config
	window	*Window
	camera	*Camera
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

	app := &App{
		config: config,
		window: window,
		core: core,
		camera: NewCamera(mgl32.Vec3{0, 0, 3}, mgl32.Vec3{0, 1, 0}, -90, 0),
		objects: nil,
		controller: nil, 
		ScreenSize: rende.ScreenSize{
			Height: float32(config.Window().Height),
			Width: float32(config.Window().Width),
		},
	}

	app.controller = newController(app)

	app.controller.BindMouseControl()

	return app, nil
}

func (a *App) Process() {
	a.controller.processInput(a.window.Core(), a.camera)
}

func (a *App) Camera() *Camera {
	return a.camera
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
