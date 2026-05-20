package geom

import (
	"fmt"
	"math"
)

type Vec3 [3]float32

func (v Vec3) X() float32 {
	return v[0]
}

func (v Vec3) Y() float32 {
	return v[1]
}

func (v Vec3) Z() float32 {
	return v[2]
}

func (v Vec3) ToArray() [3]float32 {
	return [3]float32{v.X(), v.Y(), v.Z()}
}

func (v Vec3) ToString() string	{
	return fmt.Sprintf("Vec3(%f, %f, %f)", v.X(), v.Y(), v.Z())
}

func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{
		v.X() + other.X(),
		v.Y() + other.Y(),
		v.Z() + other.Z(),
	}
}

func (v Vec3) Sub(other Vec3) Vec3 {
	return Vec3{
		v.X() - other.X(),
		v.Y() - other.Y(),
		v.Z() - other.Z(),
	}
}

func (v Vec3) Mul(scalar float32) Vec3 {
	return Vec3{
		v.X() * scalar,
		v.Y() * scalar,
		v.Z() * scalar,
	}
}

func (v Vec3) Div(scalar float32) Vec3 {
	return Vec3{
		v.X() / scalar,
		v.Y() / scalar,
		v.Z() / scalar,
	}
}

func (v Vec3) Normalize() Vec3 {
	length := v.Length()
	if length == 0 {
		return Vec3{0, 0, 0}
	}
	return v.Div(length)
}

func (v Vec3) Length() float32 {
	return float32(math.Sqrt(float64(v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z())))
}

func (v Vec3) Cross(other Vec3) Vec3 {
	return Vec3{
		v.Y()*other.Z() - v.Z()*other.Y(),
		v.Z()*other.X() - v.X()*other.Z(),
		v.X()*other.Y() - v.Y()*other.X(),
	}
}

func (v Vec3) Dot(other Vec3) float32 {
	return v.X()*other.X() + v.Y()*other.Y() + v.Z()*other.Z()
}