package objectParsing

// vertexStorage holds the raw global vertex arrays shared across all objects in one .obj file.
// In OBJ format, v/vt/vn indices are global — they don't reset per 'o' declaration.
type vertexStorage struct {
	coords   [][3]float32
	textures [][2]float32
	normals  [][3]float32
}

func newVertexStorage() *vertexStorage {
	return &vertexStorage{}
}