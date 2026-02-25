package objectParsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
)

func materialUseHandler(object *objectParsingProcess, args []string) error {

	if len(args) != 2 {
		return ErrInvalidUseMaterialLine
	}

	var (
		materialName string = args[1]
		material	 *rende.Material
	)


	for _, tempMaterial := range object.materials {
		if materialName == tempMaterial.Name() {
			material = tempMaterial
		}
	}

	if material == nil {
		return ErrInvalidUseMaterialLine
	}

	material.StartRange(len(object.indices))
	object.currentMaterial = material

	return nil
}