package objectParsing

import (
	"path/filepath"
	"log"

	materialParsing "github.com/tmazitov/42_scop/internal/parsing/material"
)

func materialHandler(object *objectParsingProcess, args []string) error {

	if len(args) < 2 {
		return ErrInvalidMaterialLine
	}

	mtlPaths := args[1:]

	for _, mtlPath := range mtlPaths {
		if !isPath(mtlPath) {
			mtlPath = filepath.Join(filepath.Dir(object.filePath), mtlPath)
		}

		materials, err := materialParsing.ParseMtl(mtlPath)
		if err != nil {
			log.Printf("obj parsing warn : failed to parse material file '%s' with error '%s'\n", mtlPath, err)
			continue
		}
		object.materialStorage.Add(materials...)
	}

	return nil
}
