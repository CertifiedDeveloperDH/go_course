package main

import (
	"fmt"
)

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println("My variable is", myIntVar)

	var myUintVar uint
	myUintVar = 12
	fmt.Println("My variable is", myUintVar)

	var myStringVar string
	myStringVar = "my string variable"
	fmt.Println("My variable is", myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("My variable is", myBoolVar)

	fmt.Println("My variable address is:", &myStringVar)

	myIntVar2 := 12
	fmt.Println("My variable is:", myIntVar2)

	myStringVar2 := "my string variable with :="
	fmt.Println("My variable is:", myStringVar2)

	myBoolVar2 := 12
	fmt.Println("My variable is:", myBoolVar2)

	const myIntConst int = 12
	fmt.Println("Mi constante es: ", myIntConst)

	const myFirstStringConst = "a12"
	fmt.Println("Mi constante es: ", myFirstStringConst)

	const myStringConst string = "a12"
	fmt.Println("Mi constante es: ", myStringConst)
}
