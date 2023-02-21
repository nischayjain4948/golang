package main

import (
	"fmt"
)

type User struct {
	Name, City, Email string
	Pincode, Age      int
	Status            bool
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func main() {
	fmt.Println("Welocome to the methods in GOLang!")
	u1 := User{"nischay", "jaipur", "nischayjain4948", 311602, 23, true}
	fmt.Printf("The complete struct is available in the %+v\n", u1)
	u1.GetStatus()

}
