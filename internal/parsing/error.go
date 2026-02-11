package parsing

import (
	"errors"
)

var (
	ErrInvalidVertexLine = 	errors.New("obj parsing error : 'v' row is invalid")
	ErrInvalidFaceLine = 	errors.New("obj parsing error : 'f' row is invalid")
	ErrNormFailed =			errors.New("obj parsing error : vertices normalization failed")
)
