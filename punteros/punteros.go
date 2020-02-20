package punteros

import "fmt"

type persona struct {
	nombre, apellido string
}

func (p *persona) cambiarNombre(nuevo string) {
	(*p).nombre = nuevo
}

// Examplepuntero is very importatn
func Examplepuntero() {
	hola := 90
	fmt.Println(hola)
	fmt.Println(&hola)

	cambiarValor(&hola)

	fmt.Println(hola)
	fmt.Println(&hola)

	cristian := persona{"Cristian", "Gaviria"}

	cristian.cambiarNombre("Camilo")
	fmt.Println(cristian)
}

func cambiarValor(val *int) {
	*val = 50
}
