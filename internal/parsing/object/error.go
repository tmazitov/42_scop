package objectParsing

import (
	"errors"
)

var (
	ErrInvalidVertexLine        = errors.New("obj parsing error : 'v' row is invalid")
	ErrInvalidVertexTextureLine = errors.New("obj parsing error : 'vt' row is invalid")
	ErrInvalidVertexNormalLine  = errors.New("obj parsing error : 'vn' row is invalid")
	ErrInvalidFaceLine          = errors.New("obj parsing error : 'f' row is invalid")
	ErrInvalidInitObjectLine    = errors.New("obj parsing error : 'o' row is invalid")
	ErrInvalidMaterialLine      = errors.New("obj parsing error : 'mtllib' row is invalid")
	ErrInvalidUseMaterialLine   = errors.New("obj parsing error : 'usemtl' row is invalid")
	ErrObjectNotDeclared        = errors.New("obj parsing error : model object is not declared")
)
