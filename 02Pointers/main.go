package main

import "fmt"

const metersToYard float64 = 1.09361

func main(){
	number := 3
	fmt.Printf("The value is %v \n", number)
	fmt.Println("The address is ", &number) //printing memory address of variable

	var pointer *int = &number
	fmt.Println(pointer) //should pointing to the same address where number reside
	*pointer = 5 //changing value that reside at certain address
	fmt.Println(number) //is now 5
}