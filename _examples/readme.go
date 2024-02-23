package main

import (
	"fmt"
	"os"

	"github.com/NorskHelsenett/outfmt"
)

type Data struct {
	Id     string
	Name   string
	Active bool
}

func init() {
	outfmt.Register(Data{}, &outfmt.Spec{
		"default": {
			{"ID", "Id"},
			{"NAME", "Name"},
		},
		"wide": {
			{"ID", "Id"},
			{"NAME", "Name"},
			{"ACTIVE", "Active"},
		},
	})
}

func main() {
	data := []Data{
		{
			Id:     "12323",
			Name:   "cool",
			Active: true,
		},
		{
			Id:     "534324",
			Name:   "cool",
			Active: false,
		},
		{
			Id:     "1gerfs",
			Name:   "cool",
			Active: true,
		},
	}

	output, err := outfmt.Format(data, &outfmt.Config{
		// default
		Format: outfmt.OutputFormatTable,

		// wide
		//Format:          outfmt.OutputFormatCondition,
		//AdditionalField: "wide",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not format data: %s", err.Error())
		return
	}

	fmt.Print(string(output))
}
