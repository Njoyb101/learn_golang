package main

import (
	"fmt"
)

func main() {
	var welcome string = "Welcome to Array"
	fmt.Println(welcome)

	// array with 5 elements
	var fruits [5] string

	fruits[0] = "apple"
	fruits[4] = "jamun"
	
	fmt.Println("Fruit List = ",fruits) // 1,2,3 positions are blank  and no comma 
	fmt.Println("Fruit List length= ",len(fruits)) // len = 5 even if no data

	var vegetables = [3] string {"potato","beans","cabbage"} // another way to declare this

	fmt.Println(vegetables)
}
