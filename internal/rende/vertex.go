package rende

import (
	"fmt"
)

type Vertex struct {
	Pos *Pos
}

func NewVertex(vector [3]float32) *Vertex {
	return &Vertex{
		Pos: &Pos{
			X: vector[0],
			Y: vector[1],
			Z: vector[2],
		},
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
