package main

import (
	"fmt"
)

func main() {
	myArrayVar := [5]int{3,6,9,10, 16}
	fmt.Println("array: ", myArrayVar, " - len:", len(myArrayVar))

	mySliceVar := []int{}
	mySliceVar = append(mySliceVar, 12, 34, 54)
	fmt.Println("slice: ", mySliceVar, " - len:", len(mySliceVar))

	mySliceVar2 := myArrayVar[2:4]
	fmt.Println("slice: ", mySliceVar2, " - len:", len(mySliceVar2), " address: ", &mySliceVar2[0])

	mySliceVar4 := make([]int, 3)
	fmt.Println("slice: ", mySliceVar4, " - len:", len(mySliceVar4))

	mySliceVar5 := []int{1,2,6,11,20,5,1,0}
	fmt.Println(mySliceVar5)
}
