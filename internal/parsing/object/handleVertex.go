package objectParsing

import (
	"strconv"
)

func vertexHandler(object *objectParsingProcess, args []string) error {
	
	if len(args) < 4 {
		return ErrInvalidVertexLine
	}

	var coords [3]float32
	for i, part := range args[1:4] {
		value, err := strconv.ParseFloat(part, 32)
		if err != nil {
			return ErrInvalidVertexLine
		}
		coords[i] = float32(value)
	}
	// args[4] (w) is the homogeneous coordinate; ignored since w=1 in practice.

	object.vertexStorage.coords = append(object.vertexStorage.coords, coords)

	return nil
}