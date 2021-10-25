package controllers

type Controller struct{}

const CommandGenerate = "g"
const CommandCalculate = "c"
const CommandValidate = "v"

func NewController(command string) iController {
	switch command {
	case CommandGenerate:
		return Generate{}
	case CommandCalculate:
		return Calculate{}
	case CommandValidate:
		return Validate{}
	default:
		panic("unkown command")
	}
}
