package objectParsing

import (
	"path/filepath"
	materialParsing "github.com/tmazitov/42_scop/internal/parsing/material"
)

func materialHandler(object *objectParsingProcess, args []string) error {

	if len(args) != 2 {
		return ErrInvalidMaterialLine
	}

	mtlPath := args[1]
	if !isPath(mtlPath) {
		mtlPath = filepath.Join(filepath.Dir(object.filePath), mtlPath)
	}
	
	material, err := materialParsing.ParseMtl(mtlPath)
	if err != nil {
		return err
	}
	object.materials = append(object.materials, material)

	return nil
}