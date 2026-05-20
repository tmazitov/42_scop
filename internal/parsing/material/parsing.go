package materialParsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
)

type mtlParsingFunc func(proc *mtlParsingProcess, args []string) error

type mtlParsingProcess struct {
	materials []*rende.Material
	sourcePath string
}

func newMtlParsingProcess(sourcePath string) *mtlParsingProcess {
	return &mtlParsingProcess{
		materials: nil,
		sourcePath: sourcePath,
	}
}

func (m *mtlParsingProcess) currentMaterial() *rende.Material {
	if len(m.materials) == 0 {
		return nil
	}
	return m.materials[len(m.materials) - 1]
}

var (
	mtlParsingActionsDictionary = map[mtlLineType]mtlParsingFunc{
		mtlNewMaterial:        newMaterialHandler,
		mtlShininess:          shininessHandler,
		mtlAmbientColor:       ambientColorHandler,
		mtlDiffuseColor:       diffuseColorHandler,
		mtlSpecularColor:      specularColorHandler,
		mtlEmissiveColor:      emissiveColorHandler,
		mtlOpticalDensity:     opticalDensityHandler,
		mtlDissolve:           dissolveHandler,
		mtlTransparency:       transparencyHandler,
		mtlTransmissionFilter: transmissionFilterHandler,
		mtlIlluminationModel:  illuminationModelHandler,
		mtlDiffuseTextureMap:  diffuseTextureMapHandler,
		mtlAmbientTextureMap:  ambientTextureMapHandler,
	}
)