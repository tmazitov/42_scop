package parsing

import (
	"strings"
	"strconv"
	"fmt"
)

func parseFace(line string) ([]uint32, error) {
	parts := strings.Split(line, " ")
	if len(parts) < 4 {
		return nil, ErrInvalidFLine
	}

	// fmt.Println(parts)

	vector := make([]uint32, 0, len(parts) - 1)

	for _, part := range parts[1:] {
		v, err := strconv.ParseUint(part, 10, 32)
		if err != nil {
			return nil, ErrInvalidFLine
		}

		vector = append(vector, uint32(v))
	}
	
	if len(vector) == 3 {
		return vector, nil
	}

	// Triangulated version

	triangleCount := (len(vector) - 3) + 1
	triangulated := make([]uint32, 0, triangleCount * 3)

	for index := range vector {
		if index == 0 || index == 1 {
			continue
		}

		triangulated = append(triangulated, []uint32{
			vector[0],
			vector[index - 1],
			vector[index],
		}...)
	}
	
	fmt.Println("triangulated to :", triangulated)

	return triangulated, nil
}
