package materialParsing

import (
	"strconv"
	"fmt"
	"github.com/tmazitov/42_scop/internal/rende"
)

func shininessHandler(material *rende.Material, args []string) error {

	
	if len(args) != 2 {
		return ErrInvalidMaterialLine
	}

	rawValue, err := strconv.ParseFloat(args[1], 32)
	
	if err != nil {
		return fmt.Errorf("%w : %w", ErrInvalidShininessLine, err)
	}
	
	value := float32(rawValue)
	if value < 0.0 || value > 1000.0 {
		return ErrInvalidShininessLine
	}

	material.SetShininess(value)

	return nil
}