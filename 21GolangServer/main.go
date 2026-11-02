package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Server Video")
	// performGetReq()
	// performPostReq()
	performPostFormReq()
}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
func performGetReq() {
	const url string = "http://localhost:8000/get"
	res, err := http.Get(url)

	checkNilError(err)

	defer res.Body.Close()

	fmt.Println("Status Code: ", res.StatusCode)
	fmt.Println("ContentLength: ", res.ContentLength)

	contentByte, err := io.ReadAll(res.Body)
	checkNilError(err)

	// Normal
	fmt.Println("Normal Way :")
	fmt.Println(string(contentByte))
	fmt.Println("================")

	// Advanced
	fmt.Println("Old/Other Way :")
	var responseString strings.Builder
	byteCount, _ := responseString.Write(contentByte)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
	fmt.Println(responseString.Cap())
	fmt.Println("================")
}

func performPostReq() {

	const url string = "http://localhost:8000/post"
	const JSON string = "application/json" // from post man header "Content-Type" of the response

	// Fake JSON Payload
	requestBody := strings.NewReader(`
		{
			"country":"india"
		}
	`)
	res, err := http.Post(url, JSON, requestBody)

	checkNilError(err)

	defer res.Body.Close()

	contentByte, err := io.ReadAll(res.Body)
	checkNilError(err)
	fmt.Println(string(contentByte))
}
func performPostFormReq() {
	/*
		Encrypted Data
	*/
	const myUrl string = "http://localhost:8000/postform"
	const JSON string = "application/json" // from post man header "Content-Type" of the response

	// Form Data
	data := url.Values{}
	data.Add("country", "India")

	res, err := http.PostForm(myUrl, data)

	checkNilError(err)

	defer res.Body.Close()

	contentByte, err := io.ReadAll(res.Body)
	checkNilError(err)
	fmt.Println(string(contentByte))
}
