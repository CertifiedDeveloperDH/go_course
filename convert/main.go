package main

import (
	"fmt"
	//"unsafe"
	"strconv"
)

func main() {

	floatVar := 33.11
	fmt.Printf("type: %T, value: %f \n", floatVar, floatVar)

	floatStrVar := fmt.Sprintf("%.2f", floatVar)
	fmt.Printf("type: %T, value: %s\n", floatStrVar, floatStrVar)

	intVar2 := 342
	intStrVar2 := strconv.Itoa(intVar2)
	fmt.Printf("type: %T, value: %s\n", intStrVar2, intStrVar2)

	strIntVar, err := strconv.Atoi("15")
	fmt.Printf("type: %T, value: %d, err:%v \n", strIntVar, strIntVar, err)

	strInvVar3, _ := strconv.ParseInt("10",10,64)
	fmt.Printf("type: %T, value: %d \n", strInvVar3, strInvVar3)

	strFloatVar, err := strconv.ParseFloat("-11.2", 64)
	fmt.Printf("type: %T, value: %f \n", strFloatVar, strFloatVar)
}
