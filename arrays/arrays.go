package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")
	var intArr [3]int32
	// alternativas para declarar arrays con valores predeterminados
	var arr2 [3]int32 = [3]int32{1,2,3}
	fmt.Println(arr2)
	arr3 := [3]int32{1,2,3}
	fmt.Println(arr3)

	arr4 := [...]int32{1,2,3}
	fmt.Println(arr4)



	intArr[1] = 123
	fmt.Println(intArr[1])
	//imprime el elemento desde la posición 1 a la 3 
	fmt.Println(intArr[1:3])
	fmt.Println(intArr[0:1])

	//imprime la unbicación en memoria del elemento
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])




}