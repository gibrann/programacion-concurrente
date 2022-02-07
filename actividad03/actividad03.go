// Ejemplo de programa que muestra el uso de mutex para definir las secciones
// criticas a las que el código necesita acceso sincrono.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// contador es una variable incrementado por todas las gorutines.
	contador int

	// wg es usada para esperar a que el rpograma finalice.
	wg sync.WaitGroup

	// mutex es usado para definir el acceso a una seccion critica en el código.
	mutex sync.Mutex
)

// main es el metodo de ejecución principal a todos los programas de go.
func main() {
	// Agrega dos contadores de espera, uno por cada goroutine.
	wg.Add(2)

	// Crea dos goroutines.
	go incContador(1)
	go incContador(2)

	// Espera a que las goroutines terminen.
	wg.Wait()
	fmt.Printf("Contador Final: %d\n", contador)
}

// incContador incrementa el la variable Contador
// con uso de  Mutex para sincronizar y proveer acceso seguro.
func incContador(id int) {
	for count := 0; count < 2; count++ {
		// Solo permite que una goroutine acceda a la
		// seccion critica a la vez.
		mutex.Lock()
		{
			// Captura el valor del contador.
			value := contador
			fmt.Printf("Contador local %d desde id: %d\n", value, id)

			// Cade el hilo y vuelve a colocarse en la cola.
			runtime.Gosched()

			// Incrementa nuestro valor local de contador.
			value++

			// Almaceno el valor en el contador.
			contador = value
		}
		mutex.Unlock()
		// Libera el bloqueo y permite
		// que otra gonrutine en espera
	}

	// Le dice al main que termino.
	wg.Done()
}
