package rende

type Object struct {
	name string
	shape []*Vertex
	indices []uint32
}

func NewObject(name string, shape []*Vertex, indices []uint32) *Object {
	return &Object{
		name: name,
		shape: shape,
		indices: indices,
	}
}

func (o *Object) Shape() []*Vertex{
	return o.shape
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