package main

import (
	"errors"
	"fmt"
)

func main() {
	var result string =	retunrSomething("funcion que retorna algo")
	fmt.Println(result)
	printMe("funcion void")
	// se puede inicializar más de una variable a la vez
	var returnedValue1, returnedValue2 = returnTwoThings("uno", "dos")
	printMe(returnedValue1+" "+returnedValue2)
	printFormatted("primer argumento", "segundo argumento")
	var numerator int = 4
	var denominator int = 0 


	var divisionResult, remainder , err = division(numerator,denominator)
	// manejo de errores 
	if err != nil{
	   fmt.Print(err.Error())
	}else if remainder==0 {
		fmt.Printf("Resultado: %v", divisionResult)
	}else{
		fmt.Printf("Resultado: %v Resto: %v", divisionResult,remainder)
	}
}

// función void
func printMe(value string){
	fmt.Println(value)
}
// función que retorna algo 
func retunrSomething(value string) string{
	return value
}
// función que retorna más de un valor
func returnTwoThings(value1 string , value2 string)(string, string){
	return "valor1: "+value1,"valor2: "+value2
}
// manejo de string con formato 
func printFormatted(value1 string, value2 string){
	fmt.Println("Utilizando printF")
	fmt.Printf("valor 1 = %v, valor 2 = %v" , value1, value2)
}

func division(numerator int , denominator int) (int, int, error){
	var err error
	if denominator==0{
		err = errors.New("No se puede dividir por cero")
		return 0,0, err
	}

	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
}
