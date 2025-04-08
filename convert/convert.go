package convert

import (
	"errors"
	"fmt"
	"strconv"
)

type Converter struct {
	CategoryMap map[string]string
}

type UnitConverter interface {
	CheckUnitIsSupported(unit string) bool
	Convert(from string, to string, value float64) (float64, error)
}

func NewConverter() *Converter {
	var c Converter
	c.CategoryMap = make(map[string]string)
	// initialize the map of unit to their category

	for _, unit := range supportedLengthUnits {
		c.CategoryMap[unit] = "length"
	}

	for _, unit := range supportedTemperatureUnits {
		c.CategoryMap[unit] = "temperature"
	}

	for _, unit := range supportedTimeUnits {
		c.CategoryMap[unit] = "time"
	}

	for _, unit := range supportedWeightUnits {
		c.CategoryMap[unit] = "weight"
	}

	return &c
}

func (c *Converter) getConverterForCategory(category string) (UnitConverter, error) {
	switch category {
	case "length":
		return &LengthConverter{}, nil
	case "temperature":
		return &TemperatureConverter{}, nil
	case "time":
		return &TimeConverter{}, nil
	case "weight":
		return &WeightConverter{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Unknown category (%v)", category))
	}
}

func (c *Converter) validateInputs(from, to, val string) error {
	msg := ""
	if from == "" {
		msg = msg + " (from) unit must be supplied"
	}
	if to == "" {
		msg = msg + ", (to) unit must be supplied"
	}
	if val == "" {
		msg = msg + ", (val) must be supplied"
	}
	if len(msg) > 0 {
		return errors.New("Input Error:" + msg + "\n")
	}
	return nil
}

func (c *Converter) ConvertUnits(from, to, valueToConvert string) (float64, error) {
	if err := c.validateInputs(from, to, valueToConvert); err != nil {
		return 0.0, err
	}
	value, err := strconv.ParseFloat(valueToConvert, 64)
	if err != nil {
		return 0.0, errors.New(fmt.Sprintf("Invalid number :%v\n", valueToConvert))
	}
	unitCategory, ok := c.CategoryMap[from]
	if ok == false {
		return 0.0, errors.New(fmt.Sprintf("Invalid from value: %v\n", from))
	}
	converter, err := c.getConverterForCategory(unitCategory)
	if err != nil {
		return 0.0, errors.New(fmt.Sprintf("Unable to load converter for category (%v)\n", unitCategory))
	}

	conversion, err := converter.Convert(from, to, value)
	if err != nil {
		return 0.0, err
	}
	return conversion, nil

}
