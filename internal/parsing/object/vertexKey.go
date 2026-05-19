package objectParsing

type vertexKey [3]int

func newVertexKey(slots [3]int) vertexKey {
	return vertexKey(slots)
}

func (k *vertexKey) Pos() int {
	return k[0]
}

func (k *vertexKey) Texture() int {
	return k[1]
}

func (k *vertexKey) Norm() int {
	return k[2]
}