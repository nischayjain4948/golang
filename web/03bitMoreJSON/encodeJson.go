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

func Encodejson() {

	learnCourse := []Course{
		{"reactNative", "youtube", "helloNative", 4500, []string{"reactNative", "web-development", "bootcamp"}},
		{"reactjs", "youtube", "helloreact", 4000, []string{"reactjs", "front-end", "bootcamp"}},
		{"nodejs", "youtube", "hellonode", 5100, []string{"nodejs", "web-development", "bootcamp"}},
		{"golang", "youtube", "helloGO", 2500, nil},
	}

	//    package this data as JSON data
	finalJson, err := json.MarshalIndent(learnCourse, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func main() {
	fmt.Println("Welcome to JSON video")
	Encodejson()
}
