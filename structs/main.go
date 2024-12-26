package main

import (
	"fmt"
	"encoding/json"
	"github.com/CertifiedDeveloperDH/go_course/structs/commerce"
)

type User struct{
	ID int `json:"id, omitempty"`
	Name string `json:"name, omitempty"`
	Address string `json:"address, omitempty"`
	Age uint8 `json:"age, omitempty"`
	LastName string `json:"last_name, omitempty"`
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
}

