package materialParsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
	"fmt"
	"strconv"
)

// dissolveHandler parsing dissolve value that has a float type and belongs to the range from 0 to 1.
func illuminationModelHandler(material *rende.Material, args []string) error {

	if len(args) != 2 {
		return ErrInvalidIlluminationModelLine
	}

	value, err := strconv.Atoi(args[1])
	
	if err != nil {
		return fmt.Errorf("%w : %w", ErrInvalidIlluminationModelLine, err)
	}
	
	if value < 0 || value > 10 {
		return ErrInvalidIlluminationModelLine
	}

	material.SetIlluminationModel(value)

	return nil
}