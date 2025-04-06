package convert

var supportedLengthUnits = [6]string{"km", "m", "yds", "ft", "in", "cm"}

type LengthConverter struct {
	From           string
	To             string
	SupportedUnits []string
}

func (c *LengthConverter) getCalculation(from string) (float64, error) {
	return 0.0, nil
}

func (c *LengthConverter) Convert(from, to string, value float64) (float64, error) {
	return 9999.99, nil

}

func (c *LengthConverter) toCentimetres(from, to string, value float64) (float64, error) {
	if from == "cm" {
		return value, nil
	}
	return value, nil

}
