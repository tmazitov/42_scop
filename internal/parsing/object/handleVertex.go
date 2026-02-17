package objectParsing

import (
	"strconv"
	"github.com/tmazitov/42_scop/internal/geom"
)

func vertexHandler(object *objectParsingProcess, args []string) error {
	
	if len(args) != 4 {
		return ErrInvalidVertexLine
	}

	vector := [3]float32{0, 0, 0}

	for index, part := range args[1:] {
		value, err := strconv.ParseFloat(part, 32)
		if err != nil {
			return ErrInvalidVertexLine
		}
		vector[index] = float32(value)
	}

	newVertex := geom.NewVertex(vector)
	object.vertices = append(object.vertices, newVertex)

	return nil
}