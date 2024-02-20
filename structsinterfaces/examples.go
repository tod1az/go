package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

type engine interface{
	milesLeft() uint8
}

func (e electricEngine) milesLeft() uint8{
	return e.kwh*e.mpkwh
}
func (e gasEngine) milesLeft() uint8{
	return e.gallons*e.mpg
}

func canMakeIt(e engine, miles uint8){
	if miles <= e.milesLeft(){
		fmt.Println("ÿes 😎")
	}else{
		fmt.Println("no 😢")
	}
}

// agregar un "*" en frente del struct hace que se obtenga la referencia en memoria del objeto
// por lo tanto no se producen copias innecesarias y se pueden modificar directamente los valores de sus atributos

func (engine *gasEngine) modify(command string){
	if command == "increment"{
		engine.gallons = engine.gallons+1
	}
	if command == "decrement"{
		engine.gallons = engine.gallons-1
	}
}

func main() {
	var defaultEngine gasEngine
	// al no definir las propiedades, las propiedades del objeto son almacenadas con un valor por defecto
	fmt.Println(defaultEngine.gallons, defaultEngine.mpg)
	
	var myEngine gasEngine = gasEngine{mpg:25, gallons: 15}
	fmt.Println(myEngine.gallons, myEngine.mpg)
	myEngine.modify("increment")
	fmt.Println(myEngine.gallons, myEngine.mpg)

	var myElectricEngine electricEngine = electricEngine{25,15}
	var myGasEnince gasEngine = gasEngine{25,15}
	canMakeIt(myElectricEngine, 255) 
	canMakeIt(myGasEnince, 50) 
}