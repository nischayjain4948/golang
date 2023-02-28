package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name               string `json:"coursename"`
	Platform, Password string
	Price              int
	Tags               []string `json:"tags,omitempty"`
}

func DecodeJson() {

	jsonDataFromWeb := []byte(`  {
                "Name": "reactNative",
                "Platform": "youtube",
                "Password": "helloNative",
                "Price": 4500,
                "tags": [
                        "reactNative",
                        "web-development",
                        "bootcamp" ]
        }`)

	var course Course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json Data is valid")
		json.Unmarshal(jsonDataFromWeb, &course)
		fmt.Printf("%#v", course)

	} else {
		fmt.Println("JSON is not valid")
	}

}

func main() {
	fmt.Println("This file is mainly used for handling the json data and consuming the json data")
	DecodeJson()

}
