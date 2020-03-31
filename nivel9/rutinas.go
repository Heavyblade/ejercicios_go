package nivel9

import (
	"fmt"
	"runtime"
	"sync"
)

// Ejercicio1 es genial
func Ejercicio1() {
	fmt.Println("Hola se van a ejecutar 2")
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("Primera Go rutina")
		wg.Done()
	}()

	go func() {
		fmt.Println("Segunda Go rutina")
		wg.Done()
	}()

	wg.Wait()
}

// Persona es una persona
type Persona struct {
	Name     string
	LastName string
	Age      int
}

type humano interface {
	hablar()
}

func (p *Persona) hablar() {
	fmt.Println("Hola me llamo: ", p.Name)
}

func diAlgo(h humano) {
	h.hablar()
}

// Ejercicio2 Repaso method sets
func Ejercicio2() {
	yo := Persona{
		Name:     "cristian",
		LastName: "Vasquez",
		Age:      37,
	}
	yo.hablar()
	diAlgo(&yo)
}

// Ejercicio3 Race conditions
func Ejercicio3() {

	gs := 100

	var wg sync.WaitGroup
	wg.Add(gs)

	contador := 0

	for i := 0; i < gs; i++ {
		go func() {
			val := contador
			runtime.Gosched()
			val++
			contador = val
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(contador)
}

// Ejercicio3 Arreglando el Race
func Ejercicio3FixMutex() {
	gs := 100

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(gs)

	contador := 0

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			val := contador
			val++
			contador = val
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("El contador bien", contador)
}
