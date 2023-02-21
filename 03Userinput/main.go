package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input :)"
	fmt.Println(welcome)

	// reading the data from terminal

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the raing for our Pizza :")

	// Comma ok ,  err  syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("The type of rating is %T ", input)

}
