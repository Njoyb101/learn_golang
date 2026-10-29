package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var welcome string = "Welcome to Switch Case"
	fmt.Println(welcome)

	// rand.Seed(time.Now().UnixNano()) //-->  depricated 

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is =", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1 so you can open")
	case 2:
		fmt.Println("Dice Value is 2 Move 2 Forward")
	case 3:
		fmt.Println("Dice Value is 3 Move 3 Forward")
	case 4:
		fmt.Println("Dice Value is 4 Move 4 Forward")
	case 5:
		fmt.Println("Dice Value is 5 Move 5 Forward")
	case 6:
		fmt.Println("Dice Value is 6 Move 6 Forward & you get an extra turn")
		fallthrough //  runs the next  here if you get a 6 then default will also run 
	default:
		fmt.Println("what was it ??")
	}

}
