package materialParsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
)

func newMaterialHandler(material *rende.Material, args []string) error {

	if len(args) != 2 {
		return ErrInvalidMaterialLine
	}

	if material == nil {
		return ErrNilMaterial
	}

	material.SetName(args[1])

	return nil
}