package objectParsing

import (
	"errors"
)

var (
	ErrInvalidVertexLine = 			errors.New("obj parsing error : 'v' row is invalid")
	ErrInvalidVertexTextureLine = 	errors.New("obj parsing error : 'vt' row is invalid")
	ErrInvalidVertexNormalLine = 	errors.New("obj parsing error : 'vn' row is invalid")
	ErrInvalidFaceLine = 			errors.New("obj parsing error : 'f' row is invalid")
	ErrInvalidMaterialLine =		errors.New("obj parsing error : 'mtllib' row is invalid")
	ErrNormFailed =					errors.New("obj parsing error : vertices normalization failed")
)
