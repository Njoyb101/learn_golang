package main

import (
	"fmt"
)

type User struct {
	// First letter capital to make it exportable
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	var welcome string = "Welcome to Structs. It is base of classes"
	fmt.Println(welcome)
	fmt.Println("There is no Inheritance or Super or Parent in Go !!")

	njoy :=User{"njoy","njoy@go.dev",true,30}
	fmt.Println(njoy)
	fmt.Printf("details are = %+v \n", njoy)
	fmt.Printf("name is = %+v \n", njoy.Name)
}
