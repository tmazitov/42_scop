package parsing

import (
	"errors"
)

var (
	ErrInvalidVertexLine = 	errors.New("parsing error : invalid 'v' row appeared")
	ErrInvalidFaceLine = 	errors.New("parsing error : invalid 'f' row appeared")
	ErrNormFailed =		errors.New("parsing error : vertices normalization failed")
)