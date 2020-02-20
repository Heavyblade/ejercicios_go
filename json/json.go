package json

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

// Encodepersona is pretty
func Encodepersona() {
	p1 := persona{
		Nombre:   "Cristian",
		Apellido: "vasquez",
		Edad:     37,
	}
	p2 := persona{
		Nombre:   "Milena",
		Apellido: "Osorno",
		Edad:     36,
	}

	personas := []persona{p1, p2}
	fmt.Println(personas)

	bs, err := json.Marshal(personas)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
