package controller

type Controller struct {
	app			App		
	lastX       float64
	lastY       float64
	firstMouse  bool
	deltaTime   float32
	lastFrame   float32
}

func NewController(app App) *Controller {
	return &Controller{
		app:     	app,
		lastX:      float64(app.ScreenSize().Width) / 2,
		lastY:      float64(app.ScreenSize().Height) / 2,
		firstMouse: true,
		deltaTime:  0.0,
		lastFrame:  0.0,
	}
}



// Bind mouse control to window
func (c *Controller) BindMouseControl() {
	window := c.app.Window().Core()
	
	// Capture and hide cursor
	// window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)

	// Set cursor position callback
	window.SetScrollCallback(c.scrollCallback)
	window.SetCursorPosCallback(c.mouseMoveCallback)
	window.SetMouseButtonCallback(c.mouseClickCallback)
}