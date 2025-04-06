package convert

import (
	"fmt"
	"io"
)

type Converter struct {
	CategoryMap map[string]string
}

type UnitConverter interface {
	Convert(from string, to string, value float64) (float64, error)
}

func NewConverter() *Converter {
	var c Converter
	c.CategoryMap = make(map[string]string)
	// set all of the length units
	for _, unit := range supportedLengthUnits {
		c.CategoryMap[unit] = "length"
	}
	return &c
}

func (c *Converter) getConverterForCategory(category string) UnitConverter {
	switch category {
	default:
		return &LengthConverter{}
	}
}

func (c *Converter) ConvertUnits(from, to string, value float64) (float64, error) {
	if from == "" || to == "" {
		fmt.Fprintf(w, "-from & -to units must be provided\nFor help try --help\nFor supported conversions try --supported")
		return
	}
	unitCategory, ok := c.CategoryMap[from]
	if ok == false {
		fmt.Fprint(w, UnsupportedUnitMessage)
		return
	}
	converter := c.getConverterForCategory(unitCategory)
	conversion, err := converter.Convert(from, to, value)
	if err != nil {
		fmt.Fprintf(w, "Failed to convert %v to %v: %v", from, to, err.Error())
		return
	}
	fmt.Fprintf(w, "%.2f %s = %.2f %s\n", value, from, conversion, to)

}
