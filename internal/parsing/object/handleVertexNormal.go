package objectParsing

import (
	"math"
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

	length := float32(math.Sqrt(float64(result[0]*result[0] + result[1]*result[1] + result[2]*result[2])))
	if length > 0 {
		result[0] /= length
		result[1] /= length
		result[2] /= length
	}

	object.vertexStorage.normals = append(object.vertexStorage.normals, result)

	return nil
}