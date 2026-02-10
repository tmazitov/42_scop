package clr

import (
	"github.com/go-gl/gl/v2.1/gl"
	"fmt"
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
		b = params[3]
	}

	color := &Color{
		R: float32(r) / 255.0,
		G: float32(g) / 255.0,
		B: float32(b) / 255.0,
		Alpha: (255.0 - float32(a)) / 255.0,
	}
    rand.Seed(time.Now().Unix())
    
	go func () {
		for {

			plus := rand.Intn(3)

			switch plus {
			case 0:
				color.R += 0.1
			case 1:
				color.G += 0.1
			case 2:
				color.B += 0.1
			}

			if color.R > 1.0 {
				color.R = 0.0
			}
			if color.G > 1.0 {
				color.G = 0.0
			}
			if color.B > 1.0 {
				color.B = 0.0
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()

	return color
}

func (c *Color) Enable() {
	fmt.Println(c.R, c.G, c.B, c.Alpha)
	gl.Color4f(c.R, c.G, c.B, c.Alpha)
}