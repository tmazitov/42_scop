package geom

import (
	"fmt"
)

type Vertex struct {
	Pos		*Pos 
	Norm	*Pos
	U, V	float32  // Texture coordinates (ADD THESE)
}

func NewVertex(vector [3]float32) *Vertex {
	return &Vertex{
		Pos: &Pos{
			X: vector[0],
			Y: vector[1],
			Z: vector[2],
		},
		Norm: nil,
		U: 0,
		V: 0,
	}
}

func (p *Vertex) Vector() []float32 {
	return []float32{
		p.Pos.X,
		p.Pos.Y,
		p.Pos.Z,
	}
}

func (p *Vertex) ToString() string {
	return fmt.Sprintf("Vertex %v", p.Vector())
}

func (p *Vertex) SetNorm(norm *Vertex) {
	p.Norm = &Pos{
		X: norm.Pos.X,
		Y: norm.Pos.Y,
		Z: norm.Pos.Z,
	}
} 

func (p *Vertex) SetNormByVector(norm [3]float32) {
	p.Norm = &Pos{
		X: norm[0],
		Y: norm[1],
		Z: norm[2],
	}
}

func (p *Vertex) SetTextureCoords(texture [2]float32) {
	p.U = texture[0]
	p.V = texture[1]
}