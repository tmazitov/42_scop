package parsing

import (
	"errors"
)

var (
	ErrInvalidVLine = 	errors.New("parsing error : invalid 'v' row appeared")
	ErrInvalidFLine = 	errors.New("parsing error : invalid 'f' row appeared")
	ErrNormFailed =		errors.New("parsing error : vertices normalization failed")
)