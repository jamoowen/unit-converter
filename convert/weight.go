package convert

import (
	"errors"
	"fmt"
	"slices"
)

var supportedWeightUnits = [5]string{"g", "kg", "lbs", "st", "oz"}

type WeightConverter struct {
	SupportedUnits []string
}

func (c *WeightConverter) CheckUnitIsSupported(unit string) bool {
	return slices.Contains(c.SupportedUnits, unit)
}

func (c *WeightConverter) Convert(from, to string, value float64) (float64, error) {
	gramVal, err := c.toGrams(from, value)
	if err != nil {
		return 0.0, errors.New(fmt.Sprintf("Failed to convert val to grams: %v", err.Error()))
	}
	result, err := c.fromGrams(to, gramVal)
	if err != nil {
		return 0.0, errors.New(fmt.Sprintf("Failed to convert grams to desired unit: %v", err.Error()))
	}
	return result, nil
}

func (c *WeightConverter) getCalculation(from string) (float64, error) {
	switch from {
	case "g":
		return 1, nil
	case "kg":
		return 1000, nil
	case "lbs":
		return 453.59237, nil
	case "st":
		return 6350.29318, nil
	case "oz":
		return 28.349523125, nil
	default:
		return 0.0, errors.New("No calculation found to convert to grams...")
	}
}

func (c *WeightConverter) toGrams(from string, value float64) (float64, error) {
	formula, err := c.getCalculation(from)
	if err != nil {
		return 0.0, err
	}
	return value * formula, nil
}

func (c *WeightConverter) fromGrams(to string, value float64) (float64, error) {
	formula, err := c.getCalculation(to)
	if err != nil {
		return 0.0, err
	}
	return value / formula, nil
}
