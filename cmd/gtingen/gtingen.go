package gtingen

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/MykytaPopov/gtingo"
)

var allowedFormats = "allowed formats 8(default), 12, 13, 14"
var gHelp = " -g [format] generate random GTIN with passed format\n             " + allowedFormats + "\n\n"
var cHelp = " -c [GTIN*] calculate check sum of provided number\n            *provided number should be without last digit - with outh checksum\n\n"
var vHelp = " -v [GTIN] validate provided GTIN\n"

var help string = "List of flags:\n" + gHelp + cHelp + vHelp

func Run() {
	arg := "-h"

	g := gtingo.NewGtin()

	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	switch arg {
	case "-h":
		fmt.Println(help)
	case "-g":
		format := gtingo.Gtin8

		if len(os.Args) > 2 {
			var err error
			format, err = strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatalln(allowedFormats)
			}
		}

		gtin, err := g.Generate(format)
		if err != nil {
			log.Fatalln(allowedFormats)
		}

		fmt.Println(gtin)
	case "-c":
		if len(os.Args) < 3 {
			log.Fatalln(cHelp)
		}

		partGtin := os.Args[2]
		gtin, err := g.Calculate(partGtin)
		if err != nil {
			log.Fatalln(cHelp)
		}

		fmt.Println(gtin)
	case "-v":
		if len(os.Args) < 3 {
			log.Fatalln(vHelp)
		}

		gtin := os.Args[2]
		ok := g.Validate(gtin)

		v := "valid"
		s := "+"
		if !ok {
			v = "invalid"
			s = "-"
		}

		fmt.Printf("%s Provided GTIN: %s is %s\n", s, gtin, v)
	default:
		fmt.Println("Unknown flag")
	}

}
