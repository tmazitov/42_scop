package materialParsing

import (
	"github.com/tmazitov/42_scop/internal/rende"
)

type mtlParsingFunc func(material *rende.Material, args []string) error

var (
	mtlParsingActionsDictionary = map[mtlLineType]mtlParsingFunc{
		mtlNewMaterial: newMaterialHandler, 
		mtlShininess: shininessHandler,
		mtlAmbientColor: ambientColorHandler,
		mtlDiffuseColor: diffuseColorHandler,
		mtlOpticalDensity: opticalDensityHandler,
		mtlSpecularColor: specularColorHandler,
		mtlDissolve: dissolveHandler,
		mtlIlluminationModel: illuminationModelHandler,
		mtlDiffuseTextureMap: diffuseTextureMapHandler,
	}
)