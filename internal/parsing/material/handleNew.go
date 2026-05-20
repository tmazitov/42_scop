package materialParsing

import (
	"strings"

	"github.com/tmazitov/42_scop/internal/rende"
)

// newMaterialHandler handles parsing of a material definition.
func newMaterialHandler(proc *mtlParsingProcess, args []string) error {

	if len(args) < 2 {
		return ErrInvalidMaterialLine
	}

	name := strings.Join(args[1:], " ")

	newMaterial := rende.NewMaterial()
	newMaterial.SetName(name)

	proc.materials = append(proc.materials, newMaterial)

	return nil
}