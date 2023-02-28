package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GETRequest() {
	const myURL = "http://localhost:3000/api/golang"
	response, err := http.Get(myURL)
	if err != nil {
		fmt.Println("This is an err block")
		panic(err)
	}
	fmt.Println("count is: ", response.ContentLength)

	defer response.Body.Close() // this line will be execute after all whole block will be completely execute..
	var responseString strings.Builder
	responseByte, err := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(responseByte)

	fmt.Println("byteCount is: ", byteCount)
	fmt.Println(responseString.String())

	if err != nil {
		fmt.Println("This is an err block")
		panic(err)
	}

	fmt.Println("Response from an API: ", string(responseByte))

}

func main() {
	fmt.Println("This file is used to read a web service!")
	GETRequest()

}
