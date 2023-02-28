package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("This file is used to consume the json Data")
	decodeJson()
}

func decodeJson() {
	const myURL = "http://localhost:3000/api/golang/json"
	response, err := http.Get(myURL)

	defer response.Body.Close()

	if err != nil {
		panic(err)
	}
	responseByte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	jsonDataFromWeb := string(responseByte)
	checkValid := json.Valid([]byte(jsonDataFromWeb))

	if checkValid {
		fmt.Println("Json Data is valid")
		fmt.Println(jsonDataFromWeb)
	} else {
		fmt.Println("Not a valid Json Data")
	}

}
