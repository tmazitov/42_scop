package rende

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/tmazitov/42_scop/internal/geom"
)

type Grid struct {
	y         float32
	halfSize  float32
	step      float32
}

func NewGrid(vertices []*geom.Vertex) *Grid {
	if len(vertices) == 0 {
		return &Grid{y: 0, halfSize: 10, step: 1}
	}

	minX, maxX := vertices[0].Pos.X, vertices[0].Pos.X
	minY := vertices[0].Pos.Y
	minZ, maxZ := vertices[0].Pos.Z, vertices[0].Pos.Z

	for _, v := range vertices[1:] {
		if v.Pos.X < minX { minX = v.Pos.X }
		if v.Pos.X > maxX { maxX = v.Pos.X }
		if v.Pos.Y < minY { minY = v.Pos.Y }
		if v.Pos.Z < minZ { minZ = v.Pos.Z }
		if v.Pos.Z > maxZ { maxZ = v.Pos.Z }
	}

	dx := maxX - minX
	dz := maxZ - minZ
	maxDim := dx
	if dz > maxDim { maxDim = dz }

	step := float32(1.0)
	for step < maxDim/10 { step *= 10 }
	for step > maxDim/2  { step /= 10 }
	if step < 0.01 { step = 0.01 }

	halfSize := maxDim * 2.5
	if halfSize < step*5 { halfSize = step * 5 }

	return &Grid{
		y:        minY,
		halfSize: halfSize,
		step:     step,
	}
}

func (g *Grid) Draw() {
	gl.Disable(gl.LIGHTING)
	gl.Disable(gl.TEXTURE_2D)
	gl.LineWidth(1.0)

	gl.Begin(gl.LINES)
	half := g.halfSize
	step := g.step
	y := g.y

	for x := -half; x <= half+step*0.5; x += step {
		if abs32(x) < step*0.1 {
			continue // skip axis lines, drawn separately
		}
		gl.Color4f(0.35, 0.35, 0.35, 1.0)
		gl.Vertex3f(x, y, -half)
		gl.Vertex3f(x, y, half)
	}

	for z := -half; z <= half+step*0.5; z += step {
		if abs32(z) < step*0.1 {
			continue
		}
		gl.Color4f(0.35, 0.35, 0.35, 1.0)
		gl.Vertex3f(-half, y, z)
		gl.Vertex3f(half, y, z)
	}

	// X axis — red
	gl.Color4f(0.8, 0.2, 0.2, 1.0)
	gl.Vertex3f(-half, y, 0)
	gl.Vertex3f(half, y, 0)

	// Z axis — blue
	gl.Color4f(0.2, 0.2, 0.8, 1.0)
	gl.Vertex3f(0, y, -half)
	gl.Vertex3f(0, y, half)

	gl.End()

	gl.Enable(gl.LIGHTING)
}

func abs32(v float32) float32 {
	if v < 0 { return -v }
	return v
}
