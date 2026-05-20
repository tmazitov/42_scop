package materialParsing

import (
	"fmt"

	"github.com/tmazitov/42_scop/internal/rende"
)

func ambientTextureMapHandler(proc *mtlParsingProcess, args []string) error {

	if len(args) < 2 {
		return ErrInvalidAmbientTextureMapLine
	}

	var material *rende.Material
	if material = proc.currentMaterial(); material == nil {
		return ErrInvalidFile
	}

	texturePath := resolveTexturePath(proc.sourcePath, args[1:])

	textureID, err := loadTexture(texturePath)
	if err != nil {
		return fmt.Errorf("%w : %w", ErrInvalidAmbientTextureMapLine, err)
	}

	// map_Kd takes priority: only set if no diffuse texture is loaded yet
	if material.TextureId() == 0 {
		material.SetTextureId(textureID)
	}
	return nil
}
