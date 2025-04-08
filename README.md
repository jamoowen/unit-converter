# Unit Converter CLI

A command-line tool for converting between different units of measurement.

## Features

- Convert between common units of:
  - Length (kilometers, meters, yards, feet, inches, centimeters)
  - Weight (grams, kilograms, pounds, stone, ounces)
  - Temperature (Celsius, Fahrenheit)
  - Time (milliseconds, seconds, minutes, hours, days, weeks, months, years)

## Usage

```bash
# Convert 5 meters to feet
unit-converter 5 -from m -to ft

# Convert 100 Celsius to Fahrenheit
unit-converter 100 -from c -to f

# Convert 250 pounds to kilograms
unit-converter 250 -from lbs -to kg

# Convert 2 hours to minutes
unit-converter 2 -from hrs -to min
```

## Supported Units

### Length
- km (kilometers)
- m (meters)
- yds (yards)
- ft (feet)
- in (inches)
- cm (centimeters)

### Weight
- g (grams)
- kg (kilograms)
- lbs (pounds)
- st (stone)
- oz (ounces)

### Temperature
- c (Celsius)
- f (Fahrenheit)

### Time
- ms (milliseconds)
- sec (seconds)
- min (minutes)
- hrs (hours)
- dys (days)
- wks (weeks)
- mts (months)
- yrs (years)
