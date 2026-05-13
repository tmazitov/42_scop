package geom

import "math"

type normCacheKey struct {
	vidx     uint32
	nx, ny, nz int32
}

func faceNormal(v0, v1, v2 *Pos) [3]float64 {
	e1x := float64(v1.X - v0.X)
	e1y := float64(v1.Y - v0.Y)
	e1z := float64(v1.Z - v0.Z)
	e2x := float64(v2.X - v0.X)
	e2y := float64(v2.Y - v0.Y)
	e2z := float64(v2.Z - v0.Z)
	nx := e1y*e2z - e1z*e2y
	ny := e1z*e2x - e1x*e2z
	nz := e1x*e2y - e1y*e2x
	length := math.Sqrt(nx*nx + ny*ny + nz*nz)
	if length == 0 {
		return [3]float64{}
	}
	return [3]float64{nx / length, ny / length, nz / length}
}

// ComputeSmoothNormals computes per-vertex normals with crease angle support.
// Faces whose normals differ by more than creaseAngleCos (cos of max angle) are treated
// as hard edges — the vertex is split so each face group gets its own normal.
// creaseAngleCos=0.5 means 60° crease angle (good default for hard-surface models).
func ComputeSmoothNormals(vertices []*Vertex, indices []uint32, creaseAngleCos float64) ([]*Vertex, []uint32) {
	triCount := len(indices) / 3

	// Step 1: normalized face normal for each triangle.
	faceNormals := make([][3]float64, triCount)
	for i := 0; i < triCount; i++ {
		i0, i1, i2 := indices[i*3], indices[i*3+1], indices[i*3+2]
		faceNormals[i] = faceNormal(vertices[i0].Pos, vertices[i1].Pos, vertices[i2].Pos)
	}

	// Step 2: adjacency — which triangles share each vertex
	vertexToTris := make([][]int, len(vertices))
	for i := 0; i < triCount; i++ {
		for k := 0; k < 3; k++ {
			vidx := indices[i*3+k]
			vertexToTris[vidx] = append(vertexToTris[vidx], i)
		}
	}

	// Step 3: for each (triangle, vertex) pair compute smoothed normal limited by crease angle,
	// then deduplicate vertices that ended up with the same normal.
	const quantScale = 32767
	cache := make(map[normCacheKey]uint32)
	newVertices := make([]*Vertex, 0, len(vertices))
	newIndices := make([]uint32, len(indices))

	for i := 0; i < triCount; i++ {
		fn := faceNormals[i]
		for k := 0; k < 3; k++ {
			vidx := indices[i*3+k]

			var acc [3]float64
			for _, adjTri := range vertexToTris[vidx] {
				an := faceNormals[adjTri]
				dot := fn[0]*an[0] + fn[1]*an[1] + fn[2]*an[2]
				if dot >= creaseAngleCos {
					acc[0] += an[0]
					acc[1] += an[1]
					acc[2] += an[2]
				}
			}

			length := math.Sqrt(acc[0]*acc[0] + acc[1]*acc[1] + acc[2]*acc[2])
			var norm [3]float32
			if length > 0 {
				norm = [3]float32{
					float32(acc[0] / length),
					float32(acc[1] / length),
					float32(acc[2] / length),
				}
			}

			key := normCacheKey{
				vidx,
				int32(norm[0] * quantScale),
				int32(norm[1] * quantScale),
				int32(norm[2] * quantScale),
			}
			if existingIdx, ok := cache[key]; ok {
				newIndices[i*3+k] = existingIdx
			} else {
				newIdx := uint32(len(newVertices))
				v := vertices[vidx].Clone()
				v.SetNormByVector(norm)
				newVertices = append(newVertices, v)
				cache[key] = newIdx
				newIndices[i*3+k] = newIdx
			}
		}
	}

	return newVertices, newIndices
}

func NormalizeVertices(vertices []*Vertex) []*Vertex {
	if len(vertices) == 0 {
		return nil
	}

	// 1. AABB
	min := NewVertex([3]float32{
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
	center := NewVertex([3]float32{
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

	result := make([]*Vertex, 0, len(vertices))

	for i := range vertices {
		result = append(result, NewVertex([3]float32{
			(vertices[i].Pos.X - center.Pos.X) * scale,
			(vertices[i].Pos.Y - center.Pos.Y) * scale,
			(vertices[i].Pos.Z - center.Pos.Z) * scale,
		}))
	}

	return result
}
