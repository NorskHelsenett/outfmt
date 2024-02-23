package main

import (
	"fmt"
	"os"

	"github.com/NorskHelsenett/outfmt"
)

type Information struct {
	Phone  int
	Hidden string
}

type Credentials struct {
	FirstName string
	LastName  string
}

type Person struct {
	Credentials Credentials
	Information Information
}

func init() {
	outfmt.Register(Person{}, &outfmt.Spec{
		"default": {
			{"FIRST NAME", "Credentials.FirstName"},
			{"PHONE", "Information.Phone"},
		},
	})
}

func main() {
	people := []Person{
		Person{
			Credentials: Credentials{
				FirstName: "Linus",
				LastName:  "Johansen",
			},
			Information: Information{
				Phone:  1231241,
				Hidden: ":o",
			},
		},
		Person{
			Credentials: Credentials{
				FirstName: "Vidar",
				LastName:  "Norman",
			},
			Information: Information{
				Phone:  1231,
				Hidden: "jeg er ikke kul",
			},
		},
		Person{
			Credentials: Credentials{
				FirstName: "Haakon",
				LastName:  "Reppen",
			},
			Information: Information{
				Phone:  1231241,
				Hidden: "jeg er kul",
			},
		},
	}

	output, err := outfmt.Format(people, &outfmt.Config{
		Format: outfmt.OutputFormatTable,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not format data: %s", err.Error())
		return
	}

	fmt.Print(string(output))
}
