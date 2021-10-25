package parsers

import (
	"fmt"

	"github.com/MykytaPopov/gtingen/internal/controllers"
)

type Parser struct{}

func (p Parser) GetCommand(osArgs []string) (string, error) {
	for _, a := range osArgs {
		switch a {
		case controllers.CommandGenerate:
			return a, nil
		case controllers.CommandCalculate:
			return a, nil
		case controllers.CommandValidate:
			return a, nil
		}
	}

	return "", fmt.Errorf("no command provided")
}

func (p Parser) GetArgs(command string, osArgs []string) map[string]string {
	argsParser := getArgsParser(command)

	return argsParser.GetArgs(osArgs)
}

func getArgsParser(command string) IArgsParser {
	switch command {
	case controllers.CommandGenerate:
		return GArgsParser{}
	case controllers.CommandCalculate:
		return CArgsParser{}
	case controllers.CommandValidate:
		return VArgsParser{}
	default:
		panic("unknown command")
	}
}
