package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var welcome string = "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the Pizza rating ( out of  10) : ")

	// Comma Ok  || Error Ok Syntax  instead of try catch
	input, _ := reader.ReadString('\n') // '_' is throw away like python

	// input has the value 'X\n' this\n gets added automatically

	// Now convert this rating to  Integer
	// Trim the extra '\n'
	inputWithoutNewLine := strings.TrimSpace(input)
	numRating,err := strconv.ParseFloat(inputWithoutNewLine,64)

	if err != nil {
		fmt.Println(err)
		panic(err) // exit(1)
	} else{
		// Add 1 to the rating
		numRating = numRating +1
	}

	// Final Message
	fmt.Println("Thanks  for rating us ", numRating)
}
