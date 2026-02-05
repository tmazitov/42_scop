package parsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
	
)

func NormalizeVertices(vertices []*rende.Vertex) []*rende.Vertex {
	if len(vertices) == 0 {
		return nil
	}

	// 1. AABB
	min := rende.NewVertex([3]float32{
		vertices[0].Pos.X, 
		vertices[0].Pos.Y, 
		vertices[0].Pos.Z,
	})
	max := min

	for _, v := range vertices {
		p := v.Pos
		if p.X < min.Pos.X { min.Pos.X = p.X }
		if p.Y < min.Pos.Y { min.Pos.Y = p.Y }
		if p.Z < min.Pos.Z { min.Pos.Z = p.Z }

		if p.X > max.Pos.X { max.Pos.X = p.X }
		if p.Y > max.Pos.Y { max.Pos.Y = p.Y }
		if p.Z > max.Pos.Z { max.Pos.Z = p.Z }
	}

	// 2. Центр
	center := rende.NewVertex([3]float32{
		(min.Pos.X + max.Pos.X) * 0.5,
		(min.Pos.Y + max.Pos.Y) * 0.5,
		(min.Pos.Z + max.Pos.Z) * 0.5,
	})

	// 3. Масштаб (по самой большой оси)
	sizeX := max.Pos.X - min.Pos.X
	sizeY := max.Pos.Y - min.Pos.Y
	sizeZ := max.Pos.Z - min.Pos.Z

	maxSize := sizeX
	if sizeY > maxSize { maxSize = sizeY }
	if sizeZ > maxSize { maxSize = sizeZ }

	if maxSize == 0 {
		return vertices// вырожденная модель
	}

	scale := 2.0 / maxSize // -> влезет в [-1,1]

	// 4. Применяем

	result := make([]*rende.Vertex, 0, len(vertices))

	for i := range vertices {
		result = append(result, rende.NewVertex([3]float32{
			(vertices[i].Pos.X - center.Pos.X) * scale,
			(vertices[i].Pos.Y - center.Pos.Y) * scale,
			(vertices[i].Pos.Z - center.Pos.Z) * scale,
		}))
	}

	return result
}
