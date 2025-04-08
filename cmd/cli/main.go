package main

import (
	"flag"
	"fmt"
	"github.com/jamoowen/unit-converter/convert"
	"os"
)

func main() {
	var from, to, val string
	var help bool

	flag.StringVar(&from, "from", "", "The base unit")
	flag.StringVar(&to, "to", "", "The desired unit")
	flag.StringVar(&val, "val", "", "The value to be converted")
	flag.BoolVar(&help, "help", false, "Show help message")

	flag.Parse()

	if help {
		fmt.Fprint(os.Stdout, convert.HelpMessage)
		return
	}

	converter := convert.NewConverter()
	result, err := converter.ConvertUnits(from, to, val)
	if err != nil {
		fmt.Fprint(os.Stdout, err.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "%v %s\n", result, to)
}
