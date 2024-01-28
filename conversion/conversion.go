package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {

	var floats []float64

	for _, strVal := range strings {
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return nil, errors.New("Error converting to float")
		}

		floats = append(floats, floatVal)
	}
	return floats, nil
}
