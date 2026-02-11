package materialParsing

import (
	"errors"
)

var (
	ErrInvalidMaterialLine = errors.New("mtl parsing error : 'newmtl' row is invalid")
	ErrInvalidShininessLine = errors.New("mtl parsing error : 'Ns' row is invalid") 
	ErrNilMaterial = errors.New("mtl parsing error : material is nil")
)