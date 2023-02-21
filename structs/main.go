package main

import "fmt"

func main() {
	fmt.Println("Welcome to the lecture of structs..")

	// no inheritance in golang
	// only composition

	nischay := User{"nischay", "jaipur", "nischayjain4948@gmail.com", true, 311602}
	fmt.Println(nischay)

	fmt.Printf("nschay details are %+v\n", nischay)
	fmt.Printf("Name is %v and email is %v", nischay.name, nischay.email)

}

type User struct {
	name, city, email string
	status            bool
	pincode           int
}
