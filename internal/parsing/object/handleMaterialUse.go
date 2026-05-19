package objectParsing

import "log"

func materialUseHandler(object *objectParsingProcess, args []string) error {

	if len(args) < 2 {
		return ErrInvalidUseMaterialLine
	}

	materialName := args[1]

	material := object.materialStorage.Find(materialName)
	if material == nil {
		log.Printf("obj parsing warn : material '%s' not found, skipping\n", materialName)
		return nil
	}

	material.StartRange(len(object.indices))
	object.materials = append(object.materials, material)

	return nil
}
