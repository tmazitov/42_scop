package materialParsing

import (
	"fmt"

	"github.com/tmazitov/42_scop/internal/rende"
)

// Tf r g b — transmission filter: color of light transmitted through the material.
// Stored but not applied in fixed-function OpenGL rendering.
func transmissionFilterHandler(proc *mtlParsingProcess, args []string) error {

	if len(args) != 4 {
		return ErrInvalidTransmissionFilterLine
	}

	var material *rende.Material
	if material = proc.currentMaterial(); material == nil {
		return ErrInvalidFile
	}

	color, err := parseColor(args[1:])
	if err != nil {
		return fmt.Errorf("%w : %w", ErrInvalidTransmissionFilterLine, err)
	}

	material.SetTransmissionFilter(color)
	return nil
}
