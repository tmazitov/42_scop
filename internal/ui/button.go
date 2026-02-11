package ui

import (
	"github.com/tmazitov/42_scop/internal/geom"
	"github.com/tmazitov/42_scop/internal/clr"
	"github.com/go-gl/gl/v2.1/gl"
	// "fmt"
)

type Button struct {
	id 		int
	onClickHandler ElementHandleFunc
	pos		*geom.Pos
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
		onClickHandler: nil,
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

func (b *Button) SetPos(pos *geom.Pos) *Button{
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

func (b *Button) SetOnClick(onClickHandler ElementHandleFunc) *Button{
	b.onClickHandler = onClickHandler
	return b
}

func (b *Button) IsPressed(xpos, ypos float32) bool {

	return xpos >= b.pos.X &&
			xpos <= b.pos.X + b.width &&
			ypos >= b.pos.Y &&
			ypos <= b.pos.Y + b.height 
}

func (b *Button) OnClickHandler() ElementHandleFunc {
	return b.onClickHandler
}