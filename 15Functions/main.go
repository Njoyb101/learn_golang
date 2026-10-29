package main

import (
	"fmt"
)

func greeter() {
	fmt.Println("Welcome to Functions")
}
func greeter2() {
	fmt.Println("Defining Function Inside Another Function Definition is not Allowed")
}

func adder(a int, b int) (int, int) {
	return (a + b), (a + b + 1)
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func main() {
	greeter()
	greeter2()
	fmt.Println(adder(4, 7))
	fmt.Println(proAdder(4, 7, 10, 15))

}
