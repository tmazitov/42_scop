package objectParsing

import (
	"strconv"
)

func vertexTextureHandler(object *objectParsingProcess, args []string) error {

	if len(args) < 2 {
		return ErrInvalidVertexTextureLine
	}

	var result = [2]float32{0, 0}
	limit := len(args)
	if limit > 3 {
		limit = 3 // ignore w in "vt u v w"
	}
	for index, part := range args[1:limit] {
		value, err := strconv.ParseFloat(part, 32)
		if err != nil {
			return ErrInvalidVertexTextureLine
		}
		result[index] = float32(value)
	}

	result[1] = 1.0 - result[1] // important to flip

	object.vertexStorage.textures = append(object.vertexStorage.textures, result)

	return nil
}	