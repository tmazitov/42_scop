package appx

import (
	"github.com/go-gl/gl/v2.1/gl"
	// "github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/tmazitov/42_scop/internal/rende"
	"github.com/tmazitov/42_scop/internal/ui"
	"github.com/tmazitov/42_scop/internal/clr"
	"log"
)

type App struct {
	controller *controller
	config     *Config
	window     *Window
	camera     *Camera
	ui		   *ui.UI
	objects    []*rende.Object
	ScreenSize rende.ScreenSize
}

// initOpenGL initializes OpenGL (no shaders needed)
func initOpenGL() error {
	if err := gl.Init(); err != nil {
		return err
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	// Enable depth testing
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	// Set clear color (background)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	return nil
}

func NewApp(config *Config) (*App, error) {
	window, err := NewWindow(config.Window)
	if err != nil {
		return nil, err
	}

	err = initOpenGL()
	if err != nil {
		return nil, err
	}

	app := &App{
		config:  config,
		window:  window,
		camera:  NewCamera(mgl32.Vec3{0, 0, 3}, mgl32.Vec3{0, 1, 0}, -90, 0),
		objects: nil,
		controller: nil,
		ui:			nil,
		ScreenSize: rende.ScreenSize{
			Height: float32(config.Window.Height),
			Width:  float32(config.Window.Width),
		},
	}

	app.controller = newController(app)
	app.controller.BindMouseControl()

	app.ui = ui.NewUI(app.ScreenSize)
	app.ui.AddButton( ui.NewButton().
		SetPos(&rende.Pos{X: 10, Y: 10, Z: 1}).
		SetSize(40, 40).
		SetColor(clr.NewColor(66, 96, 150)))

	return app, nil
}

// Rest of your methods remain the same...
func (a *App) Process() {
	a.controller.processInput(a.window.Core(), a.camera)
}

func (a *App) Camera() *Camera {
	return a.camera
}

func (a *App) AddObjects(objs ...*rende.Object) {
	a.objects = append(a.objects, objs...)
}

func (a *App) Objects() []*rende.Object {
	return a.objects
}

func (a *App) Window() *Window {
	return a.window
}

func (a *App) Close() {
	a.window.Close()
}