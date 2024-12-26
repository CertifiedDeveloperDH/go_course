package main

import (
	"fmt"
)

func main() {
	myOtherScopeVariable := 50
	{
		myScopeVariable := 40

		fmt.Println("Mi variable: ", myOtherScopeVariable)
		fmt.Println("Mi variable: ", myScopeVariable)
	}
}
