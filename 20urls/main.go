/*
You need to Close the connection manually
*/
package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?course_name=sql&course_id=dfwwrr32mkcdmgvisql"

func main() {
	fmt.Println("Handling URLs...")
	result, _ := url.Parse(myUrl)
	fmt.Println("result.... = ", result)
	fmt.Println("result ForceQuery = ", result.ForceQuery)
	fmt.Println("result Fragment = ", result.Fragment)
	fmt.Println("result Host = ", result.Host)
	fmt.Println("result OmitHost = ", result.OmitHost)
	fmt.Println("result Opaque = ", result.Opaque)
	fmt.Println("result Path = ", result.Path)
	fmt.Println("result RawFragment = ", result.RawFragment)
	fmt.Println("result RawPath = ", result.RawPath)
	fmt.Println("result RawQuery = ", result.RawQuery)
	fmt.Println("result Scheme = ", result.Scheme)
	fmt.Println("result User = ", result.User)
	fmt.Println("result Query = ", result.Query())

	for key, val := range result.Query() {
		fmt.Println("key, value is : ", key, "===>", val)
	}

	partsOfUrl := &url.URL{ // pass a ref : '&' is a must
		Scheme:  "https",
		Host:    "lco.dev",
		RawPath: "user-njoy",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Created URL =", anotherUrl)
}
