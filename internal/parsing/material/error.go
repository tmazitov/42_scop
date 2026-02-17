package materialParsing

import (
	"errors"
)

var (
	ErrInvalidMaterialLine = errors.New("mtl parsing error : 'newmtl' row is invalid")
	ErrInvalidShininessLine = errors.New("mtl parsing error : 'Ns' row is invalid") 
	ErrInvalidAmbientColorLine = errors.New("mtl parsing error : 'Ka' row is invalid")
	ErrInvalidDiffuseColorLine = errors.New("mtl parsing error : 'Kd' row is invalid")
	ErrInvalidSpecularColorLine = errors.New("mtl parsing error : 'Ks' row is invalid")
	ErrInvalidDissolveLine = errors.New("mtl parsing error : 'd' row is invalid")
	ErrInvalidOpticalDensityLine = errors.New("mtl parsing error : 'Ni' row is invalid")
	ErrInvalidIlluminationModelLine = errors.New("mtl parsing error : 'illum' row is invalid")
	ErrInvalidDiffuseTextureMapLine = errors.New("mtl parsing error : 'map_Kd' row is invalid")
	ErrNilMaterial = errors.New("mtl parsing error : material is nil")
)