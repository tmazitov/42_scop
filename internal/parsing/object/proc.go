package objectParsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
	"github.com/tmazitov/42_scop/internal/geom"
	// "fmt"
)

type objectParsingProcess struct {
	vertices			[]*geom.Vertex
	verticesTextures	[][2]float32
	verticesNormals		[][3]float32
	verticesCoords		[][3]float32
	verticesCache      	map[vertexKey]uint32
	indices				[]uint32
	materials			[]*rende.Material
	currentMaterial		*rende.Material
	filePath			string
}

func newObjectParsingProcess(filePath string) *objectParsingProcess {
	return &objectParsingProcess{
		filePath: filePath,
		verticesCache: make(map[vertexKey]uint32),
	}
}

func (o *objectParsingProcess) Prepare() error {
	if len(o.verticesNormals) == 0 {
		o.vertices, o.indices = geom.ComputeSmoothNormals(o.vertices, o.indices, 0.5)
	}
	return nil
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