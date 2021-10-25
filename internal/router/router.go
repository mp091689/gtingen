package router

import (
	"github.com/MykytaPopov/gtingen/internal/controllers"
	"github.com/MykytaPopov/gtingen/internal/router/parsers"
)

type Router struct {
	command string
	args    map[string]string
}

func NewRouter(osArgs []string) Router {
	r := Router{}

	p := parsers.Parser{}

	var err error
	r.command, err = p.GetCommand(osArgs)
	if err != nil {
		panic(err)
	}

	r.args = p.GetArgs(r.command, osArgs)

	return r
}

func (r *Router) Process() {
	controller := controllers.NewController(r.command)

	controller.Index(r.args)
}
