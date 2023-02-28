package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func postRequest() {
	const myurl = "http://localhost:3000/api/golang/post"

	// fake json payload
	requestBody := strings.NewReader(`{"name":"nischay", "courseName":"golang", "price":3999, "pincode":311602, "isActive":true}`)
	reponse, err := http.Post(myurl, "application/json", requestBody) // url, content-type, requestBody
	if err != nil {
		fmt.Println("Error while calling the api", err)
		panic(err)
	}

	defer reponse.Body.Close()

	byteData, err := ioutil.ReadAll(reponse.Body)
	if err != nil {
		fmt.Println("Error wile reading the body")
	}
	fmt.Println("reading the reponse : ", string(byteData))
}

func main() {
	fmt.Println("This file is used for post request....")
	postRequest()

}
