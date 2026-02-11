package parsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
	"github.com/tmazitov/42_scop/internal/geom"
	materialParsing "github.com/tmazitov/42_scop/internal/parsing/material"
	"bufio"
	"fmt"
	"strings"
	"path/filepath"
	"os"
)


func ParseObj(filePath string) (*rende.Object, error) {

	var (
		vertices []*geom.Vertex
		indices []uint32
		object *rende.Object
		materialPaths []string
		materials []*rende.Material
	)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	object = rende.NewObject(filepath.Base(filePath))

    // Create a new scanner to read the file line by line
    scanner := bufio.NewScanner(file)

    // Loop through the file and read each line
	var counter = -1
	var lineArgs []string
    for scanner.Scan() {
		counter++
        line := scanner.Text() // Get the line as a string
		
		if lineArgs = filterObjFileLine(line); len(lineArgs) == 0{
			continue
		}
		
		if lineArgs[0] == "v" {
			newVertex, err := parseVertex(line)
			if err != nil {
				return nil, fmt.Errorf("parsing line %d error : %w", counter, err)
			}
			vertices = append(vertices, newVertex)
		} else if lineArgs[0] == "f" {
			newIndices, err := parseFace(line)
			if err != nil {
				return nil, fmt.Errorf("parsing line %d error : %w", counter, err)
			}
			indices = append(indices, newIndices...)
		} else if lineArgs[0] == "mtllib" {

			if len(lineArgs) != 2 {
				return nil, fmt.Errorf("parsing line %d error : invalid 'newmtl' line", counter)
			}

			materialPaths = append(materialPaths, lineArgs[1])
		}
    }

    if err := scanner.Err(); err != nil {
		return nil, err
    }

	normVertices := geom.NormalizeVertices(vertices)
	if normVertices == nil {
		return nil, ErrNormFailed 
	}

	materials = make([]*rende.Material, 0, len(materialPaths))
	for _, rawMtlPath := range materialPaths {

		path := rawMtlPath

		if !isPath(path) {
			path = filepath.Join(filepath.Dir(filePath), rawMtlPath)
		}

		material, err := materialParsing.ParseMtl(path)
		if err != nil {
			return nil, err
		}
		materials = append(materials, material)
	}

	object.
		SetShape(normVertices).
		SetIndices(indices).
		SetMaterials(materials)

	return object, nil
}

func isPath(s string) bool {
    // Check for common path separators
    return strings.Contains(s, "/") || 
           strings.Contains(s, "\\") ||
           strings.Contains(s, string(filepath.Separator))
}