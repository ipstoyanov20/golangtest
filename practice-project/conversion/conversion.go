package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(stringSlice []string) ([]float64, error) {
	var floats []float64
	for _, stringVal := range stringSlice {
		floatVal, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("Failed to convert")
		}
		floats = append(floats, floatVal)
	}
	return floats, nil
}