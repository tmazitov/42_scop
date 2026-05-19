package objectParsing

import "math"

type vec2 struct{ x, y float32 }

func sub2(a, b vec2) vec2      { return vec2{a.x - b.x, a.y - b.y} }
func cross2(a, b vec2) float32 { return a.x*b.y - a.y*b.x }

// polygonNormal computes polygon normal via Newell's method.
func polygonNormal(indices []uint32, pos func(uint32) [3]float32) [3]float64 {
	var nx, ny, nz float64
	n := len(indices)
	for i := 0; i < n; i++ {
		c := pos(indices[i])
		k := pos(indices[(i+1)%n])
		nx += float64(c[1]-k[1]) * float64(c[2]+k[2])
		ny += float64(c[2]-k[2]) * float64(c[0]+k[0])
		nz += float64(c[0]-k[0]) * float64(c[1]+k[1])
	}
	return [3]float64{nx, ny, nz}
}

// projAxes returns the two coordinate axes for projection (drops the dominant axis).
func projAxes(n [3]float64) (ax, ay int) {
	a0, a1, a2 := math.Abs(n[0]), math.Abs(n[1]), math.Abs(n[2])
	if a0 >= a1 && a0 >= a2 {
		return 1, 2 // drop X
	}
	if a1 >= a0 && a1 >= a2 {
		return 0, 2 // drop Y
	}
	return 0, 1 // drop Z
}

func pointInTri(p, a, b, c vec2) bool {
	d1 := cross2(sub2(b, a), sub2(p, a))
	d2 := cross2(sub2(c, b), sub2(p, b))
	d3 := cross2(sub2(a, c), sub2(p, c))
	hasNeg := d1 < 0 || d2 < 0 || d3 < 0
	hasPos := d1 > 0 || d2 > 0 || d3 > 0
	return !(hasNeg && hasPos)
}

// earClipping triangulates a polygon (convex or concave) using the ear clipping algorithm.
// polygon is a list of vertex indices; pos maps each index to its 3D position.
func earClipping(polygon []uint32, pos func(uint32) [3]float32) []uint32 {
	n := len(polygon)
	if n < 3 {
		return nil
	}
	if n == 3 {
		return []uint32{polygon[0], polygon[1], polygon[2]}
	}

	normal := polygonNormal(polygon, pos)
	ax, ay := projAxes(normal)

	pts := make([]vec2, n)
	for i, idx := range polygon {
		p := pos(idx)
		pts[i] = vec2{p[ax], p[ay]}
	}

	// Signed area determines winding (positive = CCW, negative = CW).
	var area float32
	for i := 0; i < n; i++ {
		area += cross2(pts[i], pts[(i+1)%n])
	}

	rem := make([]int, n)
	for i := range rem {
		rem[i] = i
	}

	out := make([]uint32, 0, (n-2)*3)

	for len(rem) > 3 {
		m := len(rem)
		earFound := false

		for i := 0; i < m; i++ {
			ip := rem[(i-1+m)%m]
			ic := rem[i]
			in_ := rem[(i+1)%m]
			p0, p1, p2 := pts[ip], pts[ic], pts[in_]

			// Convexity: signed area of ear triangle must match polygon winding.
			cp := cross2(sub2(p1, p0), sub2(p2, p0))
			if (area > 0 && cp <= 0) || (area < 0 && cp >= 0) {
				continue
			}

			// No other remaining vertex inside this ear triangle.
			isEar := true
			for j := 0; j < m; j++ {
				if j == (i-1+m)%m || j == i || j == (i+1)%m {
					continue
				}
				if pointInTri(pts[rem[j]], p0, p1, p2) {
					isEar = false
					break
				}
			}
			if !isEar {
				continue
			}

			out = append(out, polygon[ip], polygon[ic], polygon[in_])
			rem = append(rem[:i], rem[i+1:]...)
			earFound = true
			break
		}

		if !earFound {
			// Degenerate polygon (e.g. collinear vertices): fall back to fan.
			for i := 1; i < len(rem)-1; i++ {
				out = append(out, polygon[rem[0]], polygon[rem[i]], polygon[rem[i+1]])
			}
			return out
		}
	}

	out = append(out, polygon[rem[0]], polygon[rem[1]], polygon[rem[2]])
	return out
}
