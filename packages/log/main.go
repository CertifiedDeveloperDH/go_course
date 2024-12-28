package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("test")

	//log.Fatal("my error")

	//log.Panic("my error")

	date := time.Now()
	file, err := os.Create(fmt.Sprintf("%d.log", date.UnixNano()))
	if err != nil {
		log.Panic(err.Error())
	}
	l := log.New(file, "success: ", log.LstdFlags|log.Lshortfile)
	l.Println("test 1")
	l.Println("test 2")
	l.Println("test 3")

	l.SetPrefix("errors: ")
	l.Println("test 1")
	l.Println("test 2")
	l.Println("test 3")
}
