package convert

import (
	"errors"
	"fmt"
	"slices"
)

var supportedTimeUnits = [8]string{"ms", "sec", "min", "hrs", "dys", "wks", "mts", "yrs"}

type TimeConverter struct {
	SupportedUnits []string
}

func (c *TimeConverter) CheckUnitIsSupported(unit string) bool {
	return slices.Contains(c.SupportedUnits, unit)
}

func (c *TimeConverter) Convert(from, to string, value float64) (float64, error) {
	seconds, err := c.toSeconds(from, value)
	if err != nil {
		return 0.0, errors.New(fmt.Sprintf("Failed to convert val to seconds: %v", err.Error()))
	}
	result, err := c.fromSeconds(to, seconds)
	if err != nil {
		return 0.0, errors.New(fmt.Sprintf("Failed to convert seconds to desired unit: %v", err.Error()))
	}
	return result, nil
}

func (c *TimeConverter) getCalculation(from string) (float64, error) {
	//base case seconds
	switch from {
	case "sec":
		return 1, nil
	case "ms":
		return 0.001, nil
	case "min":
		return 60, nil
	case "hrs":
		return 60 * 60, nil
	case "dys":
		return 24 * 60 * 60, nil
	case "wks":
		return 7 * 24 * 60 * 60, nil
	case "mts":
		return 30 * 24 * 60 * 60, nil
	case "yrs":
		return 365 * 24 * 60 * 60, nil
	default:
		return 0.0, errors.New("No calculation found to convert to seconds...")
	}
}

func (c *TimeConverter) toSeconds(from string, value float64) (float64, error) {
	formula, err := c.getCalculation(from)
	if err != nil {
		return 0.0, err
	}
	return value * formula, nil
}

func (c *TimeConverter) fromSeconds(to string, value float64) (float64, error) {
	formula, err := c.getCalculation(to)
	if err != nil {
		return 0.0, err
	}
	return value / formula, nil
}
