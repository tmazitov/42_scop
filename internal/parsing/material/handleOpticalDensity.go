package materialParsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
	"fmt"
	"strconv"
)

// dissolveHandler parsing dissolve value that has a float type and belongs to the range from 0 to 1.
func opticalDensityHandler(material *rende.Material, args []string) error {

	if len(args) != 2 {
		return ErrInvalidOpticalDensityLine
	}

	rawValue, err := strconv.ParseFloat(args[1], 32)
	
	if err != nil {
		return fmt.Errorf("%w : %w", ErrInvalidOpticalDensityLine, err)
	}
	
	value := float32(rawValue)
	if value < 0.0 || value > 10.0 {
		return ErrInvalidOpticalDensityLine
	}

	material.SetDensity(value)

	return nil
}