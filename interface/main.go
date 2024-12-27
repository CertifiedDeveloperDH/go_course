package main

import (
	"fmt"
	"github.com/CertifiedDeveloperDH/go_course/interface/vehicles"
)

func main(){
	Display(123)
	Display("123")
	Display(true)
	Display(123.28)

	vArray := []string{"CAR", "MOTORCYCLE", "TRUCK", "MOTORCYCLE", "TRUCK", "CAR"}

	var d float64
	for _, v := range vArray{
		fmt.Printf("Vehicle: %s\n", v)

		veh, err := vehicles.New(v, 400)
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println()
			continue
		}

		distance := veh.Distance()
		fmt.Printf("Distance: %.2f\n", distance)
		fmt.Println()
		d += distance
	}
	fmt.Println("Total distance: %.2\n", d)
}

func Display(value interface{}){
	fmt.Println(value)
}