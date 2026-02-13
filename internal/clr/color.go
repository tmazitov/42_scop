package clr

import (
	"github.com/go-gl/gl/v2.1/gl"
	// "fmt"
	"time"
	"math/rand"
)

type Color struct {
	R float32
	G float32
	B float32
	Alpha float32
}

func NewColor(params ...int) *Color {

	var (
		r int
		g int
		b int
		a int = 0
		paramsCount = len(params)
	)

	if paramsCount > 0 {
		r = params[0]
	}

	if paramsCount > 1 {
		g = params[1]
	}

	if paramsCount > 2 {
		b = params[2]
	}

	if paramsCount > 3 {
		a = params[3]
	}

	color := &Color{
		R: float32(r) / 255.0,
		G: float32(g) / 255.0,
		B: float32(b) / 255.0,
		Alpha: (255.0 - float32(a)) / 255.0,
	}
    rand.Seed(time.Now().Unix())
    
	return color
}

func NewColorF(params ...float32) *Color {
	var (
		r,g,b,a float32
		paramsCount = len(params)
	)



	if paramsCount > 0 {
		r = params[0]
	}

	if paramsCount > 1 {
		g = params[1]
	}

	if paramsCount > 2 {
		b = params[2]
	}

	if paramsCount > 3 {
		a = params[3]
	}

	return &Color{
		R: r,
		G: g,
		B: b,
		Alpha: a,
	}
}

func (c *Color) Apply() {
	gl.Color4f(c.R, c.G, c.B, c.Alpha)
}

func (c *Color) Vector() []float32 {
	return []float32{c.R, c.G, c.B, c.Alpha}
}