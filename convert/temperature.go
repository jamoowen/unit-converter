package convert

import (
	"errors"
	"slices"
)

var supportedTemperatureUnits = [2]string{"c", "f"}

type TemperatureConverter struct {
	SupportedUnits []string
}

func (c *TemperatureConverter) CheckUnitIsSupported(unit string) bool {
	return slices.Contains(c.SupportedUnits, unit)
}

func (c *TemperatureConverter) Convert(from, to string, value float64) (float64, error) {
	switch from {
	case "c":
		return c.toF(value)
	case "f":
		return c.toC(value)
	default:
		return 0, errors.New("from value not supported")
	}
}

func (c *TemperatureConverter) toC(value float64) (float64, error) {
	return (value - 32) * 5 / 9, nil
}

func (c *TemperatureConverter) toF(value float64) (float64, error) {
	return (value * 9 / 5) + 32, nil
}
