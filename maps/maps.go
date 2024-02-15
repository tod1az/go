package main

import "fmt"

func main(){
	fmt.Println("Maps")

	var myMap  map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"tomas":27, "gonzalo":17}
	fmt.Println(myMap2)
	// acceder a un valor especifico del map
	fmt.Println(myMap2["tomas"])
	// acceder a un valor que no esta registrado regresara un valor por defecto
	// en este caso 0
	fmt.Println(myMap2["asdfqwerty"])
	//los maps devuelven el valor mas un valor booleando que valida la existencia de la clave buscada
	var age, ok = myMap2["tomas"]
	fmt.Println(ok)
	fmt.Println(age)
	var age2, ok2 = myMap2["asdgasdas"]
	fmt.Println(ok2)
	fmt.Println(age2)	

	var value, exist = myMap2["gonzalo"]
	if exist{
		fmt.Println(value)
	}else{
		fmt.Println("Nombre no registrado")
	}

	for key, value := range myMap2{
		fmt.Printf("Nombre:%v Edad:%v \n" , key, value)
	}


	// se elimina por referencia por lo que no devuelve ningun valor
	delete(myMap2, "tomas")
	fmt.Println(myMap2["tomas"])

}