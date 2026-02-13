package materialParsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
	"fmt"
)

// specularColorHandler parsing specular color following RGB standard format,
// where each of values has a float type and belongs to the range from 0 to 1.
func specularColorHandler(material *rende.Material, args []string) error {

	if len(args) != 4 {
		return ErrInvalidSpecularColorLine
	}

	color, err := parseColor(args[1:])
	if err != nil {
		return fmt.Errorf("%w : %w", ErrInvalidSpecularColorLine, err)
	}

	material.SetSpecularColor(color)

	return nil
}