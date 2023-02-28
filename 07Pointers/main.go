package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This file is used to take inputs from cmd")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating -: ")
	myrating, _ := reader.ReadString('\n')
	fmt.Println("my ratings are :- ", myrating)
}
