package ui

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/tmazitov/42_scop/internal/geom"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"github.com/go-gl/gl/v2.1/gl"
)

type Text struct {
	textureID uint32
	width     int32
	height    int32
	pos		  geom.Pos
}

func NewText(text string, x, y float32) (*Text, error) {
	// Measure text width automatically
	face := basicfont.Face7x13
	width := font.MeasureString(face, text).Ceil()
	height := 16 // Fixed height for basicfont

	// Create RGBA image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill with transparent background
	draw.Draw(img, img.Bounds(), image.Transparent, image.Point{}, draw.Src)

	// Draw text onto image
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.White),
		Face: face,
		Dot:  fixed.P(0, 13), // baseline position
	}
	d.DrawString(text)

	// Create OpenGL texture from image
	var textureID uint32
	gl.GenTextures(1, &textureID)
	gl.BindTexture(gl.TEXTURE_2D, textureID)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(width),
		int32(height),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(img.Pix),
	)

	gl.BindTexture(gl.TEXTURE_2D, 0)

	return &Text{
		textureID: textureID,
		width:     int32(width),
		height:    int32(height),
		pos:	   geom.Pos{
			X: x,
			Y: y,
		},
	}, nil
}

func (b *Text) Draw() {
	w := float32(b.width)
	h := float32(b.height)

	gl.Enable(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, b.textureID)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	// Disable lighting so text color isn't affected by scene lights
	gl.Disable(gl.LIGHTING)

	gl.Color4f(1.0, 1.0, 1.0, 1.0) // White color

	gl.Begin(gl.QUADS)
	gl.TexCoord2f(0, 0); gl.Vertex2f(b.pos.X, b.pos.Y)
	gl.TexCoord2f(1, 0); gl.Vertex2f(b.pos.X+w, b.pos.Y)
	gl.TexCoord2f(1, 1); gl.Vertex2f(b.pos.X+w, b.pos.Y+h)
	gl.TexCoord2f(0, 1); gl.Vertex2f(b.pos.X, b.pos.Y+h)
	gl.End()

	gl.Disable(gl.TEXTURE_2D)
	gl.Disable(gl.BLEND)
}

func (b *Text) Cleanup() {
	if b.textureID != 0 {
		gl.DeleteTextures(1, &b.textureID)
		b.textureID = 0
	}
}