package geom

import (
	"fmt"
	"math"
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

func (p *Vertex) Clone() *Vertex {

	var result = &Vertex{}

	if p.Pos != nil {
		result.Pos = &Pos{
			X : p.Pos.X,
			Y : p.Pos.Y,
			Z : p.Pos.Z,
		}
	}

	if p.Norm != nil {
		result.Norm = &Pos{
			X : p.Norm.X,
			Y : p.Norm.Y,
			Z : p.Norm.Z,
		}
	}

	result.U = p.U
	result.V = p.V

	return result
}

func rotateVec(x, y, z, angleX, angleY, angleZ float32) (float32, float32, float32) {
	cosX := float32(math.Cos(float64(angleX)))
	sinX := float32(math.Sin(float64(angleX)))
	y1 := y*cosX - z*sinX
	z1 := y*sinX + z*cosX

	cosY := float32(math.Cos(float64(angleY)))
	sinY := float32(math.Sin(float64(angleY)))
	x2 := x*cosY + z1*sinY
	z2 := -x*sinY + z1*cosY

	cosZ := float32(math.Cos(float64(angleZ)))
	sinZ := float32(math.Sin(float64(angleZ)))
	x3 := x2*cosZ - y1*sinZ
	y3 := x2*sinZ + y1*cosZ
	return x3, y3, z2
}

func (v *Vertex) Rotate(angleX, angleY, angleZ float32) *Vertex {
	v.Pos.X, v.Pos.Y, v.Pos.Z = rotateVec(v.Pos.X, v.Pos.Y, v.Pos.Z, angleX, angleY, angleZ)
	if v.Norm != nil {
		v.Norm.X, v.Norm.Y, v.Norm.Z = rotateVec(v.Norm.X, v.Norm.Y, v.Norm.Z, angleX, angleY, angleZ)
	}
	return v
}