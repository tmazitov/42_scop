package materialParsing

import (
	"strconv"
	"errors"
	"github.com/tmazitov/42_scop/internal/clr"
)

func parseColor(colorParts []string) (*clr.Color, error) {

	rawColor := make([]float32, 0, len(colorParts))
	for _, arg := range colorParts {
		rawValue, err := strconv.ParseFloat(arg, 32)
		
		if err != nil {
			return nil, err
		}
		
		if rawValue < 0.0 || rawValue > 1.0 {
			return nil, errors.New("color value doesn't belong to the range from 0 to 1")
		}

		rawColor = append(rawColor, float32(rawValue))
	}

	return clr.NewColorF(rawColor...), nil
}