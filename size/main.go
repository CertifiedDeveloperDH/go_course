package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var my8BitsUnitVar uint8
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my8BitsUnitVar, my8BitsUnitVar, unsafe.Sizeof(my8BitsUnitVar), unsafe.Sizeof(my8BitsUnitVar)*8)

	var myFloat32Var float32
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", myFloat32Var, myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myStringVar3 string = "test1234"
	fmt.Printf("type: %T, value: %s, bytes: %d, bits: %d \n", myStringVar3, myStringVar3, unsafe.Sizeof(myStringVar3), unsafe.Sizeof(myStringVar3)*8)

	myStringVar4 := `Mi variable de tipo texto en go
con multiples
lineas
:)`
	fmt.Printf("mi valor es %s \n", myStringVar4)
}
