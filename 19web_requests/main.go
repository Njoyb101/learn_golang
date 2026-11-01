/*
You need to Close the connection manually
*/
package main

import (
	"fmt"
	"io"
	"net/http"
)

const worldsFirstWebPage string = "https://info.cern.ch/hypertext/WWW/TheProject.html"

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	fmt.Println("Handling web requests using nethttp")

	res, err := http.Get(worldsFirstWebPage)

	checkNilError(err)

	fmt.Printf("response is of type : %T\n", res)

	defer res.Body.Close()

	dataBytes, err := io.ReadAll(res.Body)

	checkNilError(err)

	content := string(dataBytes)
	fmt.Println(content)
}
