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

// Method
func (user User) GetStatus() string {
	// Passes a copy of it
	if user.Status {
		return "active"
	} else {
		return "inactive"
	}
}

func (user User) NewEmail()  {
	// Passes a copy of it
	user.Email = "123@gmail.com"
	fmt.Println("Email from Method is =", user.Email)
}

func main() {
	var welcome string = "Welcome to Structs. It is base of classes"
	fmt.Println(welcome)
	fmt.Println("There is no Inheritance or Super or Parent in Go !!")

	njoy := User{"njoy", "njoy@go.dev", true, 30}
	fmt.Println(njoy)
	fmt.Printf("details are = %+v \n", njoy)
	fmt.Printf("name is = %+v \n", njoy.Name)

	fmt.Println("User is ", njoy.GetStatus())

	njoy.NewEmail()
	fmt.Println("Email from Main is =", njoy.Email)
	fmt.Println("Take Away : In the Method the value changed but not in the original Object")
}
