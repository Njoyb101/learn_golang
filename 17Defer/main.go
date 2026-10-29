// LIFO Last Differ executes first
package main

import (
	"fmt"
)

func main() {
	var welcome string = "Welcome to Defer"
	
	// on Stack (Open Coded)  as  it is not in a loop Go 1.14+
	defer fmt.Println("Hello World1") // Puts at the last line of the function
	defer fmt.Println("Hello World2") // Puts at the 2nd last line of the function REVERSE ORDER (LIFO)

	fmt.Println(welcome)
	myDiffer()

}

// on Heap Mem ( Dynamic Stack) Go 1.14+
func myDiffer(){
	for i:=0; i<5; i++ {
		defer fmt.Println(i)
	}
}

/*
Welcome to Defer
4
3
2
1
0
Hello World2
Hello World1
*/