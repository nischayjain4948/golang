package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("This file is used to web reques")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The type of response is %T\n", response)
	defer response.Body.Close() // caller's responseblity to close the connections while calling the api!!

	dataBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataBytes))

}
