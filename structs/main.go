package main

import (
	"fmt"
	"encoding/json"
	"github.com/CertifiedDeveloperDH/go_course/structs/commerce"
)

type User struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	Age uint8 `json:"age"`
	LastName string `json:"last_name"`
}

func (u User) Display(){
	v, err := json.Marshal(u)

	fmt.Println(err)
	fmt.Println(string(v))
}

func (u *User) SetName (name string){
	u.Name = name
}
func main(){
	user := User {
		ID: 123, 
		Name: "Nahuel",

	}
	fmt.Println(user)
	
	user.Display()
	user.SetName("Azul")
	user.Display()

	p1 := commerce.Product{
		Name: "Heladera marca sarasa",
		Price: 200000,
		Type: commerce.Type{
			Code: "A",
			Description: "Electrodomestico",
		},
		Tags: []string{"heladera", "freezer", "saraza", "refrigerador"},
		Count: 5,
	}
	
	p2 := commerce.Product{
		Name: "Naranja",
		Price: 50,
		Type: commerce.Type{
			Code: "B",
			Description: "Alimento",
		},
		Tags: []string{"alimento", "verdura"},
		Count: 20,
	}
	
	car := commerce.NewCar(1234)
	car.AddProducts(p1, p2)
	
	fmt.Println("Products Car")
	fmt.Println("Total Products: ", len(car.Products))
	fmt.Println("Total Car %.2f\n", car.Total())
	fmt.Println()
}

