package main

import (
	"log"
	"runtime"

	// OR: github.com/go-gl/gl/v2.1/gl
	"github.com/tmazitov/42_scop/internal/appx"
	"github.com/tmazitov/42_scop/internal/parsing"
)

// makeVao initializes and returns a vertex array from the points provided.

func main() {
	runtime.LockOSThread()

	config, err := SetupConfig()
	if err != nil {
		log.Fatal(err)
	}
	app, err := appx.NewApp(config.ToAppConfig())
	if err != nil {
		log.Fatal(err)
	}
	defer app.Close()

	obj, err := parsing.ParseObj(config.ObjectPath)
	if err != nil {
		log.Fatal(err)
	}
	app.AddObject(obj)

	render(app, config)
}




// initGlfw initializes glfw and returns a Window to use.
