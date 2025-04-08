package convert

import (
	"errors"
	"fmt"
	"slices"
)

var supportedLengthUnits = [6]string{"km", "m", "yds", "ft", "in", "cm"}

type LengthConverter struct {
	SupportedUnits []string
}

func (c *LengthConverter) CheckUnitIsSupported(unit string) bool {
	return slices.Contains(c.SupportedUnits, unit)
}
func (c *LengthConverter) Convert(from, to string, value float64) (float64, error) {
	// get to cm first
	cmVal, err := c.toCentimetres(from, value)
	if err != nil {
		return 0.0, errors.New(fmt.Sprintf("Failed to convert val to cm: %v", err.Error()))
	}
	result, err := c.fromCentimetres(to, cmVal)
	if err != nil {
		return 0.0, errors.New(fmt.Sprintf("Failed to cm to desired unit: %v", err.Error()))
	}
	return result, nil
}

func (c *LengthConverter) getCalculation(from string) (float64, error) {
	switch from {
	case "cm":
		return 1, nil
	case "in":
		return 2.54, nil
	case "ft":
		return 30.48, nil
	case "yds":
		return 91.44, nil
	case "m":
		return 100, nil
	case "km":
		return 100000, nil
	default:
		return 0.0, errors.New("No calculation found to convert to cm...")
	}
}

func (c *LengthConverter) toCentimetres(from string, value float64) (float64, error) {
	formula, err := c.getCalculation(from)
	if err != nil {
		return 0.0, err
	}
	return value * formula, nil
}

func (c *LengthConverter) fromCentimetres(to string, value float64) (float64, error) {
	formula, err := c.getCalculation(to)
	if err != nil {
		return 0.0, err
	}
	return value / formula, nil
}
