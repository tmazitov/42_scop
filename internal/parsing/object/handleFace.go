package objectParsing

import (
	"strings"
	"strconv"
	"fmt"
	"errors"
)

// Face: format is v/vt/vn
func parseTexturesAndNormals(object *objectParsingProcess, vertexData []uint32) error {
	
	// Create vertex
	vertexIndex := vertexData[0]
	if int(vertexIndex) >= len(object.vertices) {
		return errors.New("vertex index out of range")
	}

	vertex := object.vertices[vertexIndex]
	
	if len(vertexData) >= 2 && int(vertexData[1]) < len(object.verticesTextures){
		vertex.SetTextureCoords(object.verticesTextures[vertexData[1]])
	}
	
	if len(vertexData) == 3 && int(vertexData[2]) < len(object.verticesNormals){
		vertex.SetNormByVector(object.verticesNormals[vertexData[2]])
	}
	
	return nil
}


func faceHandler(object *objectParsingProcess, args []string) error {

	if len(args) < 4 {
		return ErrInvalidFaceLine
	}

	// Parse vertex indices (handle v, v/vt, v/vt/vn formats)
	vector := make([]uint32, 0, len(args)-1)
	for _, part := range args[1:] {
		
		// Split by "/" to handle texture and normal coordinates
		// Example: "1/2/3" -> we want just "1"
		vertexRawData := strings.Split(part, "/")
		if len(vertexRawData) == 0 {
			return ErrInvalidFaceLine
		}

		// 0 - vertex index
		// 1 - vertex texture array index
		// 2 - vertex normal array index
		vertexData := make([]uint32, 0, len(vertexRawData))
		for _, rawElem := range vertexRawData {

			convertedElem, err := strconv.ParseUint(rawElem, 10, 32)
			if err != nil {
				return ErrInvalidFaceLine
			}

			if convertedElem == 0 {
				return fmt.Errorf("%w : invalid vertex index: 0 (OBJ uses 1-based indexing)", ErrInvalidFaceLine)
			}

			vertexData = append(vertexData, uint32(convertedElem))
		}

		vertexData[0]--

		if len(vertexData) > 1 {
			err := parseTexturesAndNormals(object, vertexData)
			if err != nil {
				return fmt.Errorf("%w : during textures and normals parsing error occurred : %w", ErrInvalidFaceLine, err)
			}
		}
		
		vector = append(vector, uint32(vertexData[0]))
	}

	// Triangulate polygon using fan triangulation
	triangleCount := len(vector) - 2
	triangulated := make([]uint32, 0, triangleCount*3)
	
	for i := 1; i < len(vector)-1; i++ {
		triangulated = append(triangulated, 
			vector[0],
			vector[i],
			vector[i+1],
		)
	}

	fmt.Printf("vector : '%+v' triangulated to '%+v'\n", vector, triangulated)

	object.indices = append(object.indices, triangulated...)

	return nil
}