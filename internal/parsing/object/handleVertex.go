package objectParsing

import (
	"strconv"
)

func vertexHandler(object *objectParsingProcess, args []string) error {
	
	if len(args) != 4 {
		return ErrInvalidVertexLine
	}

	coords := [3]float32{0, 0, 0}

	for index, part := range args[1:] {
		value, err := strconv.ParseFloat(part, 32)
		if err != nil {
			return ErrInvalidVertexLine
		}
		coords[index] = float32(value)
	}

	
	object.verticesCoords = append(object.verticesCoords, coords)

	return nil
}