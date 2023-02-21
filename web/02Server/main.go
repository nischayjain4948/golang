package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func PerformGetRequest() {
	const myurls = "http://localhost:3000/api/golang"

	response, err := http.Get(myurls)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	readResponse, error := ioutil.ReadAll(response.Body)

	if error != nil {
		panic(error)
	}
	fmt.Println("reponse body is :", string(readResponse))
	fmt.Println("response length is", response.ContentLength)
	fmt.Println("Status code is", response.StatusCode)
}

func main() {
	fmt.Println("This file is having the code of creating the server!!")
	PerformGetRequest()
}
