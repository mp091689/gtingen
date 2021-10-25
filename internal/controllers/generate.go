package controllers

import (
	"fmt"
	"strconv"

	"github.com/MykytaPopov/gtingo"
)

type Generate struct{}

func (g Generate) Index(args map[string]string) error {
	gtin := gtingo.Gtin{}

	q, _ := strconv.Atoi(args["-quantity="])
	if q < 1 {
		q = 1
	}

	f, _ := strconv.Atoi(args["-format="])

	for i := q; i > 0; i-- {
		result, err := gtin.Generate(f)
		if err != nil {
			return err
		}

		fmt.Println(result)
	}

	return nil
}
