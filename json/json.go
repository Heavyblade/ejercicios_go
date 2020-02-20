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

// DecodeEjercicio is great
func DecodeEjercicio() {
	s := `[{"Nombre":"Eduar","Apellido":"Tua","Edad":32,"Dichos":["Cachicamo diciéndole a morrocoy conchudo","La mona, aunque se vista de seda, mona se queda","Poco a poco se anda lejos"]},{"Nombre":"Carlos","Apellido":"Pérez","Edad":27,"Dichos":["Camarón que se duerme se lo lleva la corriente","A ponerse las alpargatas que lo que viene es joropo","No gastes pólvora en zamuro"]},{"Nombre":"M","Apellido":"Hmmmm","Edad":54,"Dichos":["Ni lava ni presta la batea","Hijo de gato, caza ratón","Más vale pájaro en mano que cien volando"]}]`

	var usuarios []usuario
	err := json.Unmarshal([]byte(s), &usuarios)

	if err != nil {
		fmt.Println(err)
	}
}
