package main

import (
	"fmt"
	"github.com/CertifiedDeveloperDH/go_course/functions/function"
)

func main() {
	//var myIntVar int64
	//function.Display(myIntVar)

	v :=function.Add(4,2)
	fmt.Println(v)

	function.RepeatString(10, "LA")

	value, err := function.Calc(function.SUM, 20.12, 34)
	fmt.Println("value: ", value, " - error: ", err)

	xVal, yVal := function.Split(40)
	fmt.Println(xVal, yVal)

	val2 := function.MSum(1,2,3,1,2,3,4)
	fmt.Println(val2)

	mOperVal, err := function.MOperations(function.SUM, 4,4,4,4,10,80)
	fmt.Println("value: ", mOperVal, " - err:", err)

	factOpFunc := function.FactoryOperation(function.SUM)
	fmt.Println(factOpFunc(10,15))
}