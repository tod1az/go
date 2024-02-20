package main

import (
	"fmt"
	"sync"
	"time"
)

// el waitgroup permite controlar que todas las operaciones realizadas en paralelo esten listas antes de finalizar
// la ejecución del programa
var wg = sync.WaitGroup{}
// un mutex asegura que no se pierdan datos a la hora de almacenarlos de forma paralela
var m = sync.RWMutex{}

// simulación de una base de datos
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string {}
func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		// añade un elemento al waitgroup; es un contador con todas las tareas paralelas que hay ejecutandose
		wg.Add(1)
		// la expresion go en frente de la funcion, permite que esta se ejecute en paralelo con el resto del programa
		go dbCall(i)
	}
	// espera a que el waitgroup llegue a 0; es decir que se terminen de ejecutar todas las tareas
	wg.Wait()
	fmt.Printf("\n Total execution time: %v", time.Since(t0))
	fmt.Printf("\n The results are %v", results)
}
//función que simula lo que tardaría una query a una base de datos
func dbCall(i int){
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	save(dbData[i])
	log()
	// la funcion Done, resta uno al contador de tareas activas,
	// esta debe ejecutarse cuando sepamos que una tarea finalizó
	wg.Done()
}

func save(result string){
	// al ejecutar un lock, el mutex revisa si hay otro lock ejecutado y espera a que se desbloquee
	// para ejecutarse
	m.Lock()
	results = append(results, result)
	// libera el bloqueo
	m.Unlock()
}

func log(){
	// este RLock  hace lo mismo que el Lock, pero en vez de prevenirnos de escribir datos, nos previene de leerlos
	// hasta que el proceso actual se termine de ejecutar y leamos data correcta
	m.RLock()
	fmt.Println("the results are:",results)
	m.RUnlock()
}