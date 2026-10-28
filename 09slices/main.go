package main

import (
	"fmt"
	"sort"
)

func main() {
	var welcome string = "Welcome to Slices"
	fmt.Println(welcome)

	var fruits = []string{"mango", "guava"}
	fmt.Printf("Types of fruits is %T \n", fruits)
	fmt.Println(fruits)

	fruits = append(fruits, "banana")
	fmt.Println("New Fruits list = ", fruits)

	fruits = append(fruits, "peach")
	fmt.Println("fruits =", fruits)
	fmt.Println("fruits[1:] =", fruits[1:])
	fmt.Println("fruits[1:3] =", fruits[1:3])
	fmt.Println("fruits[:2] =", fruits[:2])
	fmt.Println("fruits[0] =", fruits[0])
	// fmt.Println("fruits[-1] =",fruits[-1]) // no negatives

	fmt.Println("=========================================\n")
	highScores := make([]int, 4)
	highScores[0] = 365
	highScores[1] = 366
	highScores[2] = 367
	highScores[3] = 368
	// highScores[4] = 888 // --> Fails  as mem overshoot

	fmt.Println("highScores = ", highScores)

	highScores = append(highScores, 658, 123) // memory reallocation . List holds

	fmt.Println("New highScores = ", highScores)
	// does not return value   sort in  place then print
	sort.Ints(highScores)
	fmt.Println("sorted highScores = ", highScores)

	fmt.Println("=========================================\n")
	fmt.Println("How to Remove a value from slice based on Index")
	
	var  courses = []string{"python","pyspark","sql","aws","gcp","snoflake","dbt"}
	fmt.Println("course =", courses)
	var indexToRemove  int = 2
	fmt.Println("Value at index ",indexToRemove, " = ",courses[indexToRemove])
	courses = append(courses[:indexToRemove],courses[indexToRemove+1:]...)
	fmt.Println("course after removing index",indexToRemove," =", courses)
}
