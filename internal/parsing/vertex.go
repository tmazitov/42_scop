package parsing

import (
	"strings"
	"strconv"
	"github.com/tmazitov/42_scop/internal/rende"
)

func parseVertex(line string) (*rende.Vertex, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 4 {
		return nil, ErrInvalidVertexLine
	}

	vector := [3]float32{0, 0, 0}

	for index, part := range parts[1:] {
		value, err := strconv.ParseFloat(part, 32)
		if err != nil {
			return nil, ErrInvalidVertexLine
		}
		vector[index] = float32(value)
	}

	return rende.NewVertex(vector), nil
}