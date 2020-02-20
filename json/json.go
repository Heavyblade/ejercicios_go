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

type usuario struct {
	Nombre string
	Edad   int
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

// EncodeEjercicio muy importante
func EncodeEjercicio() {
	u1 := usuario{
		Nombre: "Eduar",
		Edad:   32,
	}

	u2 := usuario{
		Nombre: "Juan",
		Edad:   27,
	}

	u3 := usuario{
		Nombre: "Che",
		Edad:   54,
	}

	usuarios := []usuario{u1, u2, u3}

	jotason, err := json.Marshal(usuarios)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jotason))
}
