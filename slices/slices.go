package main

import "fmt"

func main(){
	fmt.Println("Slices")
	// un slice se declara de la misma forma que un array, con la diferencia
	// de que se omite el largo de este
	var intSlice []int32 = []int32{4,5,6}
	// alternativa para declarar un slice, tipo, largo , capacidad
	// la capidad puede ser definida opcionalmente, definirla puede ayudar a la performance
	var intSlice2 []int32 = make([]int32, 2, 8)

	

	fmt.Println(intSlice)
	fmt.Printf("El largo del slice es: %v y la capacidad es %v \n" , len(intSlice), cap(intSlice))
	// agrega un elemento al final del slice 
	intSlice = append(intSlice, 8)
	fmt.Println(intSlice)
	fmt.Printf("El largo del slice es: %v y la capacidad es %v \n" , len(intSlice), cap(intSlice))

	// agregando multiples elementos al slice
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	fmt.Printf("El largo del slice es: %v y la capacidad es %v \n" , len(intSlice), cap(intSlice))
}