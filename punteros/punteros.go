package punteros

import "fmt"

type persona struct {
	nombre, apellido string
}

func (p *persona) cambiarNombre(nuevo string) {
	(*p).nombre = nuevo
}

func (p persona) sinPuntero() {
	fmt.Println("llama a sin puntero")
	(&p).nombre = "Milena"
}

func (p *persona) conPuntero() {
	fmt.Println("llama a con puntero")
	(*p).nombre = "Santiago"
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

// TestPunteros es muy chevere
func TestPunteros() {
	per := persona{"Juan", "Perez"}
	punt := &per

	fmt.Println("xxxxxxxxxxx objeto xxxxxxx")
	per.sinPuntero()
	fmt.Println(per.nombre)
	per.conPuntero()
	fmt.Println(per.nombre)

	fmt.Println("xxxxxxxxxxx puntero xxxxxxx")
	punt.sinPuntero()
	fmt.Println(per.nombre)
	punt.conPuntero()
	fmt.Println(per.nombre)
}
