package geom

import (
	"fmt"
)

type Vertex struct {
	Pos *Pos
	Norm *Pos
}

func NewVertex(vector [3]float32) *Vertex {
	return &Vertex{
		Pos: &Pos{
			X: vector[0],
			Y: vector[1],
			Z: vector[2],
		},
		Norm: nil,
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
