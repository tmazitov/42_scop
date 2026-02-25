package appx

import (
	"github.com/tmazitov/42_scop/internal/geom"
)

func calculateCameraPosition(vertices []*geom.Vertex) (geom.Pos) {
	minX, maxX := vertices[0].Pos.X, vertices[0].Pos.X
    minY, maxY := vertices[0].Pos.Y, vertices[0].Pos.Y
    minZ, maxZ := vertices[0].Pos.Z, vertices[0].Pos.Z

    for _, v := range vertices {
        if v.Pos.X < minX { minX = v.Pos.X }
        if v.Pos.X > maxX { maxX = v.Pos.X }
        if v.Pos.Y < minY { minY = v.Pos.Y }
        if v.Pos.Y > maxY { maxY = v.Pos.Y }
        if v.Pos.Z < minZ { minZ = v.Pos.Z }
        if v.Pos.Z > maxZ { maxZ = v.Pos.Z }
    }

	maxDifference := (maxX - minX + maxY - minY + maxZ - minZ) / 3
	
	return geom.Pos{
		X: (minX + maxX) / 2,
		Y: (minY + maxY) / 2,
		Z: (minZ + maxZ) / 2 + maxDifference,
	}
}

func calculateCameraSpeed(vertices []*geom.Vertex) float32 {
	minX, maxX := vertices[0].Pos.X, vertices[0].Pos.X
	minY, maxY := vertices[0].Pos.Y, vertices[0].Pos.Y
	minZ, maxZ := vertices[0].Pos.Z, vertices[0].Pos.Z

    for _, v := range vertices {
        if v.Pos.X < minX { minX = v.Pos.X }
        if v.Pos.X > maxX { maxX = v.Pos.X }
        if v.Pos.Y < minY { minY = v.Pos.Y }
        if v.Pos.Y > maxY { maxY = v.Pos.Y }
        if v.Pos.Z < minZ { minZ = v.Pos.Z }
        if v.Pos.Z > maxZ { maxZ = v.Pos.Z }
    }

	return (maxX - minX + maxY - minY + maxZ - minZ) / 3 * 0.5
}