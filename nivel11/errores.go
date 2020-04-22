package nivel11

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type persona struct {
	Nombre   string
	Apellido string
	Frases   []string
}

type errorPre struct {
	Msg string
}

func (err errorPre) Error() string {
	return err.Msg
}

type raizError struct {
	lat  string
	long string
	err  error
}

func (re raizError) Error() string {
	return fmt.Sprintf("error matemático: %v %v %v", re.lat, re.long, re.err)
}

func ejercicio1() {
	p1 := persona{
		Nombre:   "James",
		Apellido: "Bond",
		Frases:   []string{"Shaken, not stirred", "¿Algún último deseo?", "Nunca digas nunca."},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("Fallo el parseo a json")
	}
	fmt.Println(string(bs))
}

func ejercicio2() {
	p1 := persona{
		Nombre:   "James",
		Apellido: "Bond",
		Frases:   []string{"Shaken, not stirred", "¿Algún último deseo?", "Nunca digas nunca."},
	}

	bs, err := aJSON(p1)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(bs))
}

func ejercicio3() {
	err := errorPre{Msg: "Este es un mensaje de error"}
	foo(err)
}

func foo(e error) {
	fmt.Println(e)
}

// aJSON también necesita retorna un error
func aJSON(a interface{}) (bs []byte, err error) {
	bs, err = json.Marshal(a)

	if err != nil {
		err = fmt.Errorf("Hubo un error al formatear el JSON")
	}
	return
}

func ejercicio4() {
	sqrt, err := raiz(-1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sqrt)
}

func raiz(f float64) (float64, error) {
	if f < 0 {
		err := raizError{lat: "50.2289 N", long: "99.4656 W", err: errors.New("EL numero no puede ser menor que 0")}
		return 0, err
	}
	return 42, nil
}

//Exec ejecuta los ejercicios
func Exec() {
	// ejercicio1()
	// ejercicio2()
	// ejercicio3()
	ejercicio4()
}
