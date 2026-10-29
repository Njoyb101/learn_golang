package main

import (
	"fmt"
)

func main() {
	var welcome string = "Welcome to Maps"
	fmt.Println(welcome)

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "Python"
	languages["rb"] = "Ruby"

	fmt.Println("List of all Languages =",languages)
	fmt.Println("'py' is short for ",languages["py"])
	delete(languages,"rb")
	fmt.Println("List of all Languages After delete =",languages)

	for key,value := range languages {
		fmt.Printf("For key %v, value is %v \n", key,value)
	}
	for key, _ := range languages {
		fmt.Printf("key is  %v\n", key)
	}

}
