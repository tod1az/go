package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var p *int32 = new(int32)
	var i int32 = 99
    //para acceder al valor del puntero se antepone un "*"
	fmt.Printf("pointer: %v regular variable: %v\n",*p,i)
    //para modificar el valor del puntero se antepone un "*"
	*p = 10
	fmt.Printf("pointer: %v regular variable: %v\n",*p,i)
	//para cambiar el lugar en memoria al que apunta el puntero se utiliza el simbolo "&"
	fmt.Printf("p is pointing => %v\n", p)
	fmt.Printf("i is pointing => %v\n", &i)
	fmt.Println("------------------------------")
	p = &i
	// en este caso el p apunta al lugar en memoria que se almaceno el valor de i 
	fmt.Printf("p is pointing => %v\n", p)
	fmt.Printf("i is pointing => %v\n", &i)
	// por lo tanto modificar el valor de p, modificaria tambien el valor de i, ya que p esta apuntando al mismo lugar 
	// en memoria que i 
	fmt.Printf("value of p: %v\n", *p)
	fmt.Printf("value of i: %v\n", i)
	*p =5
	fmt.Println("------------------------------")
	fmt.Printf("value of p: %v\n", *p)
	fmt.Printf("value of i: %v\n", i)
	//* nota, los slices utilizan punteros implicitos, por lo que modificar una copia de un slice 
	// tambien modificara al slice original
}