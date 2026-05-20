package camera

import (
	"github.com/tmazitov/42_scop/internal/geom"
)

func (cam *Camera) SetMovementSpeed(vertices []*geom.Vertex) float32 {
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

	cam.MovementSpeed = (maxX - minX + maxY - minY + maxZ - minZ) / 3 * 0.5

	return cam.MovementSpeed
}