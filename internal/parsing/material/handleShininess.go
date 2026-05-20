package materialParsing

import (
	"strconv"
	"fmt"
	"github.com/tmazitov/42_scop/internal/rende"
)

// shininessHandler handles parsing of shininess.
// Parsing value has a float type and belongs to the range from 0 to 1000. 
func shininessHandler(proc *mtlParsingProcess, args []string) error {

	
	if len(args) < 2 {
		return ErrInvalidShininessLine
	}

	var material *rende.Material
	if material = proc.currentMaterial(); material == nil {
		return ErrInvalidFile
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