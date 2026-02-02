package rende

type Object struct {
	name string
	pos *Point
	shape []*Point
	indices []uint32
}

func NewObject(name string, pos *Point, shape []*Point) *Object {
	return &Object{
		name: name,
		shape: shape,
		pos: pos,
		indices: []uint32{
			0, 1, 2, // First triangle
			1, 2, 3, // Second triangle
		},
	}
}

func (o *Object) Shape() []*Point{
	return o.shape
}

func (o *Object) Pos() *Point{
	return o.pos
}

func (o *Object) VAO(screen ScreenSize) uint32 {
	return MakeVao(screen, o.indices, o.shape)
}

func (o *Object) NodeCount() int32 {
	return int32(len(o.shape))
}

func (o *Object) IndicesCount() int32 {
	return int32(len(o.indices))
}