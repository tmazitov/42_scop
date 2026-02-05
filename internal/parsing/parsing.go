package parsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
	"bufio"
	"fmt"
	"os"
)

func ParseObj(filePath string) (*rende.Object, error) {

	var (
		vertexes []*rende.Vertex
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
				return nil, fmt.Errorf("parsing line %d erorr : %w", counter, err)
			}
			vertexes = append(vertexes, newVertex)
		} else if lineType == "f" {
			newIndices, err := parseFace(line)
			if err != nil {
				return nil, fmt.Errorf("parsing line %d erorr : %w", counter, err)
			}
			indices = append(indices, newIndices...)
		}

    }

    // Check for errors during the scan
    if err := scanner.Err(); err != nil {
		return nil, err
    }

	fmt.Println("Before :")
	for _, v := range vertexes {
		fmt.Println(v.ToString())
	}

	normVerticises := NormalizeVertices(vertexes)
	if normVerticises == nil {
		return nil, ErrNormFailed 
	}

	fmt.Println("After :")
	for _, v := range vertexes {
		fmt.Println(v.ToString())
	}

	return rende.NewObject(filePath, normVerticises, indices), nil
}