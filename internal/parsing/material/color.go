package materialParsing

import (
	"strconv"

	"github.com/tmazitov/42_scop/internal/clr"
)

func parseColor(colorParts []string) (*clr.Color, error) {

	rawColor := make([]float32, 0, len(colorParts))
	for _, arg := range colorParts {
		rawValue, err := strconv.ParseFloat(arg, 32)
		if err != nil {
			return nil, err
		}

		// Clamp to [0, 1] — some exporters write slightly out-of-range values
		if rawValue < 0.0 {
			rawValue = 0.0
		} else if rawValue > 1.0 {
			rawValue = 1.0
		}

		rawColor = append(rawColor, float32(rawValue))
	}

	return clr.NewColorF(rawColor...), nil
}