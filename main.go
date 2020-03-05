package main

import (
	"ejercicios_golang/race"
	"fmt"
)

func main() {
	// punteros.Examplepuntero()

	// json.Encodepersona()
	// json.EncodeEjercicio()
	// json.DecodeEjercicio()

	// json.EncodeToStdout()
	// punteros.TestPunteros()

	race.CheckRaceCondition()
}

type persona struct {
	nombre, apellido string
	edad             int
}

type circulo struct {
	radio int
}

type cuadrado struct {
	lado int
}

func (c circulo) area() float64 {
	return (float64(3.1416) * float64(c.radio) * float64(c.radio))
}

func (c cuadrado) area() float64 {
	return (float64(c.lado) * float64(c.lado))
}

func (per persona) presentar() {
	fmt.Println("hola mi nombre es: "+per.nombre+" "+per.apellido+" y mi edad es:", per.edad)
}

type forma interface {
	area() float64
}

func verAreaForma(f forma) {
	fmt.Println(f.area())
}

func exampleAnonima() {
	anonima := func() {
		fmt.Println("Soy una funcion anonima")
	}
	anonima()
}

func exampleCallbacks() {
	fmt.Println(suma(10, 20, 30, 40))
	fmt.Println(suma([]int{10, 20, 30, 40, 50}...))
	sumaConCallback([]int{10, 10, 10, 10, 10}, func(total int) {
		fmt.Println(total)
	})
}

func exampleStructs() {
	persona{"Cristian", "Vasquez", 37}.presentar()

	cuad := cuadrado{2}
	circ := circulo{2}

	verAreaForma(cuad)
	verAreaForma(circ)

	fmt.Println(cuad.area())
	fmt.Println(circ.area())
}

func suma(numeros ...int) (total int) {

	defer func() {
		fmt.Println("Esto corre despues de sumar")
	}()

	for _, i := range numeros {
		total += i
	}
	return
}

func sumaConCallback(numeros []int, callback func(t int)) {
	total := 0
	for _, i := range numeros {
		total += i
	}
	callback(total)
}
