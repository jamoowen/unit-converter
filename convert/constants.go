package convert

import "fmt"

var HelpMessage = `Usage: unit-converter -val <value> -from <unit> -to <unit>
For list of available conversions: $ unit-converter -units
`

var Version = "0.0.1"

var SupportedConversionsMessage = `length: 
    kilometres (km)
    metres (m)
    yards (yds)
    feet (ft)
    inches (in)
    centimetres (cm)

weight:
    grams (g)
    kilograms (kg)
    pounds (lbs)
    stone (st)
    ounces (oz)

temperature:
    Celsius (c)
    Fahrenheit (f)

time:
    milliseconds (ms)
    seconds (sec)
    minutes (min)
    hours (hrs)
    days (dys)
    weeks (wks)
    months (mts)
    years (yrs)
`

var UnsupportedUnitMessage = fmt.Sprintf("Unsupported units...\nSupported:\n%v\n", SupportedConversionsMessage)
