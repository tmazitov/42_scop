package geom

import "math"

type Mat4 [16]float32

func IdentityMat4() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func Perspective(fov, aspect, near, far float32) Mat4 {
	f := 1.0 / float32(math.Tan(float64(DegToRad(fov))/2))
	return Mat4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (far + near) / (near - far), -1,
		0, 0, (2 * far * near) / (near - far), 0,
	}
}	

func LookAtV(eye, target, up Vec3) Mat4 {
	f := target.Sub(eye).Normalize()
	s := f.Cross(up.Normalize()).Normalize()
	u := s.Cross(f)

	return Mat4{
		s.X(), u.X(), -f.X(), 0,
		s.Y(), u.Y(), -f.Y(), 0,
		s.Z(), u.Z(), -f.Z(), 0,
		-s.Dot(eye), -u.Dot(eye), f.Dot(eye), 1,
	}
}