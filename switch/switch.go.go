package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to switch & case in golang")
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)
	switch diceNumber {

	case 1:
		fmt.Println("The dice value is one", diceNumber)

	case 2:
		fmt.Println("The dice value is two", diceNumber)

	case 3:
		fmt.Println("The dice value is three", diceNumber)
		fmt.Println("The fallthrough is used to fall down in the switch case!!!!", diceNumber)
		fallthrough //  fallthrough case is used to fall down in the switch CASE.

	default:
		fmt.Println("This is default case!!", diceNumber)

	}
}
