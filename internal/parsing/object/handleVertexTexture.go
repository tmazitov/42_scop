package objectParsing

import (
	"strconv"
)

func vertexTextureHandler(object *objectParsingProcess, args []string) error {

	if len(args) != 3 {
		return ErrInvalidVertexTextureLine
	}

	var result = [2]float32{0, 0}
	for index, part := range args[1:] {
		value, err := strconv.ParseFloat(part, 32)
		if err != nil {
			return ErrInvalidVertexTextureLine
		}
		result[index] = float32(value)
	}

	result[1] = 1.0 - result[1] // important to flip

	object.verticesTextures = append(object.verticesTextures, result)

	return nil
}	