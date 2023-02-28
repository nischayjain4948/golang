package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func PerformPostJsonRequest() {
	const myURL = "http://localhost:3000/api/golang/formdata"

	// form Data
	data := url.Values{}

	data.Add("firstName", "nischay")
	data.Add("lastName", "jain")
	data.Add("email", "nischayjain4948@gmail.com")

	response, err := http.PostForm(myURL, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("The reponse from the server is: ", string(content))

}

func main() {
	fmt.Println("This file is used to sen the form data")
	PerformPostJsonRequest()

}
