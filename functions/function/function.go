package function

import "fmt"

func Display(myValue int) {
	fmt.Println()
	fmt.Printf("type: %T, value: %d, address: %v\n", myValue, myValue, &myValue)
	fmt.Println()
}