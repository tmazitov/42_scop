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
	indices				[]uint32
	materials			[]*rende.Material
	filePath			string
}

func (o *objectParsingProcess) Prepare() error {

	// fmt.Println("lens", len(o.vertices), len(o.verticesNormals))

	if len(o.vertices) != len(o.verticesNormals) {
		normVertices := geom.NormalizeVertices(o.vertices)
		if normVertices == nil {
			return ErrNormFailed 
		}
		for index, nv := range normVertices {
			o.vertices[index].SetNorm(nv)
		}
	}

	return nil
}