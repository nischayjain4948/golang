package main

import "fmt"

func main() {
	fmt.Println("Welcome to the  course of maps Data structure")

	language := make(map[string]string)

	language["JS"] = "javaScript"
	language["RB"] = "Ruby"
	language["PY"] = "python"
	fmt.Println("List of maps: ", language)
	fmt.Println("RB in language is: ", language["RB"])

	// iterate to the language maps

	for key, value := range language {
		fmt.Println(key, "-", value)
	}

}
