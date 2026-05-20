package materialParsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
	"fmt"
	"strconv"
)

// dissolveHandler parsing dissolve value that has a float type and belongs to the range from 0 to 1.
func opticalDensityHandler(proc *mtlParsingProcess, args []string) error {

	if len(args) < 2 {
		return ErrInvalidOpticalDensityLine
	}

	var material *rende.Material
	if material = proc.currentMaterial(); material == nil {
		return ErrInvalidFile
	}

	rawValue, err := strconv.ParseFloat(args[1], 32)
	if err != nil {
		return fmt.Errorf("%w : %w", ErrInvalidOpticalDensityLine, err)
	}

	value := float32(rawValue)
	if value < 0.001 {
		value = 0.001
	} else if value > 10.0 {
		value = 10.0
	}

	material.SetDensity(value)

	return nil
}