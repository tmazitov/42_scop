package ui

import (
	"github.com/tmazitov/42_scop/internal/rende"
	"github.com/tmazitov/42_scop/internal/clr"
	"github.com/go-gl/gl/v2.1/gl"
	// "fmt"
)

type Button struct {
	id 		int
	onClick func()
	pos		*rende.Pos
	color	*clr.Color
	width 	float32
	height	float32
}

var buttonIdCounter int = 0

var buttonDefaultColor = clr.NewColor() // rgba : 0 0 0 0

func NewButton() *Button {

	buttonIdCounter++

	return &Button{
		id: buttonIdCounter,
		pos: nil,
		width: 0,
		height: 0,
		onClick: nil,
		color: buttonDefaultColor,
	}
}

func (b *Button) Draw() {
	b.color.Enable()

    gl.Begin(gl.QUADS)

    gl.Vertex2f(b.pos.X, b.pos.Y)
    gl.Vertex2f(b.pos.X + b.width, b.pos.Y)
    gl.Vertex2f(b.pos.X + b.width, b.pos.Y + b.height)
    gl.Vertex2f(b.pos.X, b.pos.Y + b.height)
	
    gl.End()
}

func (b *Button) SetColor(color *clr.Color) *Button {
	b.color = color
	return b
}

func (b *Button) SetPos(pos *rende.Pos) *Button{
	b.pos = pos
	return b
} 

func (b *Button) SetHeight(height float32) *Button{
	b.height = height
	return b
}

func (b *Button) SetWidth(width float32) *Button{
	b.width = width
	return b
}

func (b *Button) SetSize(height, width float32) *Button{
	b.width = width
	b.height = height
	return b
}

func (b *Button) SetOnClick(onClickHandler func ()) *Button{
	b.onClick = onClickHandler
	return b
}
