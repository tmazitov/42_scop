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

func uploadTextTexture(text string) (id uint32, w, h int32) {
	face := basicfont.Face7x13
	width := font.MeasureString(face, text).Ceil()
	if width == 0 {
		return 0, 0, 0
	}
	height := 16

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), image.Transparent, image.Point{}, draw.Src)
	(&font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.White),
		Face: face,
		Dot:  fixed.P(0, 13),
	}).DrawString(text)

	gl.GenTextures(1, &id)
	gl.BindTexture(gl.TEXTURE_2D, id)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(width), int32(height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))
	gl.BindTexture(gl.TEXTURE_2D, 0)

	return id, int32(width), int32(height)
}

func NewText(text string, x, y float32) *Text {
	id, w, h := uploadTextTexture(text)
	return &Text{
		textureID: id,
		width:     w,
		height:    h,
		pos:       geom.Pos{X: x, Y: y},
	}
}

func (t *Text) SetText(text string) {
	if t.textureID != 0 {
		gl.DeleteTextures(1, &t.textureID)
	}
	t.textureID, t.width, t.height = uploadTextTexture(text)
}

func (b *Text) Draw() {
	if b.textureID == 0 {
		return
	}
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