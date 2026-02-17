package objectParsing

import (
	"strconv"
)

func vertexNormalHandler(object *objectParsingProcess, args []string) error {

	if len(args) != 4 {
		return ErrInvalidVertexNormalLine
	}

	var result = [3]float32{0, 0, 0}
	for index, part := range args[1:] {
		value, err := strconv.ParseFloat(part, 32)
		if err != nil {
			return ErrInvalidVertexNormalLine
		}
		result[index] = float32(value)
	}

	object.verticesNormals = append(object.verticesNormals, result)

	return nil
}