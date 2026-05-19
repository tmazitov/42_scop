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
	filePath        string
	materialStorage *materialStorage
	vertexStorage   *vertexStorage
	name            string
	hasNormals      bool
	smoothGroup     int // 0 = flat shading, >0 = smooth shading group
}

func newObjectParsingProcess(name, filePath string, materialStorage *materialStorage, vertexStorage *vertexStorage) *objectParsingProcess {
	return &objectParsingProcess{
		name:            name,
		filePath:        filePath,
		materialStorage: materialStorage,
		vertexStorage:   vertexStorage,
		verticesCache:   make(map[vertexKey]uint32),
		smoothGroup:     1, // default: smooth shading on
	}
}

func (o *objectParsingProcess) ToObject() (*rende.Object, error) {
	if !o.hasNormals {
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

