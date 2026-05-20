package materialParsing

import (
	"fmt"

	"github.com/tmazitov/42_scop/internal/rende"
)

func emissiveColorHandler(proc *mtlParsingProcess, args []string) error {

	if len(args) != 4 {
		return ErrInvalidEmissiveColorLine
	}

	var material *rende.Material
	if material = proc.currentMaterial(); material == nil {
		return ErrInvalidFile
	}

	color, err := parseColor(args[1:])
	if err != nil {
		return fmt.Errorf("%w : %w", ErrInvalidEmissiveColorLine, err)
	}

	material.SetEmissiveColor(color)
	return nil
}
