package main

import "fmt"

func main(){
	fmt.Println("For loops")
	arr := [3]int32{1,2,3}

	for index, value:= range arr{
		fmt.Printf("index:%v value: %v \n" , index, value)
	}

	// comportamiento de un while 
	var i = 0
	for i<10{
		fmt.Println(i)
		i = i+1
	}
	fmt.Println("usando break")
	var j = 0 
	for{
		if j>10{
			break
		}
		fmt.Println(j)
		j = j+1
	}
	// js way
	fmt.Println("javascript")
	for i:=0; i<10;i++{
		fmt.Println(i)
	}
}