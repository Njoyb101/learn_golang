package main

import (
	"fmt"
	"time"
)

func main() {
	var welcome string = "Welcome to Time management"
	fmt.Println(welcome)

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:05:05 Mon"))
	fmt.Println(presentTime.Format(time.UnixDate)) 

	createdDate := time.Date(2026,time.October,25,00,00,00,00,time.UTC)
	fmt.Println(createdDate)

}
