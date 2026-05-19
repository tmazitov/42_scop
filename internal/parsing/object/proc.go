package objectParsing

import (
	"github.com/tmazitov/42_scop/internal/geom"
	"github.com/tmazitov/42_scop/internal/rende"
	// "fmt"
)

type objectParsingProcess struct {
	vertices        []*geom.Vertex
	verticesCache   map[vertexKey]uint32
	indices         []uint32
	materials       []*rende.Material
	currentMaterial *rende.Material
	filePath        string
	materialStorage *materialStorage
	vertexStorage   *vertexStorage
	name            string
}

func newObjectParsingProcess(name, filePath string, materialStorage *materialStorage, vertexStorage *vertexStorage) *objectParsingProcess {
	return &objectParsingProcess{
		name:            name,
		filePath:        filePath,
		materialStorage: materialStorage,
		vertexStorage:   vertexStorage,
		verticesCache:   make(map[vertexKey]uint32),
	}
}

func (o *objectParsingProcess) ToObject() (*rende.Object, error) {
	if len(o.vertexStorage.normals) == 0 {
		o.vertices, o.indices = geom.ComputeSmoothNormals(o.vertices, o.indices, 0.5)
	}

	object := rende.NewObject(o.name).
		SetShape(o.vertices).
		SetIndices(o.indices).
		SetMaterials(o.materials)

	return object, nil
}

func (o *objectParsingProcess) IsEmpty() bool {
	return len(o.indices) == 0
}

type vertexKey [3]int

func newVertexKey(values []uint32) vertexKey {
	var result = vertexKey{-1, -1, -1}

	for index, value := range values {
		result[index] = int(value)
	}
	return result
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
