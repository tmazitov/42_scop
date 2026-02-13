package materialParsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
	"fmt"
	"strconv"
)

// dissolveHandler parsing dissolve value that has a float type and belongs to the range from 0 to 1.
func dissolveHandler(material *rende.Material, args []string) error {

	if len(args) != 2 {
		return ErrInvalidDissolveLine
	}

	rawValue, err := strconv.ParseFloat(args[1], 32)
	
	if err != nil {
		return fmt.Errorf("%w : %w", ErrInvalidDissolveLine, err)
	}
	
	value := float32(rawValue)
	if value < 0.0 || value > 1.0 {
		return ErrInvalidDissolveLine
	}

	material.SetDissolve(value)

	return nil
}