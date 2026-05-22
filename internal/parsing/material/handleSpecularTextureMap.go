package materialParsing

import (
	"fmt"

	"github.com/tmazitov/42_scop/internal/rende"
)

func specularTextureMapHandler(proc *mtlParsingProcess, args []string) error {

	if len(args) < 2 {
		return ErrInvalidSpecularTextureMapLine
	}

	var material *rende.Material
	if material = proc.currentMaterial(); material == nil {
		return ErrInvalidFile
	}

	texturePath := resolveTexturePath(proc.sourcePath, args[1:])

	textureID, err := loadTexture(texturePath)
	if err != nil {
		return fmt.Errorf("%w : %w", ErrInvalidSpecularTextureMapLine, err)
	}

	material.SetSpecularTextureId(textureID)
	return nil
}
