package appx

type controller struct {
	app 		*App
	lastX    	float64
    lastY    	float64
    firstMouse 	bool
	lastFrame	float32
}

func newController(app *App) *controller {
	return &controller{
		app: app,
		lastX: 0,
		lastY: 0,
		firstMouse: false,
	}
}
