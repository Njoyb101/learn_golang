package main

import (
	"fmt"
)

func main() {
	var welcome string = "Welcome to If Else"
	fmt.Println(welcome)

	loginCount := 100

	var result string

	if loginCount < 10 {
		result = "Regular User"
	}else if loginCount >= 10 && loginCount< 100{
		result = "Very Regular User"
	}else{
		result = "SCAMMMM"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even Number")
	}else{
		fmt.Println("Odd Number")
	}

	if num := 3; num <= 10 {
		fmt.Println("Less than equal to 10")
	}else{
		fmt.Println("More than 10")
	}
}
