package convert

import "fmt"

var HelpMessage = `Usage: unit-converter -val <value> -from <unit> -to <unit>
	eg: to convert metres to yards
	$ unit-converter 10 -from m -to yds 
	For list of available conversions:
	$ unit-converter --units

	`

var Version = "0.0.1"

var SupportedConversionsMessage = `
	length: 
		kilometres (km)
		metres (m)
		yards (yds)
		feet (ft)
		inches (in)
		centimetres (cm)	

	`

var UnsupportedUnitMessage = fmt.Sprintf("Unsupported units...\nSupported:\n%v\n", SupportedConversionsMessage)
