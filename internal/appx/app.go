package appx

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/tmazitov/42_scop/internal/appx/camera"
	"github.com/tmazitov/42_scop/internal/appx/controller"
	"github.com/tmazitov/42_scop/internal/appx/window"
	"github.com/tmazitov/42_scop/internal/geom"
	"github.com/tmazitov/42_scop/internal/rende"
	"github.com/tmazitov/42_scop/internal/ui"
)

type App struct {
	config          *Config
	controller      *controller.Controller
	window          *window.Window
	camera          *camera.Camera
	state           *State
	ui              *ui.UI
	objects         []*rende.Object
	grid            *rende.Grid
	screenSize      rende.ScreenSize
	objectInfoTexts [4]*ui.Text
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

type Config struct {
	Window *window.WindowOptions
	RotationSpeed float32
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
		config: config,
		window: win,
		state:  NewState(),
		camera: camera.NewCamera(geom.Vec3{0, 0, 0}).
			SetMouseSensitivity(0.1).
			SetVisionAngle(0, -90),
		objects:    nil,
		controller: nil,
		ui:         ui.NewUI(),
		screenSize: rende.ScreenSize{
			Height: float32(config.Window.Height),
			Width:  float32(config.Window.Width),
		},
	}

	app.controller = controller.NewController(app)
	app.controller.BindMouseControl()
	app.SetupButtons()
	app.setupObjectInfoTexts()

	return app, nil
}

func (a *App) setupObjectInfoTexts() {
	x := a.screenSize.Width - 210
	for i := range a.objectInfoTexts {
		a.objectInfoTexts[i] = ui.NewText("", x, float32(10+i*16))
	}
}

func (a *App) updateObjectInfo() {
	obj := a.SelectedObject()
	if obj == nil {
		return
	}
	t := obj.Translation()
	p := obj.Pivot()
	lines := [4]string{
		fmt.Sprintf("Object: %s", obj.Name()),
		fmt.Sprintf("X: %.2f", p.X+t.X),
		fmt.Sprintf("Y: %.2f", p.Y+t.Y),
		fmt.Sprintf("Z: %.2f", p.Z+t.Z),
	}
	for i, line := range lines {
		a.objectInfoTexts[i].SetText(line)
	}
}

func (a *App) SelectedObject() *rende.Object {
	if len(a.objects) == 0 {
		return nil
	}
	idx := a.state.SelectedObjectIdx
	if idx < 0 || idx >= len(a.objects) {
		return nil
	}
	return a.objects[idx]
}

func (a *App) TranslateSelectedObject(dx, dy, dz float32) {
	obj := a.SelectedObject()
	if obj == nil {
		return
	}
	obj.Translate(dx, dy, dz)
}

func (a *App) RotateSelectedObject(angleX, angleY, angleZ float32) {
	obj := a.SelectedObject()
	if obj == nil {
		return
	}
	obj.Rotate(angleX, angleY, angleZ)
}

// Rest of your methods remain the same...
func (a *App) Process() {
	a.controller.ProcessInput(a.window.Core(), a.camera)

	if a.state.IsRotationEnabled {
		for _, obj := range a.objects {
			obj.Rotate(0, a.config.RotationSpeed, 0)
		}
	}

	const blendStep = float32(0.02)
	if a.state.IsTextureEnabled && a.state.TextureBlend < 1.0 {
		a.state.TextureBlend += blendStep
		if a.state.TextureBlend > 1.0 {
			a.state.TextureBlend = 1.0
		}
	} else if !a.state.IsTextureEnabled && a.state.TextureBlend > 0.0 {
		a.state.TextureBlend -= blendStep
		if a.state.TextureBlend < 0.0 {
			a.state.TextureBlend = 0.0
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

	// var y float32 = 32
	// var x float32 = a.screenSize.Width - float32(200)
	var shape []*geom.Vertex
	for _, o := range a.objects {
		shape = append(shape, o.Shape()...)
	}

	// for _, objectInfoElem := range a.objects[0].Info() {
	// 	text := ui.NewText(objectInfoElem, x, y)
	// 	a.ui.AddStaticText(text)
	// 	y += 28
	// }

	a.camera.SetMovementSpeed(shape)
	a.camera.SetPosition(shape)
	a.grid = rende.NewGrid(shape)
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
