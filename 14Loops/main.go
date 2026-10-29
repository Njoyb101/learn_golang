package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var welcome string = "Welcome to Switch Case"
	fmt.Println(welcome)

	days := []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"} // my slice

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println("=====================")
	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("=====================")
	for i, day := range days {
		fmt.Println("Index is= ", i, " value is = ", day)
	}

	fmt.Println("=====================")
	rougueValue := rand.IntN(9)+1
	for rougueValue < 10 {

		if rougueValue==2{
			goto lco
		}
		if rougueValue == 5{
			fmt.Println("Breakkkkkk!!!!")
			break
		}
		if rougueValue == 6{
			rougueValue+=2
			fmt.Println("Continue!!!!")
			continue
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

// label
lco:
	fmt.Println("Jumpinggggg Jackss")
}
