package objectParsing

import (
	"github.com/tmazitov/42_scop/internal/geom"
	"strings"
	"strconv"
	"fmt"
)

// Face: format is v/vt/vn
func parseTexturesAndNormals(object *objectParsingProcess, vertex *geom.Vertex, key vertexKey) {
	
	if textureId := key.Texture(); textureId >= 0 && textureId < len(object.verticesTextures){
		vertex.SetTextureCoords(object.verticesTextures[textureId])
	}
	
	if normId := key.Norm(); normId >= 0 && normId < len(object.verticesNormals){
		vertex.SetNormByVector(object.verticesNormals[normId])
	}
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

			vertexData = append(vertexData, uint32(convertedElem - 1))
		}

		if int(vertexData[0]) >= len(object.verticesCoords) {
			return ErrInvalidFaceLine
		}
		key := newVertexKey(vertexData)

        // Check if this exact combination already exists
		if existingIdx, ok := object.verticesCache[key]; ok {
			vector = append(vector, existingIdx)
		} else {
			vertexCoords := object.verticesCoords[key.Pos()]
			newVertex := geom.NewVertex(vertexCoords) 
			parseTexturesAndNormals(object, newVertex, key)
			
			
			newIdx := uint32(len(object.vertices))
			object.vertices = append(object.vertices, newVertex)
			object.verticesCache[key] = newIdx
			
			vector = append(vector, newIdx)
		}
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


	object.indices = append(object.indices, triangulated...)

	return nil
}