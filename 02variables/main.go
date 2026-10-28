package main

import "fmt"

// jwtToken :=10000  // walrus operator  not allowed outside the methods
const myToken string = "not a public Constant" // not a public var
const MyToken string = "A public Constant as First Letter Capital" // A public var as First Letter Capital

func main() {
	// var usename string = "njoy"
	var isLoggedin bool = true

	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type: %T \n", isLoggedin)

	var defaultInteger int

	fmt.Println(defaultInteger)
	fmt.Printf("Variable is of type: %T \n", defaultInteger)

	var defaultString string

	fmt.Println(defaultString)
	fmt.Printf("Variable is of type: %T \n", defaultString)

	// implicit non changable type
	var  website= "www.google.com" 
	fmt.Println(website)

	// No Var Style
	// Walrus operator   allowed only inside the methods
	numOfUsers := 50000
	fmt.Println(numOfUsers)
	fmt.Printf("No of Users = %d \n", numOfUsers)

	fmt.Println(myToken)
	fmt.Println(MyToken)
}
