package main

import (
	"fmt"
	"strings"
)


func main(){
	var myString = "résumé"
	var indexed = myString[1]
	fmt.Println(myString)
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString{
		fmt.Println(i,v)
	}

	var stringLength = len(myString)
	fmt.Println(stringLength)// es el largo del array de bytes que representa al string
	var myRune = []rune(myString) 
	var runeLength = len(myRune)// representa la cantidad de caracteres
	fmt.Println(runeLength)

	var strSlice = []string{"a", "b", "c","d","e"}
	var strBuilder strings.Builder
	// La forma mas eficiente de concatenar strings es con un string builder, 
	// sumando los caracteres de forma convencional se hace un uso de memoria innecesario
	// ya que se crea un nuevo string por cada caracter que es concatenado
	fmt.Println(strSlice)
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}