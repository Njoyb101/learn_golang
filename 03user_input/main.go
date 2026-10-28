package main

import (
	"bufio"
	"fmt"
	"os"
)



func main() {
	var welcome string = "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the Pizza rating ( out of  10) : ")

	// Comma Ok  || Error Ok Syntax  instead of try catch
	input, _ := reader.ReadString('\n') // '_' is throw away like python
	fmt.Println("Thanks  for rating us ", input)
	
}
