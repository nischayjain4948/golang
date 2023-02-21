package main

import "fmt"

func main() {
	fmt.Println("Variables")

	var username string = "Nischay jain"
	fmt.Println(username)
	fmt.Printf("Variable is type: %T \n ", username)

	isLoggedIn := true
	fmt.Printf("Variable is type: %T \n ", isLoggedIn)

	// implicit type
	var website string = "learntocode@nischayjain"
	fmt.Println(website)

	//  no var style
	numberOfUsers := 300000 //--->>> Walrus operator
	fmt.Println(numberOfUsers)

}
