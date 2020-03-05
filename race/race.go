package race

import (
	"fmt"
	"runtime"
	"sync"
)

// CheckRaceCondition is great
func CheckRaceCondition() {
	fmt.Println("Numero de CPUs", runtime.NumCPU())
	fmt.Println("Numero de Goroutines", runtime.NumGoroutine())
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	contador := 0

	for i := 0; i < gs; i++ {
		go func() {
			v := contador
			v++
			contador = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("El contador es: ", contador)
}
