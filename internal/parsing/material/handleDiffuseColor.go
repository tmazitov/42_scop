package materialParsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
	"fmt"
)

// ambientColorHandler parsing ambient color following RGB standard format,
// where each of values has a float type and belongs to the range from 0 to 1.
func diffuseColorHandler(material *rende.Material, args []string) error {

	if len(args) != 4 {
		return ErrInvalidDiffuseColorLine
	}

	color, err := parseColor(args[1:])
	if err != nil {
		return fmt.Errorf("%w : %w", ErrInvalidDiffuseColorLine, err)
	}

	material.SetDiffuseColor(color)

	return nil
}