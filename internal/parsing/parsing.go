package parsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
	"github.com/tmazitov/42_scop/internal/geom"
	"bufio"
	"fmt"
	"os"
)

func ParseObj(filePath string) (*rende.Object, error) {

	var (
		vertexes []*geom.Vertex
		indices []uint32
	)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

    // Create a new scanner to read the file line by line
    scanner := bufio.NewScanner(file)

    // Loop through the file and read each line
	var counter = -1

    for scanner.Scan() {
		counter++
        line := scanner.Text() // Get the line as a string
		lineType := filterLine(line)

		if lineType == "x" {
			continue
		} else if lineType == "v" {
			newVertex, err := parseVertex(line)
			if err != nil {
				return nil, fmt.Errorf("parsing line %d error : %w", counter, err)
			}
			vertexes = append(vertexes, newVertex)
		} else if lineType == "f" {
			newIndices, err := parseFace(line)
			if err != nil {
				return nil, fmt.Errorf("parsing line %d error : %w", counter, err)
			}
			indices = append(indices, newIndices...)
		}
    }

    if err := scanner.Err(); err != nil {
		return nil, err
    }

	normVertices := geom.NormalizeVertices(vertexes)
	if normVertices == nil {
		return nil, ErrNormFailed 
	}

	return rende.NewObject(filePath, normVertices, indices), nil
}