package objectParsing

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tmazitov/42_scop/internal/geom"
)

// Face: format is v/vt/vn
func parseTexturesAndNormals(object *objectParsingProcess, vertex *geom.Vertex, key vertexKey) {

	if textureId := key.Texture(); textureId >= 0 && textureId < len(object.vertexStorage.textures) {
		vertex.SetTextureCoords(object.vertexStorage.textures[textureId])
	}
	if normId := key.Norm(); normId >= 0 && normId < len(object.vertexStorage.normals) {
		vertex.SetNormByVector(object.vertexStorage.normals[normId])
	}
}

func faceHandler(object *objectParsingProcess, args []string) error {

	if len(args) < 4 {
		return ErrInvalidFaceLine
	}


	// Parse vertex indices (handle v, v/vt, v/vt/vn formats)
	vector := make([]uint32, 0, len(args)-1)
	for _, part := range args[1:] {

		// slots: [0]=pos, [1]=texture, [2]=normal; -1 means absent
		vertexRawData := strings.Split(part, "/")
		if len(vertexRawData) == 0 || len(vertexRawData) > 3 {
			return ErrInvalidFaceLine
		}

		storageLens := [3]int{
			len(object.vertexStorage.coords),
			len(object.vertexStorage.textures),
			len(object.vertexStorage.normals),
		}

		var slots [3]int = [3]int{-1, -1, -1}
		for slotIdx, rawElem := range vertexRawData {
			if len(rawElem) == 0 {
				continue // e.g. "1//3" — texture slot stays -1
			}

			convertedElem, err := strconv.ParseInt(rawElem, 10, 32)
			if err != nil {
				return ErrInvalidFaceLine
			}

			if convertedElem == 0 {
				return fmt.Errorf("%w : invalid vertex index: 0 (OBJ uses 1-based indexing)", ErrInvalidFaceLine)
			}

			if convertedElem > 0 {
				slots[slotIdx] = int(convertedElem) - 1
			} else {
				slots[slotIdx] = storageLens[slotIdx] + int(convertedElem)
			}
		}

		if slots[0] < 0 || slots[0] >= len(object.vertexStorage.coords) {
			return ErrInvalidFaceLine
		}
		if slots[2] >= 0 {
			object.hasNormals = true
		}
		key := newVertexKey(slots)

		// When smooth group is off, don't share vertices between faces so
		// ComputeSmoothNormals produces flat (per-face) normals.
		useCache := object.smoothGroup != 0

		if useCache {
			if existingIdx, ok := object.verticesCache[key]; ok {
				vector = append(vector, existingIdx)
				continue
			}
		}

		vertexCoords := object.vertexStorage.coords[key.Pos()]
		newVertex := geom.NewVertex(vertexCoords)
		parseTexturesAndNormals(object, newVertex, key)

		newIdx := uint32(len(object.vertices))
		object.vertices = append(object.vertices, newVertex)
		if useCache {
			object.verticesCache[key] = newIdx
		}
		vector = append(vector, newIdx)
	}

	posFunc := func(idx uint32) [3]float32 {
		p := object.vertices[idx].Pos
		return [3]float32{p.X, p.Y, p.Z}
	}
	triangulated := earClipping(vector, posFunc)

	object.indices = append(object.indices, triangulated...)

	if len(object.materials) > 0 {
		object.materials[len(object.materials)-1].IncreaseRange(len(triangulated))
	}

	return nil
}
