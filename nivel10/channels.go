package nivel10

import (
	"fmt"
	"runtime"
)

// Ejercicio1 problema de un channel que no funciona
func ejercicio1() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

// Ejercicio1B mismo problema pero channel con buffer
func ejercicio1B() {
	c := make(chan int, 2)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func ejercicio2() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func ejercicio3() {
	c := gen()
	recibir(c)

	fmt.Println("A punto de finalizar.")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func recibir(c <-chan int) {
	for val := range c {
		fmt.Println(val)
	}
}

func ejercicio4() {
	salir := make(chan int)
	c := gen4(salir)

	recibir4(c, salir)

	fmt.Println("A punto de finalizar(4).")
}

func gen4(salir <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func recibir4(c <-chan int, salir <-chan int) {
	for {
		select {
		case x := <-c:
			fmt.Println(x)
		case <-salir:
			fmt.Println("salir")
			return
		}
	}
}

func ejercicio5() {
	c := make(chan int)

	go func() {
		c <- 50
		c <- 0
		close(c)
	}()

	primero, ok := <-c
	fmt.Println(primero, ok)

	segndo, ok2 := <-c
	fmt.Println(segndo, ok2)

	tercero, ok3 := <-c
	fmt.Println(tercero, ok3)
}

func ejercicio6() {
	ch := make(chan int)

	go agregarTrabajo(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func agregarTrabajo(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func ejercicio7() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go agregarTrabajo2(ch)
	}

	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("GoRUTINAS", runtime.NumGoroutine())
}

func agregarTrabajo2(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
}

// Exec executes the exercises in the current package
func Exec() {
	// ejercicio1()
	// ejercicio1B()
	// ejercicio2()
	// ejercicio3()
	// ejercicio4()
	// ejercicio5()
	// ejercicio6()
	ejercicio7()
}
