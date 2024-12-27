package mypackage

import "fmt"
func Run(){

	defer func(){
		fmt.Println("end my function")
		r := recover()
		if r != nil {
			fmt.Println("error in run: ",r)
		}
	}()
	//panic("panic in run function")
}