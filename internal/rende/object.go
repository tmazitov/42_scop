package rende

type Object struct {
	name string
	points []Point
}

func NewObject(name string, points... Point) *Object {
	return &Object{
		name: name,
		points: points,
	}
}
