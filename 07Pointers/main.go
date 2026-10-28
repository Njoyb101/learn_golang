package main

import (
	"fmt"
)

func main() {
	var welcome string = "Welcome to Time management"
	fmt.Println(welcome)

	var ptr *int // Pointer to store integer  pointer '*'
	fmt.Println("Default pointer value =", ptr)

	myNumber := 20

	var myPtr = &myNumber                               // here it holds the address of myNumber Variable
	fmt.Println("New Address Pointer = ", myPtr)        // returns the address
	fmt.Println("New Address Pointer Value = ", *myPtr) // returns the value the address holds

	fmt.Println("Update Pointer not Variable")
	fmt.Println("===========================")
	*myPtr = *myPtr + 10
	fmt.Println("New Pointer Value, Address =(", *myPtr, ",", myPtr, ")")
	fmt.Println("New Variable Value =", myNumber)
	fmt.Println("===========================")
	fmt.Println("You update the Pointer --> You update the variable value but the Address remains same.")
	fmt.Println("--------------------------")

	fmt.Println("Update Variable not Pointer")
	fmt.Println("===========================")

	myNumber = myNumber + 15
	fmt.Println("New Variable Value =", myNumber)
	fmt.Println("New Pointer Value =", *myPtr)
	fmt.Println("New Pointer Address =", myPtr)

	fmt.Println("===========================")
	fmt.Println("You update the Variable --> You update the pointr value but the Address remains same.")
	fmt.Println("--------------------------")

}
