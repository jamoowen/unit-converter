package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/jamoowen/unit-converter/convert"
)

func main() {
	from := flag.String("from", "", "Base unit")
	to := flag.String("to", "", "Desired unit")

	help := flag.Bool("help", false, "Show help message")
	flag.BoolVar(help, "help", false, "Show help message")

	unitsHelp := flag.Bool("units", false, "Show supported units for conversion")
	flag.BoolVar(unitsHelp, "units", false, "Show supported units for conversion")

	versionHelp := flag.Bool("version", false, "Show version")
	flag.BoolVar(versionHelp, "version", false, "Show version")

	flag.Parse()

	if *help {
		fmt.Fprint(os.Stdout, convert.HelpMessage)
		os.Exit(0)
	}
	if *versionHelp {
		fmt.Fprint(os.Stdout, convert.Version)
		os.Exit(0)
	}
	if *unitsHelp {
		fmt.Fprint(os.Stdout, convert.SupportedConversionsMessage)
		os.Exit(0)
	}

	if len(flag.Args()) != 1 {
		fmt.Fprint(os.Stdout, "Value missing from arguments")
		os.Exit(0)
	}

	converter := convert.NewConverter()
	valueToConvert := flag.Args()[0]
	result, err := converter.ConvertUnits(*from, *to, valueToConvert)
	if err != nil {
		fmt.Fprintf(os.Stdout, err)
		os.Exit(0)
	}
	fmt.Printf("valueL %v, From: %s, To:%v\n", value, *from, *to)
}
