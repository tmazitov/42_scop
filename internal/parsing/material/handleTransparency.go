package materialParsing

import (
	"fmt"
	"strconv"

	"github.com/tmazitov/42_scop/internal/rende"
)

// Tr is the inverse of d: Tr = 1 - d (0 = fully opaque, 1 = fully transparent)
func transparencyHandler(proc *mtlParsingProcess, args []string) error {
	if len(args) != 2 {
		return ErrInvalidDissolveLine
	}

	var material *rende.Material
	if material = proc.currentMaterial(); material == nil {
		return ErrInvalidFile
	}

	rawValue, err := strconv.ParseFloat(args[1], 32)
	if err != nil {
		return fmt.Errorf("%w : %w", ErrInvalidDissolveLine, err)
	}

	value := float32(rawValue)
	if value < 0.0 || value > 1.0 {
		return ErrInvalidDissolveLine
	}

	material.SetDissolve(1.0 - value)
	return nil
}
