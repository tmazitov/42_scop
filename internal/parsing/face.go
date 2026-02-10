package parsing

import (
	"fmt"
	"strconv"
	"strings"
)

func parseFace(line string) ([]uint32, error) {
	parts := strings.Split(line, " ")
	if len(parts) < 4 {
		return nil, ErrInvalidFaceLine
	}

	// Parse vertex indices (handle v, v/vt, v/vt/vn formats)
	vector := make([]uint32, 0, len(parts)-1)
	for _, part := range parts[1:] {
		// Split by "/" to handle texture and normal coordinates
		// Example: "1/2/3" -> we want just "1"
		vertexData := strings.Split(part, "/")
		
		v, err := strconv.ParseUint(vertexData[0], 10, 32)
		if err != nil {
			return nil, ErrInvalidFaceLine
		}
		
		// OBJ format uses 1-based indexing, convert to 0-based
		if v == 0 {
			return nil, fmt.Errorf("invalid vertex index: 0 (OBJ uses 1-based indexing)")
		}
		vector = append(vector, uint32(v-1))
	}

	// Triangle - no triangulation needed
	if len(vector) == 3 {
		return vector, nil
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

	fmt.Println("line", line, "triangulated to:", triangulated)
	return triangulated, nil
}