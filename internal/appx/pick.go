package appx

import (
	"math"

	"github.com/tmazitov/42_scop/internal/appx/camera"
	"github.com/tmazitov/42_scop/internal/geom"
	"github.com/tmazitov/42_scop/internal/rende"
)

func (a *App) SelectObjectAt(x, y float32) {
	if len(a.objects) == 0 {
		return
	}
	origin := a.camera.Position
	dir := screenRay(x, y, a.screenSize, a.camera)

	bestIdx := -1
	bestDist := float32(math.MaxFloat32)
	for i, obj := range a.objects {
		t, hit := rayAABBHit(origin, dir, obj.Shape())
		if hit && t < bestDist {
			bestDist = t
			bestIdx = i
		}
	}
	if bestIdx >= 0 {
		a.state.SelectedObjectIdx = bestIdx
	}
}

func screenRay(x, y float32, screen rende.ScreenSize, cam *camera.Camera) geom.Vec3 {
	ndcX := (2*x/screen.Width) - 1
	ndcY := 1 - (2*y/screen.Height)
	aspect := screen.Width / screen.Height
	tanHalfFov := float32(math.Tan(float64(geom.DegToRad(cam.Zoom)) * 0.5))

	dir := cam.Front.
		Add(cam.Right.Mul(ndcX * aspect * tanHalfFov)).
		Add(cam.Up.Mul(ndcY * tanHalfFov))
	return dir.Normalize()
}

func rayAABBHit(origin, dir geom.Vec3, shape []*geom.Vertex) (float32, bool) {
	if len(shape) == 0 {
		return 0, false
	}
	minX, maxX := shape[0].Pos.X, shape[0].Pos.X
	minY, maxY := shape[0].Pos.Y, shape[0].Pos.Y
	minZ, maxZ := shape[0].Pos.Z, shape[0].Pos.Z
	for _, v := range shape {
		if v.Pos.X < minX { minX = v.Pos.X }
		if v.Pos.X > maxX { maxX = v.Pos.X }
		if v.Pos.Y < minY { minY = v.Pos.Y }
		if v.Pos.Y > maxY { maxY = v.Pos.Y }
		if v.Pos.Z < minZ { minZ = v.Pos.Z }
		if v.Pos.Z > maxZ { maxZ = v.Pos.Z }
	}

	tmin := float32(math.Inf(-1))
	tmax := float32(math.Inf(1))

	type slab struct{ orig, d, lo, hi float32 }
	slabs := [3]slab{
		{origin.X(), dir.X(), minX, maxX},
		{origin.Y(), dir.Y(), minY, maxY},
		{origin.Z(), dir.Z(), minZ, maxZ},
	}
	for _, s := range slabs {
		if s.d != 0 {
			t1 := (s.lo - s.orig) / s.d
			t2 := (s.hi - s.orig) / s.d
			if t1 > t2 { t1, t2 = t2, t1 }
			if t1 > tmin { tmin = t1 }
			if t2 < tmax { tmax = t2 }
		} else if s.orig < s.lo || s.orig > s.hi {
			return 0, false
		}
	}
	if tmin > tmax || tmax < 0 {
		return 0, false
	}
	t := tmin
	if t < 0 {
		t = tmax
	}
	return t, true
}
