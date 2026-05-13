package appx

import (
	"github.com/go-gl/gl/v2.1/gl"
	// "github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/tmazitov/42_scop/internal/rende"
	"github.com/tmazitov/42_scop/internal/ui"
	"github.com/tmazitov/42_scop/internal/appx/camera"
	"github.com/tmazitov/42_scop/internal/appx/window"
	"github.com/tmazitov/42_scop/internal/appx/controller"
	"log"
)

type App struct {
	config     *Config
	controller *controller.Controller
	window     *window.Window
	camera     *camera.Camera
	state	   *State
	ui		   *ui.UI
	objects    []*rende.Object
	screenSize rende.ScreenSize
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
	initLight()

	// Set clear color (background)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	return nil
}

func NewApp(config *Config) (*App, error) {
	win, err := window.NewWindow(config.Window)
	if err != nil {
		return nil, err
	}

	err = initOpenGL()
	if err != nil {
		return nil, err
	}
	app := &App{
		config:  config,
		window:  win,
		state:	 NewState(),
		camera:  camera.NewCamera(
			mgl32.Vec3{
				0, 0, 0,
			}).
			SetMouseSensitivity(0.1).
			SetVisionAngle(0, -90),
		objects: nil,
		controller: nil,
		ui:			ui.NewUI(),
		screenSize: rende.ScreenSize{
			Height: float32(config.Window.Height),
			Width:  float32(config.Window.Width),
		},
	}

	app.controller = controller.NewController(app)
	app.controller.BindMouseControl()
	app.SetupButtons()

	return app, nil
}

// Rest of your methods remain the same...
func (a *App) Process() {
	a.controller.ProcessInput(a.window.Core(), a.camera)

	if a.state.IsRotationEnabled {
		for _, obj := range a.objects {
			obj.Rotate(0, 0.001, 0)
		}
	}
}

func (a *App) Camera() *camera.Camera {
	return a.camera
}

func (a *App) UI() *ui.UI {
	return a.ui
}

func (a *App) ScreenSize() rende.ScreenSize {
	return a.screenSize
}

func (a *App) AddObjects(objs ...*rende.Object) {
	a.objects = append(a.objects, objs...)
	
	var y float32 = 32
	var x float32 = a.screenSize.Width - float32(200)
	for _, objectInfoElem := range a.objects[0].Info() {
		text := ui.NewText(objectInfoElem, x, y)
		a.ui.AddStaticText(text)
		y += 28
	}

	pos := calculateCameraPosition(a.objects[0].Shape())

	speed := calculateCameraSpeed(a.objects[0].Shape())

	a.camera.
	SetMovementSpeed(speed).
	SetPosition(mgl32.Vec3{
		pos.X,
		pos.Y,
		pos.Z,
	})
}

func (a *App) Objects() []*rende.Object {
	return a.objects
}

func (a *App) Window() *window.Window {
	return a.window
}

func (a *App) Close() {
	a.window.Close()

	for _, object := range a.objects {
		object.Cleanup()
	}
}