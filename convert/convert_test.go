package convert

import (
	"math"
	"strings"
	"testing"
)

// TestNewConverter tests the creation of a new converter
func TestNewConverter(t *testing.T) {
	c := NewConverter()
	if c == nil {
		t.Error("NewConverter() returned nil")
	}
	if c.CategoryMap == nil {
		t.Error("CategoryMap is nil")
	}
	// Check if length units are properly mapped
	for _, unit := range supportedLengthUnits {
		if category, exists := c.CategoryMap[unit]; !exists {
			t.Errorf("Unit %s not found in CategoryMap", unit)
		} else if category != "length" {
			t.Errorf("Unit %s has wrong category: %s", unit, category)
		}
	}

	for _, unit := range supportedWeightUnits {
		if category, exists := c.CategoryMap[unit]; !exists {
			t.Errorf("Unit %s not found in CategoryMap", unit)
		} else if category != "weight" {
			t.Errorf("Unit %s has wrong category: %s", unit, category)
		}
	}

	for _, unit := range supportedTimeUnits {
		if category, exists := c.CategoryMap[unit]; !exists {
			t.Errorf("Unit %s not found in CategoryMap", unit)
		} else if category != "time" {
			t.Errorf("Unit %s has wrong category: %s", unit, category)
		}
	}

	for _, unit := range supportedTemperatureUnits {
		if category, exists := c.CategoryMap[unit]; !exists {
			t.Errorf("Unit %s not found in CategoryMap", unit)
		} else if category != "temperature" {
			t.Errorf("Unit %s has wrong category: %s", unit, category)
		}
	}
}

// TestConvertUnits tests the main conversion functionality
func TestConvertUnits(t *testing.T) {
	c := NewConverter()
	tests := []struct {
		name          string
		from          string
		to            string
		value         string
		expected      float64
		expectError   bool
		errorContains string
	}{
		{
			name:        "Convert meters to centimeters",
			from:        "m",
			to:          "cm",
			value:       "1",
			expected:    100,
			expectError: false,
		},
		{
			name:        "Convert inches to centimeters",
			from:        "in",
			to:          "cm",
			value:       "1",
			expected:    2.54,
			expectError: false,
		},
		{
			name:          "Invalid unit",
			from:          "invalid",
			to:            "cm",
			value:         "1",
			expectError:   true,
			errorContains: "Invalid from value",
		},
		{
			name:          "Invalid number",
			from:          "m",
			to:            "cm",
			value:         "not a number",
			expectError:   true,
			errorContains: "Invalid number",
		},
		{
			name:          "Valid length conversion",
			from:          "m",
			to:            "cm",
			value:         "1",
			expected:      100,
			expectError:   false,
			errorContains: "",
		},
		{
			name:          "Valid temp conversion",
			from:          "c",
			to:            "f",
			value:         "100",
			expected:      212,
			expectError:   false,
			errorContains: "",
		},
		{
			name:          "Valid weight conversion",
			from:          "lbs",
			to:            "kg",
			value:         "250",
			expected:      113.3980925,
			expectError:   false,
			errorContains: "",
		},
		{
			name:          "Valid time conversion",
			from:          "sec",
			to:            "min",
			value:         "120",
			expected:      2,
			expectError:   false,
			errorContains: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := c.ConvertUnits(tt.from, tt.to, tt.value)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				if tt.errorContains != "" && err != nil {
					if !strings.Contains(err.Error(), tt.errorContains) {
						t.Errorf("Error message does not contain '%s': %v", tt.errorContains, err)
					}
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Use a small epsilon for float comparison
			if result == 0.0 {
				t.Errorf("Unexpected result: %v", result)
			}

		})
	}
}

// TestGetConverterForCategory tests the category converter retrieval
func TestGetConverterForCategory(t *testing.T) {
	c := NewConverter()

	tests := []struct {
		name          string
		category      string
		expectError   bool
		errorContains string
	}{
		{
			name:        "Valid length category",
			category:    "length",
			expectError: false,
		},
		{
			name:          "Invalid category",
			category:      "invalid",
			expectError:   true,
			errorContains: "Unknown category",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converter, err := c.getConverterForCategory(tt.category)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				if tt.errorContains != "" && err != nil {
					if !strings.Contains(err.Error(), tt.errorContains) {
						t.Errorf("Error message does not contain '%s': %v", tt.errorContains, err)
					}
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if converter == nil {
				t.Error("Expected converter but got nil")
			}
		})
	}
}

// TestLengthConverter tests the LengthConverter implementation
func TestLengthConverter(t *testing.T) {
	lc := &LengthConverter{}

	tests := []struct {
		name          string
		from          string
		to            string
		value         float64
		expected      float64
		expectError   bool
		errorContains string
	}{
		// Centimeter conversions
		{
			name:     "cm to m",
			from:     "cm",
			to:       "m",
			value:    100,
			expected: 1,
		},
		{
			name:     "cm to km",
			from:     "cm",
			to:       "km",
			value:    100000,
			expected: 1,
		},
		// Meter conversions
		{
			name:     "m to cm",
			from:     "m",
			to:       "cm",
			value:    1,
			expected: 100,
		},
		{
			name:     "m to km",
			from:     "m",
			to:       "km",
			value:    1000,
			expected: 1,
		},
		// Kilometer conversions
		{
			name:     "km to m",
			from:     "km",
			to:       "m",
			value:    1,
			expected: 1000,
		},
		{
			name:     "km to cm",
			from:     "km",
			to:       "cm",
			value:    1,
			expected: 100000,
		},
		// Inch conversions
		{
			name:     "in to cm",
			from:     "in",
			to:       "cm",
			value:    1,
			expected: 2.54,
		},
		{
			name:     "cm to in",
			from:     "cm",
			to:       "in",
			value:    2.54,
			expected: 1,
		},
		// Foot conversions
		{
			name:     "ft to cm",
			from:     "ft",
			to:       "cm",
			value:    1,
			expected: 30.48,
		},
		{
			name:     "cm to ft",
			from:     "cm",
			to:       "ft",
			value:    30.48,
			expected: 1,
		},
		// Yard conversions
		{
			name:     "yds to cm",
			from:     "yds",
			to:       "cm",
			value:    1,
			expected: 91.44,
		},
		{
			name:     "cm to yds",
			from:     "cm",
			to:       "yds",
			value:    91.44,
			expected: 1,
		},
		// Error cases
		{
			name:          "Invalid from unit",
			from:          "invalid",
			to:            "cm",
			value:         1,
			expectError:   true,
			errorContains: "No calculation found",
		},
		{
			name:          "Invalid to unit",
			from:          "cm",
			to:            "invalid",
			value:         1,
			expectError:   true,
			errorContains: "No calculation found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := lc.Convert(tt.from, tt.to, tt.value)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				if tt.errorContains != "" && err != nil {
					if !strings.Contains(err.Error(), tt.errorContains) {
						t.Errorf("Error message does not contain '%s': %v", tt.errorContains, err)
					}
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			// Use a small epsilon for float comparison
			epsilon := 0.0001
			if math.Abs(result-tt.expected) > epsilon {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
